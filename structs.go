package wrapper

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Session is the main struct for the session
type Session struct {
	Auth        string
	Socket      *websocket.Conn
	SocketMutex sync.Mutex
	Handlers    map[string]*func(Session, interface{})
	Headers     http.Header
	Rooms       []string
	Token       string
	Room        string
	Log         bool
}

// Payload is the typical payload, this should be able to be used 99% of the time
type Payload struct {
	Data interface{} `json:"data"`
	Room string      `json:"room"`
	Type string      `json:"type"`
}
