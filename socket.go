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
// Rooms is a list of rooms to join, assumed to join all but you can set specific rooms. The current all known rooms is "chat", "crash", "shop", "coinflip", "jackpot", "jackpot-low", "supply-drops", "mines"
// Room can either be "en", "tr", or "ru". If none is supplied it assumes "en"
func New(token string, rooms []string, room string) (*Session, error) {
	s := &Session{}
	if token != "" {
		s.Auth = token
	}
	headers := strings.Split("Host: rustchance.com\nPragma: no-cache\nCache-Control: no-cache\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.421\nOrigin: https://rustchance.com\nSec-WebSocket-Version: 13\nAccept-Encoding: gzip, deflate, br\nAccept-Language: en-US,en;q=0.9,zh;q=0.8\nSec-WebSocket-Extensions: permessage-deflate; client_max_window_bits", "\n")
	if s.Auth != "" {
		headers = append(headers, "Cookie: token="+s.Auth)
	}
	s.Headers = http.Header{}
	for _, header := range headers {
		parts := strings.Split(header, ": ")
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
	s.Handlers = make(map[string]func(*Session, interface{}))
	return s, nil
}

// Write writes a payload to the websocket, this is usually only used by the package but can be used by a user directly.
// toWrite should be a json payload unmarshal'd
// returns an error incase writing fails
func (s *Session) Write(toWrite interface{}) error {
	s.SocketMutex.Lock()
	err := s.Socket.WriteJSON(toWrite)
	s.SocketMutex.Unlock()
	return err
}

// Open opens the websocket connection and writes the initial payload as well as starts reading from the socket.
func (s *Session) Open() error {
	if c, _, err := websocket.DefaultDialer.Dial(SocketURL, s.Headers); err == nil {

		s.Socket = c
		err = s.Write(&Payload{
			Data: s.Rooms,
			Room: "control",
			Type: "join_rooms",
		})
		if err != nil && s.Log {
			fmt.Println(err)
		}
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
				t := m.Room + "_" + m.Type
				if f, ok := s.Handlers[t]; ok {
					switch t {
					case "shop_rules":
						p := &ShopRules{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "chat_rooms":
						p := &ChatRooms{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "chat_message":
						p := &ChatMessage{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "chat_stats":
						p := &ChatStats{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "coinflip_delete_game":
						p := &CoinflipDeleteGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "coinflip_game_status":
						p := &CoinflipGameStatus{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "coinflip_list":
						p := &CoinflipList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "coinflip_new_game":
						p := &CoinflipNewGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "coinflip_update_game":
						p := &CoinflipUpdateGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "crash_cashout":
						p := &CrashCashOut{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "crash_multiple_bets":
						p := &CrashMultipleBets{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "crash_new":
						p := &CrashNew{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "crash_start":
						p := &CrashStart{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "crash_tick":
						p := &CrashTick{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot_list":
						p := &JackpotList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot_new_deposit":
						p := &JackpotNewDeposit{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot_new_game":
						p := &JackpotNewGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot_start_timer":
						p := &JackpotStartTimer{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot-low_list":
						p := &LowJackpotList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot-low_new_deposit":
						p := &LowJackpotNewDeposit{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot-low_new_game":
						p := &LowJackpotNewGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "jackpot-low_start_timer":
						p := &LowJackpotStartTimer{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_begin_timer":
						p := &MinesBeginTimer{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_game_started":
						p := &MinesGameStarted{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_game_starting":
						p := &MinesGameStarting{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_list":
						p := &MinesList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_new_game":
						p := &MinesNewGame{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_new_player":
						p := &MinesNewPlayer{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "mines_winner":
						p := &MinesWinner{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "supply-drops_joinable":
						p := &SupplyDropsJoinable{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "supply-drops_list":
						p := &SupplyDropsList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "supply-drops_players":
						p := &SupplyDropsPlayers{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "supply-drops_result":
						p := &SupplyDropWinner{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "roulette_roll":
						p := &RouletteRoll{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "roulette_list":
						p := &RouletteList{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					case "user_set_points":
						p := &UserSetPoints{}
						err = json.Unmarshal([]byte(msg), p)
						if err != nil && s.Log {
							fmt.Println(err)
							break
						}
						f(s, p)
					default:
						break

					}
				}
			}

		}
	} else {
		return err
	}
}

// AddHandler sets something to do when an event happens, the input is a func that always has a first argument of a *Session and a second argument of another struct
func (s *Session) AddHandler(v interface{}) {
	switch a := v.(type) {
	case func(*Session, *ShopRules):
		s.Handlers["shop_rules"] = func(s *Session, v interface{}) {
			a(s, v.(*ShopRules))
		}
	case func(*Session, *ChatRooms):
		s.Handlers["chat_rooms"] = func(s *Session, v interface{}) {
			a(s, v.(*ChatRooms))
		}
	case func(*Session, *ChatMessage):
		s.Handlers["chat_message"] = func(s *Session, v interface{}) {
			a(s, v.(*ChatMessage))
		}
	case func(*Session, *ChatStats):
		s.Handlers["chat_stats"] = func(s *Session, v interface{}) {
			a(s, v.(*ChatStats))
		}
	case func(*Session, *CoinflipDeleteGame):
		s.Handlers["coinflip_delete_game"] = func(s *Session, v interface{}) {
			a(s, v.(*CoinflipDeleteGame))
		}
	case func(*Session, *CoinflipGameStatus):
		s.Handlers["coinflip_game_status"] = func(s *Session, v interface{}) {
			a(s, v.(*CoinflipGameStatus))
		}
	case func(*Session, *CoinflipList):
		s.Handlers["coinflip_list"] = func(s *Session, v interface{}) {
			a(s, v.(*CoinflipList))
		}
	case func(*Session, *CoinflipNewGame):
		s.Handlers["coinflip_new_game"] = func(s *Session, v interface{}) {
			a(s, v.(*CoinflipNewGame))
		}
	case func(*Session, *CoinflipUpdateGame):
		s.Handlers["coinflip_update_game"] = func(s *Session, v interface{}) {
			a(s, v.(*CoinflipUpdateGame))
		}
	case func(*Session, *CrashCashOut):
		s.Handlers["crash_cashout"] = func(s *Session, v interface{}) {
			a(s, v.(*CrashCashOut))
		}
	case func(*Session, *CrashMultipleBets):
		s.Handlers["crash_multiple_bets"] = func(s *Session, v interface{}) {
			a(s, v.(*CrashMultipleBets))
		}
	case func(*Session, *CrashNew):
		s.Handlers["crash_new"] = func(s *Session, v interface{}) {
			a(s, v.(*CrashNew))
		}
	case func(*Session, *CrashStart):
		s.Handlers["crash_start"] = func(s *Session, v interface{}) {
			a(s, v.(*CrashStart))
		}
	case func(*Session, *CrashTick):
		s.Handlers["crash_tick"] = func(s *Session, v interface{}) {
			a(s, v.(*CrashTick))
		}
	case func(*Session, *JackpotList):
		s.Handlers["jackpot_list"] = func(s *Session, v interface{}) {
			a(s, v.(*JackpotList))
		}
	case func(*Session, *JackpotNewDeposit):
		s.Handlers["jackpot_new_deposit"] = func(s *Session, v interface{}) {
			a(s, v.(*JackpotNewDeposit))
		}
	case func(*Session, *JackpotNewGame):
		s.Handlers["jackpot_new_game"] = func(s *Session, v interface{}) {
			a(s, v.(*JackpotNewGame))
		}
	case func(*Session, *JackpotStartTimer):
		s.Handlers["jackpot_start_timer"] = func(s *Session, v interface{}) {
			a(s, v.(*JackpotStartTimer))
		}
	case func(*Session, *LowJackpotList):
		s.Handlers["jackpot-low_list"] = func(s *Session, v interface{}) {
			a(s, v.(*LowJackpotList))
		}
	case func(*Session, *LowJackpotNewDeposit):
		s.Handlers["jackpot-low_new_deposit"] = func(s *Session, v interface{}) {
			a(s, v.(*LowJackpotNewDeposit))
		}
	case func(*Session, *LowJackpotNewGame):
		s.Handlers["jackpot-low_new_game"] = func(s *Session, v interface{}) {
			a(s, v.(*LowJackpotNewGame))
		}
	case func(*Session, *LowJackpotStartTimer):
		s.Handlers["jackpot-low_start_timer"] = func(s *Session, v interface{}) {
			a(s, v.(*LowJackpotStartTimer))
		}
	case func(*Session, *MinesBeginTimer):
		s.Handlers["mines_begin_timer"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesBeginTimer))
		}
	case func(*Session, *MinesGameStarted):
		s.Handlers["mines_game_started"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesGameStarted))
		}
	case func(*Session, *MinesGameStarting):
		s.Handlers["mines_game_starting"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesGameStarting))
		}
	case func(*Session, *MinesList):
		s.Handlers["mines_list"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesList))
		}
	case func(*Session, *MinesNewGame):
		s.Handlers["mines_new_game"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesNewGame))
		}
	case func(*Session, *MinesNewPlayer):
		s.Handlers["mines_new_player"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesNewPlayer))
		}
	case func(*Session, *MinesWinner):
		s.Handlers["mines_winner"] = func(s *Session, v interface{}) {
			a(s, v.(*MinesWinner))
		}
	case func(*Session, *SupplyDropsJoinable):
		s.Handlers["supply-drops_joinable"] = func(s *Session, v interface{}) {
			a(s, v.(*SupplyDropsJoinable))
		}
	case func(*Session, *SupplyDropsList):
		s.Handlers["supply-drops_list"] = func(s *Session, v interface{}) {
			a(s, v.(*SupplyDropsList))
		}
	case func(*Session, *SupplyDropsPlayers):
		s.Handlers["supply-drops_players"] = func(s *Session, v interface{}) {
			a(s, v.(*SupplyDropsPlayers))
		}
	case func(*Session, *SupplyDropWinner):
		s.Handlers["supply-drops_result"] = func(s *Session, v interface{}) {
			a(s, v.(*SupplyDropWinner))
		}
	case func(*Session, *RouletteRoll):
		s.Handlers["roulette_roll"] = func(s *Session, v interface{}) {
			a(s, v.(*RouletteRoll))
		}
	case func(*Session, *RouletteList):
		s.Handlers["roulette_list"] = func(s *Session, v interface{}) {
			a(s, v.(*RouletteList))
		}
	case func(*Session, *UserSetPoints):
		s.Handlers["user_set_points"] = func(s *Session, v interface{}) {
			a(s, v.(*UserSetPoints))
		}
	default:
		fmt.Println("Unknown handler type, this handler will not be called")
	}
}

// SwitchRoom takes in a room type, then a type, then a data
// Room can be for example "chat"
// Type can be for example "switch_room"
// Data can be anything, it's a type interface{}
func (s *Session) SwitchRoom(room, t string, data interface{}) error {
	err := s.Write(&Payload{
		Room: room,
		Type: t,
		Data: data,
	})
	return err
}

// SwitchChatRoom takes in a room type of "en", "tr", or "ru"
func (s *Session) SwitchChatRoom(room string) error {
	if room != "en" && room != "tr" && room != "ru" {
		return fmt.Errorf("invalid room input")
	}
	err := s.Write(&Payload{
		Room: "chat",
		Type: "switch_room",
		Data: room,
	})
	return err
}

// BetRoulette takes in an amount, a game ID, and a color to bet on roulette
// Amount is is US cents
// GameID is the game id, get this from the RouletteRoll event in the "NewGame" field of "Data"
// Color int can be 0 for blue, 1 for yellow, and 2 for red
func (s *Session) BetRoulette(amount, gameID, color int) error {
	if color < 0 || color > 2 {
		return fmt.Errorf("color out of range 0-2")
	}
	err := s.Write(&Payload{
		Room: "roulette",
		Type: "join_game",
		Data: &EnterRouletteData{
			Amount: amount,
			Color:  color,
			ID:     gameID,
		},
	})
	return err
}

// JoinSupplyDrop joins a supply drop
func (s *Session) JoinSupplyDrop() error {
	return s.Write(&Payload{
		Room: "supply-drops",
		Type: "join",
		Data: nil,
	})
}
