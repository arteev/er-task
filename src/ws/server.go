package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/arteev/er-task/src/storage"

	"github.com/gorilla/websocket"
)

type Connection interface {
	Send([]byte) bool
	Close()
}

type Server interface {
	Handler(http.ResponseWriter, *http.Request) (int, error)
	Append(Connection)
	Remove(c Connection)
	Broadcast([]byte)
}

type server struct {
	connections map[Connection]bool
	append      chan Connection
	remove      chan Connection
	broadcast   chan []byte
	notify      chan storage.Notification
}

var (
	instanceServer *server
	onceServer     sync.Once
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *server) Handler(w http.ResponseWriter, r *http.Request) (int, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not upgrader.Upgrade: %s", err)
		return http.StatusInternalServerError, err
	}
	accept(s, conn)
	return http.StatusAccepted, nil
}

func (s *server) Append(c Connection) {
	s.append <- c
}

func (s *server) Remove(c Connection) {
	s.remove <- c
}

func (s *server) Broadcast(b []byte) {
	s.broadcast <- b
}

//GetServer запускает сервет который обслуживает websocket клиентов
//при получении сообщения из канала storage.Notification отправляет
//всем подключенным клиентам
func GetServer(n chan storage.Notification) Server {
	onceServer.Do(func() {
		instanceServer = &server{
			connections: make(map[Connection]bool),
			append:      make(chan Connection),
			remove:      make(chan Connection),
			broadcast:   make(chan []byte),
			notify:      n,
		}
		go instanceServer.run()
	})
	return instanceServer
}

func (s *server) run() {
	for {
		select {
		case c := <-s.append:
			s.connections[c] = true
		case c := <-s.remove:
			if _, exists := s.connections[c]; exists {
				delete(s.connections, c)
				c.Close()
			}
		case message := <-s.broadcast:
			for c := range s.connections {
				if !c.Send(message) {
					delete(s.connections, c)
					c.Close()
				}
			}

		case n := <-s.notify:
			//Уведомление от хранилища
			nws, err := storage.RentDataFromStorage(n)
			if err != nil {
				log.Printf("Could not cast notify from storage:%v,", err)
			}
			go func() {
				b, _ := json.Marshal(nws)
				s.Broadcast(b)
			}()
		}
	}
}
