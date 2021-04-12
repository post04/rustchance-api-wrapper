package wrapper

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Session is the main struct for the session
type Session struct {
	Auth        string
	Socket      *websocket.Conn
	SocketMutex sync.Mutex
	Handlers    map[string]func(*Session, interface{})
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

type ShopRules struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data ShopRulesData `json:"data"`
}
type ShopRulesData struct {
	Enabled      bool `json:"enabled"`
	MaxItems     int  `json:"maxItems"`
	MinItemValue int  `json:"minItemValue"`
	MinItems     int  `json:"minItems"`
	MinValue     int  `json:"minValue"`
}

type SupplyDropsJoinable struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type SupplyDropsList struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data SupplyDropsListData `json:"data"`
}

type SupplyDropsListData struct {
	ID         int  `json:"id"`
	Joined     bool `json:"joined"`
	Players    int  `json:"players"`
	Processing bool `json:"processing"`
	State      int  `json:"state"`
	TimeLeft   int  `json:"timeLeft"`
}

type SupplyDropsPlayers struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type SupplyDropWinner struct {
	Room string               `json:"room"`
	Type string               `json:"type"`
	Data SupplyDropWinnerData `json:"data"`
}
type Winner struct {
	Avatar       string `json:"avatar"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Reward       int    `json:"reward"`
	SerialNumber int    `json:"serialNumber"`
	Steamid      string `json:"steamid"`
	Ticket       int    `json:"ticket"`
}
type SupplyDropWinnerData struct {
	Avatars  []string `json:"avatars"`
	TimeLeft int      `json:"timeLeft"`
	Winner   Winner   `json:"winner"`
}

type MinesWinner struct {
	Room string          `json:"room"`
	Type string          `json:"type"`
	Data MinesWinnerData `json:"data"`
}
type MinesWinnerData struct {
	Lobby  int    `json:"lobby"`
	Map    string `json:"map"`
	Secret string `json:"secret"`
	Seed   int64  `json:"seed"`
	Winner int    `json:"winner"`
}

type MinesNewPlayer struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data MinesNewPlayerData `json:"data"`
}
type Player struct {
	Avatar  string `json:"a"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	Bet     int    `json:"o"`
	SteamID string `json:"s"`
}
type MinesNewPlayerData struct {
	Lobby  int    `json:"lobby"`
	Player Player `json:"player"`
	Pot    int    `json:"pot"`
}

type MinesNewGame struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data MinesNewGameData `json:"data"`
}
type Players struct {
	Avatar  string `json:"a"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	Bet     int    `json:"o"`
	SteamID string `json:"s"`
	Empty   bool   `json:"empty,omitempty"`
}
type MinesNewGameData struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
}

type MinesList struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data MinesListData `json:"data"`
}

type InProgress struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
	Timer        int       `json:"timer,omitempty"`
}

type Joinable struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	Timer        int       `json:"timer"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
}

type MinesListSettings struct {
	Enabled  bool `json:"enabled"`
	MaxValue int  `json:"maxValue"`
	MinValue int  `json:"minValue"`
}

type MinesListData struct {
	InProgress []InProgress      `json:"inProgress"`
	Joinable   []Joinable        `json:"joinable"`
	Settings   MinesListSettings `json:"settings"`
}

type MinesGameStarting struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data MinesGameStartingData `json:"data"`
}

type MinesGameStartingData struct {
	Lobby         int            `json:"lobby"`
	NumberOfBombs int            `json:"numberOfBombs"`
	NumberOfTiles int            `json:"numberOfTiles"`
	PlayerOrder   map[string]int `json:"playerOrder"`
	Time          int            `json:"time"`
}

type MinesGameStarted struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type MinesBeginTimer struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data MinesBeginTimerData `json:"data"`
}
type MinesBeginTimerData struct {
	Lobby int `json:"lobby"`
	Timer int `json:"timer"`
}

type LowJackpotStartTimer struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type LowJackpotNewGame struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data LowJackpotNewGameData `json:"data"`
}

type LowJackpotNewGameData struct {
	NewGame NewGame `json:"newGame"`
	OldGame OldGame `json:"oldGame"`
}

type NewGame struct {
	Deposits interface{} `json:"deposits"`
	Expires  int         `json:"expires"`
	Hash     string      `json:"hash"`
	ID       int         `json:"id"`
}

