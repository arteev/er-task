
window.startws = function(timeout,message,open,close) {
    if (!window["WebSocket"])  {
        return false;
    }
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onopen = open; 
    conn.onclose = function(evt) {
        if (close) {
            close(evt);
        }
        conn = null;
        setTimeout(function () {
            window.startws();
        }, timeout)
    } 
    conn.onmessage = message;
    return true;
}