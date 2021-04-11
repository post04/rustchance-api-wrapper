package wrapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// New returns a *Session and err
// Token is your auth token, this can be left empty, you only need a token for account specific things
// Rooms is a list of rooms to join, assumed to join all but you can set specific rooms
// Room can either be "en", "tr", or "ru". If none is supplied it assumes "en"
func New(token string, rooms []string, room string) (*Session, error) {
	s := &Session{}
	s.Token = token
	headers := strings.Split("Host: rustchance.com\nPragma: no-cache\nCache-Control: no-cache\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.421\nOrigin: https://rustchance.com\nSec-WebSocket-Version: 13\nAccept-Encoding: gzip, deflate, br\nAccept-Language: en-US,en;q=0.9,zh;q=0.8\nSec-WebSocket-Extensions: permessage-deflate; client_max_window_bits", "\n")
	if s.Token != "" {
		headers = append(headers, "Cookie: "+s.Token)
	}
	s.Headers = http.Header{}
	for _, header := range headers {
		parts := strings.Split(header, ":")
		if len(parts) == 2 {
			s.Headers[parts[0]] = []string{parts[1]}
		}
	}
	if len(rooms) < 1 {
		s.Rooms = []string{"chat", "crash", "shop", "coinflip", "jackpot", "jackpot-low", "supply-drops", "mines"}
	} else {
		s.Rooms = rooms
	}
	if room != "" {
		s.Room = room
	} else {
		s.Room = "en"
	}
	s.Log = false
	return s, nil
}

// Write writes a payload to the websocket
func (s *Session) Write(toWrite interface{}) {
	s.SocketMutex.Lock()
	err := s.Socket.WriteJSON(toWrite)
	if err != nil && s.Log {
		fmt.Println(err)
	}
	s.SocketMutex.Unlock()
}

// Open opens the websocket connection and starts writing all the initiating payloads
func (s *Session) Open() error {
	if c, _, err := websocket.DefaultDialer.Dial(SocketURL, s.Headers); err == nil {
		s.Write(&Payload{
			Data: s.Rooms,
			Room: "control",
			Type: "join_rooms",
		})
		s.Socket = c
		for {
			_, message, err := c.ReadMessage()

			if err != nil {
				c.Close()
				s.Socket = nil
				s.Open()
				return err
			}
			for _, msg := range strings.Split(string(message), "\n") {
				var m Payload
				err = json.Unmarshal([]byte(msg), &m)
				switch m.Room + "_" + m.Type {
				case "chat_rooms":
					s.Write(&Payload{
						Room: "chat",
						Type: "switch_room",
						Data: s.Room,
					})
				}
			}

		}
	} else {
		return err
	}
}