type OldGame struct {
	Deposits     []Deposits
	Expires      int    `json:"expires"`
	Hash         string `json:"hash"`
	ID           int    `json:"id"`
	Mod          string `json:"mod"`
	Percentage   string `json:"percentage"`
	Secret       string `json:"secret"`
	Seed         string `json:"seed"`
	SerialNumber int    `json:"serialNumber"`
	TicketNumber int    `json:"ticketNumber"`
	Winner       string `json:"winner"`
}
type Deposits struct {
	Avatar  string  `json:"avatar"`
	Color   string  `json:"color"`
	ID      int     `json:"id"`
	Items   [][]int `json:"items"`
	Level   int     `json:"level"`
	Name    string  `json:"name"`
	SteamID string  `json:"steamid"`
	UserID  int     `json:"user_id"`
	Value   int     `json:"value"`
}

type LowJackpotNewDeposit struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data Deposits `json:"data"`
}

type LowJackpotList struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data LowJackpotListData `json:"data"`
}

type LowJackpotListData struct {
	Current  Current   `json:"current"`
	History  []History `json:"history"`
	Rolling  bool      `json:"rolling"`
	Settings Settings  `json:"settings"`
}

type Settings struct {
	CasinoPercentage       int  `json:"casinoPercentage"`
	Disabled               bool `json:"disabled"`
	GameMaxItems           int  `json:"gameMaxItems"`
	GameRoundTime          int  `json:"gameRoundTime"`
	MinItemValue           int  `json:"minItemValue"`
	NameDiscount           int  `json:"nameDiscount"`
	SecondCasinoPercentage int  `json:"secondCasinoPercentage"`
	UserMaxDeposits        int  `json:"userMaxDeposits"`
	UserMaxItems           int  `json:"userMaxItems"`
	UserMaxValue           int  `json:"userMaxValue"`
	UserMinItems           int  `json:"userMinItems"`
	UserMinValue           int  `json:"userMinValue"`
}

type History struct {
	Deposits     []Deposits `json:"deposits"`
	Expires      int        `json:"expires"`
	Hash         string     `json:"hash"`
	ID           int        `json:"id"`
	Mod          string     `json:"mod"`
	Percentage   string     `json:"percentage"`
	Secret       string     `json:"secret"`
	Seed         string     `json:"seed"`
	SerialNumber int        `json:"serialNumber"`
	TicketNumber int        `json:"ticketNumber"`
	Winner       string     `json:"winner"`
}

type Current struct {
	Deposits []Deposits `json:"deposits"`
	Expires  int        `json:"expires"`
	Hash     string     `json:"hash"`
	ID       int        `json:"id"`
}

type JackpotStartTimer struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type JackpotNewGame struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data LowJackpotNewGameData `json:"data"`
}

type JackpotNewDeposit struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data Deposits `json:"data"`
}

type JackpotList struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data LowJackpotListData `json:"data"`
}

type CrashTick struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type CrashStart struct {
	Room string         `json:"room"`
	Type string         `json:"type"`
	Data CrashStartData `json:"data"`
}

type CrashStartData struct {
	State     int    `json:"state"`
	TimeStart string `json:"timeStart"`
}

type CrashNew struct {
	Room string       `json:"room"`
	Type string       `json:"type"`
	Data CrashNewData `json:"data"`
}

type CrashNewData struct {
	Bets      interface{} `json:"bets"`
	Elapsed   int         `json:"elapsed"`
	ID        int         `json:"id"`
	State     int         `json:"state"`
	TimeStart string      `json:"timeStart"`
	Timer     int         `json:"timer"`
}

type CrashMultipleBets struct {
	Room string                  `json:"room"`
	Type string                  `json:"type"`
	Data []CrashMultipleBetsData `json:"data"`
}

type CrashMultipleBetsData struct {
	Avatar  string `json:"a"`
	Bet     int    `json:"f"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	SteamID string `json:"s"`
	UserID  int    `json:"u`
}

type CrashList struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data CrashListData `json:"data"`
}

