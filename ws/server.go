package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Connection interface {
	Send([]byte) bool
	Close()
}

type Server interface {
	Handler(http.ResponseWriter, *http.Request)
	Append(Connection)
	Remove(c Connection)
	Broadcast([]byte)
}

type server struct {
	connections map[Connection]bool
	append      chan Connection
	remove      chan Connection
	broadcast   chan []byte
}

var (
	instanceServer *server
	onceServer     sync.Once
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *server) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	//TODO : error
	if err != nil {
		log.Println(err)
		return

	}
	accept(s, conn)
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
func GetServer() Server {
	onceServer.Do(func() {
		instanceServer = &server{
			connections: make(map[Connection]bool),
			append:      make(chan Connection),
			remove:      make(chan Connection),
			broadcast:   make(chan []byte),
		}
		go instanceServer.run()
	})
	return instanceServer
}

func (s *server) run() {
	ticker := time.NewTicker(3 * time.Second)
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

		case <-ticker.C:
			//THIS IS TEST
			log.Println("Ticker")
			m := struct {
				Type     string    `json:"type"`
				Model    string    `json:"model"`
				RN       string    `json:"rn"`
				Daterent time.Time `json:"daterent"`
				Dateret  time.Time `json:"dateret"`
				Agent    string    `json:"agent"`
			}{"Мопед", "AUDI", "AAA", time.Now(), time.Now(), "Смирнов Иван Иванович"}
			b, _ := json.Marshal(&m)
			go func() { s.Broadcast(b) }()
			log.Println("Ticker done")

		}
	}
}