type Game struct {
	Bets      []CrashMultipleBetsData `json:"bets"`
	Elapsed   int                     `json:"elapsed"`
	ID        int                     `json:"id"`
	State     int                     `json:"state"`
	TimeStart time.Time               `json:"timeStart"`
	Timer     int                     `json:"timer"`
}
type CrashListDataHistory struct {
	CrashPoint float64 `json:"crashPoint"`
	ID         int     `json:"id"`
}
type CrashListDataSettings struct {
	Disabled bool `json:"disabled"`
	MaxValue int  `json:"maxValue"`
	MaxWin   int  `json:"maxWin"`
	MinValue int  `json:"minValue"`
}
type CrashListData struct {
	Game     Game                   `json:"game"`
	History  []CrashListDataHistory `json:"history"`
	Settings CrashListDataSettings  `json:"settings"`
}

type CrashEnd struct {
	Room string       `json:"room"`
	Type string       `json:"type"`
	Data CrashEndData `json:"data"`
}
type CrashEndData struct {
	CrashPoint float64 `json:"crashPoint"`
	ID         int     `json:"id"`
	State      int     `json:"state"`
	Timer      int     `json:"timer"`
}

type CrashCashOut struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data CrashCashOutData `json:"data"`
}

type CrashCashOutData struct {
	Amount    int     `json:"amount"`
	CashoutAt float64 `json:"cashoutAt"`
	ID        int     `json:"id"`
}

type CoinflipUpdateGame struct {
	Room string                 `json:"room"`
	Type string                 `json:"type"`
	Data CoinflipUpdateGameData `json:"data"`
}

type CoinflipUpdateGameData struct {
	BlueSide     Side   `json:"blue_side"`
	Diff         int    `json:"diff"`
	Hash         string `json:"hash"`
	ID           int    `json:"id"`
	InitialValue int    `json:"initial_value"`
	Owner        string `json:"owner"`
	RedSide      Side   `json:"red_side"`
	Status       string `json:"status"`
	TimeLeft     int    `json:"time_left"`
	Timer        int    `json:"timer"`
	Value        int    `json:"value"`
}

type Side struct {
	Avatar  string  `json:"avatar"`
	ID      int     `json:"id"`
	Items   [][]int `json:"items"`
	Level   int     `json:"level"`
	Name    string  `json:"name"`
	Steamid string  `json:"steamid"`
}

type CoinflipNewGame struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data CoinflipNewGameData `json:"data"`
}
type CoinflipNewGameData struct {
	Diff         int    `json:"diff"`
	Hash         string `json:"hash"`
	ID           int    `json:"id"`
	InitialValue int    `json:"initial_value"`
	Owner        string `json:"owner"`
	RedSide      Side   `json:"red_side,omniempty"`
	BlueSide     Side   `json:"blue_side,omniempty"`
	Status       string `json:"status"`
	TimeLeft     int    `json:"time_left"`
	Value        int    `json:"value"`
}

type CoinflipList struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data CoinflipListData `json:"data"`
}

type CoinflipListData struct {
	Games []CoinflipNewGameData `json:"games"`
}

type CoinflipGameStatus struct {
	Room string                 `json:"room"`
	Type string                 `json:"type"`
	Data CoinflipGameStatusData `json:"data"`
}

type CoinflipGameStatusData struct {
	ID           int    `json:"id"`
	RedSide      Side   `json:"red_side,omniempty"`
	BlueSide     Side   `json:"blue_side,omniempty"`
	Status       string `json:"status"`
	Timer        int    `json:"timer"`
	Mod          string `json:"mod"`
	Secret       string `json:"secret"`
	Seed         string `json:"seed"`
	SerialNumber int    `json:"serialNumber"`
	TicketNumber int    `json:"ticketNumber"`
	WinnerSide   string `json:"winner_side"`
}

type CoinflipDeleteGame struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

type ChatStats struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data ChatStatsData `json:"data"`
}
type ChatStatsData struct {
	Online      int `json:"online"`
	SteamStatus int `json:"steamStatus"`
}

type ChatRooms struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data []string `json:"data"`
}

type ChatMessage struct {
	Room string          `json:"room"`
	Type string          `json:"type"`
	Data ChatMessageData `json:"data"`
}
type ChatMessageProfile struct {
	UnderscoreID string `json:"_id"`
	ID           int    `json:"id"`
	Steamid      string `json:"steamid"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Rank         int    `json:"rank"`
	Level        int    `json:"level"`
}
type ChatMessageData struct {
	Profile ChatMessageProfile `json:"profile"`
	Content string             `json:"content"`
	ID      string             `json:"id"`
	Time    int                `json:"time"`
}
