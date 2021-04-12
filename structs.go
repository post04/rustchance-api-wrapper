package wrapper

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Session is the main struct for the session
type Session struct {
	// Auth is the auth token for an account, if this is blank the socket will be unauthorized and any account specific actions will return an error
	Auth string
	// Socket is the main socket, if you want to interact with the Socket directly this is what you would use
	Socket *websocket.Conn
	// SocketMutex is the socket mutex to stop concurrent writing
	SocketMutex sync.Mutex
	// Handlers is a map of handlers where string is the room_type and the func is added by the AddHandler func
	Handlers map[string]func(*Session, interface{})
	// Headers is used when connecting to the socket (maybe http requests) and are set automatically, incase you want to set them yourself though the option is always there
	Headers http.Header
	// Rooms is what rooms we want to listen for on the socket, by default it's []string{"chat", "crash", "shop", "coinflip", "jackpot", "jackpot-low", "supply-drops", "mines"}
	Rooms []string
	// Room is the chat room we join, default is "en" but it can also be "tr" or "ru"
	Room string
	// Log is logging errors to console, this is defaulted as false
	Log bool
}

// Payload is the typical payload, this should be able to be used 99% of the time when writing to the socket
type Payload struct {
	Data interface{} `json:"data"`
	Room string      `json:"room"`
	Type string      `json:"type"`
}

// ShopRules is the expected payload from the shop_rules socket event
type ShopRules struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data ShopRulesData `json:"data"`
}

// ShopRulesData is the expected data field of the shop_rules socket event
type ShopRulesData struct {
	Enabled      bool `json:"enabled"`
	MaxItems     int  `json:"maxItems"`
	MinItemValue int  `json:"minItemValue"`
	MinItems     int  `json:"minItems"`
	MinValue     int  `json:"minValue"`
}

// SupplyDropsJoinable is the expected payload from the supply-drops_joinable socket event
type SupplyDropsJoinable struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// SupplyDropsList is the expected payload from the supply-drops_list socket event
type SupplyDropsList struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data SupplyDropsListData `json:"data"`
}

// SupplyDropsListData is the expected data field from the supply-drops_list socket event
type SupplyDropsListData struct {
	ID         int  `json:"id"`
	Joined     bool `json:"joined"`
	Players    int  `json:"players"`
	Processing bool `json:"processing"`
	State      int  `json:"state"`
	TimeLeft   int  `json:"timeLeft"`
}

// SupplyDropsPlayers says how many players are currently in the supply drop
type SupplyDropsPlayers struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// SupplyDropWinner says who won a supply drop, this contains all sorts of data about the user who won including their steam ID and steam avatar
type SupplyDropWinner struct {
	Room string               `json:"room"`
	Type string               `json:"type"`
	Data SupplyDropWinnerData `json:"data"`
}

// Winner is the winner of a supply drop
type Winner struct {
	Avatar       string `json:"avatar"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Reward       int    `json:"reward"`
	SerialNumber int    `json:"serialNumber"`
	Steamid      string `json:"steamid"`
	Ticket       int    `json:"ticket"`
}

// SupplyDropWinnerData contains information about the supply drop winner including timeleft (idk why this is here, seems to always be 0) and the winner
type SupplyDropWinnerData struct {
	Avatars  []string `json:"avatars"`
	TimeLeft int      `json:"timeLeft"`
	Winner   Winner   `json:"winner"`
}

// MinesWinner represents a win in the mines gamemode
type MinesWinner struct {
	Room string          `json:"room"`
	Type string          `json:"type"`
	Data MinesWinnerData `json:"data"`
}

// MinesWinnerData is the winner of a mines game, the Winner field is probably the rustchance.com player id (rustchance has 0 documentation so I can't be sure)
type MinesWinnerData struct {
	Lobby  int    `json:"lobby"`
	Map    string `json:"map"`
	Secret string `json:"secret"`
	Seed   int64  `json:"seed"`
	Winner int    `json:"winner"`
}

// MinesNewPlayer represents a new player joining a mines game
type MinesNewPlayer struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data MinesNewPlayerData `json:"data"`
}

// Player is the struct that describes account information about the Player joining a mines game
// NOTE: Because of the lack of documentation and the terrible variable names in the json payload, I could have gotten some of this wrong
type Player struct {
	Avatar  string `json:"a"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	Bet     int    `json:"o"`
	SteamID string `json:"s"`
}

// MinesNewPlayerData is the data for a mines game when a new player joins, as well as the updated pot and the player data such as name and avatar as well as other things
type MinesNewPlayerData struct {
	Lobby  int    `json:"lobby"`
	Player Player `json:"player"`
	Pot    int    `json:"pot"`
}

// MinesNewGame is the payload for when a new mines game is made, includes stuff like the current players and how much the mines game is worth
type MinesNewGame struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data MinesNewGameData `json:"data"`
}

// Players is used in a few places but it's similar to the Player struct except it has a field Empty, the field empty basically means "is the field filled by a player or is this slot open"
type Players struct {
	Avatar  string `json:"a"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	Bet     int    `json:"o"`
	SteamID string `json:"s"`
	Empty   bool   `json:"empty,omitempty"`
}

// MinesNewGameData is the data for a new mines game, includes how much the game is worth and the game id
type MinesNewGameData struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
}

// MinesList is a list of mine games, this should only be triggered once on startup but due to lack of documentation I can't say for sure
type MinesList struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data MinesListData `json:"data"`
}

// InProgress is used for a mines game that's in progress
type InProgress struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
	Timer        int       `json:"timer,omitempty"`
}

// Joinable is used for a mines game that's joinable
type Joinable struct {
	ID           int       `json:"id"`
	JoinValue    int       `json:"joinValue"`
	Players      []Players `json:"players"`
	State        int       `json:"state"`
	Timer        int       `json:"timer"`
	TotalPlayers int       `json:"totalPlayers"`
	TotalPot     int       `json:"totalPot"`
}

// MinesListSettings is the settings for a mines game
type MinesListSettings struct {
	Enabled  bool `json:"enabled"`
	MaxValue int  `json:"maxValue"`
	MinValue int  `json:"minValue"`
}

// MinesListData is the data for a mines list, has games in progess and games that are joinable currently
type MinesListData struct {
	InProgress []InProgress      `json:"inProgress"`
	Joinable   []Joinable        `json:"joinable"`
	Settings   MinesListSettings `json:"settings"`
}

// MinesGameStarting is when a mines game is starting
type MinesGameStarting struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data MinesGameStartingData `json:"data"`
}

// MinesGameStartingData is the data for a mines game that's starting, includes the game id and number of bombs and number of tiles
type MinesGameStartingData struct {
	Lobby         int            `json:"lobby"`
	NumberOfBombs int            `json:"numberOfBombs"`
	NumberOfTiles int            `json:"numberOfTiles"`
	PlayerOrder   map[string]int `json:"playerOrder"`
	Time          int            `json:"time"`
}

// MinesGameStarted is used to tell when a mines game is started, the data is the game id
type MinesGameStarted struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// MinesBeginTimer is used for telling when a game is going to start (I think, no documentation makes it hard)
type MinesBeginTimer struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data MinesBeginTimerData `json:"data"`
}

// MinesBeginTimerData contains the game id and the time that the game is going to start (probably Unix but with no documentation I can't say for sure)
type MinesBeginTimerData struct {
	Lobby int `json:"lobby"`
	Timer int `json:"timer"`
}

// LowJackpotStartTimer is used to tell when a low jackpot timer is starting, I believe the data is the Unix time when the timer is done and the lowjackpot is rolled but with no documentation it's hard to tell
type LowJackpotStartTimer struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// LowJackpotNewGame contains data for a new Lowjackpot game starting
type LowJackpotNewGame struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data LowJackpotNewGameData `json:"data"`
}

// LowJackpotNewGameData contains data about the previous(old) game and the current(new) game
type LowJackpotNewGameData struct {
	NewGame NewGame `json:"newGame"`
	OldGame OldGame `json:"oldGame"`
}

// NewGame is the data for the current(new) lowjackpot game
type NewGame struct {
	Deposits interface{} `json:"deposits"`
	Expires  int         `json:"expires"`
	Hash     string      `json:"hash"`
	ID       int         `json:"id"`
}

// OldGame contains data for the previous(old) lowjackpot game
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

// Deposits contains information for each deposit including what items they deposited and infromation about their account. This is used in a different structs
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

// LowJackpotNewDeposit contains information about when a new deposit is added to a lowjackpot game
type LowJackpotNewDeposit struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data Deposits `json:"data"`
}

// LowJackpotList is a list of low jackpot games including the current active game and previous games, this should only be triggered once on startup but with lack of documentation I can't know for sure
type LowJackpotList struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data LowJackpotListData `json:"data"`
}

// LowJackpotListData contains data like previous low jackpot games, the current game, and if the current game is rolling
type LowJackpotListData struct {
	Current  Current   `json:"current"`
	History  []History `json:"history"`
	Rolling  bool      `json:"rolling"`
	Settings Settings  `json:"settings"`
}

// Settings contains information about the Low jackpot settings, there's a lot of stuff here but most of it speaks for it self
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

// History is data for a previous low jackpot game, there's a lot of data here but most of it speaks for it self
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

// Current is the data for the current/active low jackpot game
type Current struct {
	Deposits []Deposits `json:"deposits"`
	Expires  int        `json:"expires"`
	Hash     string     `json:"hash"`
	ID       int        `json:"id"`
}

// JackpotStartTimer is used to tell that the current highrollers jackpot game timer is starting, the data should be the Unix time that the timer is done and the game gets rolled (with lack of documentation it's hard to tell)
type JackpotStartTimer struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// JackpotNewGame contains data about a new highrollers jackpot game
type JackpotNewGame struct {
	Room string                `json:"room"`
	Type string                `json:"type"`
	Data LowJackpotNewGameData `json:"data"`
}

// JackpotNewDeposit contains data about a new highrollers jackpot deposit of skins, includes information about the user adding skins and the skins being added
type JackpotNewDeposit struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data Deposits `json:"data"`
}

// JackpotList is the list of highroller jackpot games, should only be triggered once on startup but due to lack of documentation I cannot tell for sure
type JackpotList struct {
	Room string             `json:"room"`
	Type string             `json:"type"`
	Data LowJackpotListData `json:"data"`
}

// CrashTick is a new tick of crash, it's the current crash position. I'm unsure how to actually use this data but it's there
type CrashTick struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// CrashStart has information about a new crash game starting
type CrashStart struct {
	Room string         `json:"room"`
	Type string         `json:"type"`
	Data CrashStartData `json:"data"`
}

// CrashStartData has information about the state of the crash game and the time the game started (which is a string for god knows what reason)
type CrashStartData struct {
	State     int    `json:"state"`
	TimeStart string `json:"timeStart"`
}

// CrashNew is a new crash game, it contains infromation about a new crash game but over all it's really just useful for knowing a new game start and not really useful for the information in the payload
type CrashNew struct {
	Room string       `json:"room"`
	Type string       `json:"type"`
	Data CrashNewData `json:"data"`
}

// CrashNewData has data about a new crash game being started, bets *should* always be nil
type CrashNewData struct {
	Bets      interface{} `json:"bets"`
	Elapsed   int         `json:"elapsed"`
	ID        int         `json:"id"`
	State     int         `json:"state"`
	TimeStart string      `json:"timeStart"`
	Timer     int         `json:"timer"`
}

// CrashMultipleBets is used when a new bet(s) are added to the current crash game
type CrashMultipleBets struct {
	Room string                  `json:"room"`
	Type string                  `json:"type"`
	Data []CrashMultipleBetsData `json:"data"`
}

// CrashMultipleBetsData contains data about the bet being added
// NOTE: Some of this data may be incorrect, rustchance has no documentation so I had to reverse everything my self so some things may be wrong
type CrashMultipleBetsData struct {
	Avatar  string `json:"a"`
	Bet     int    `json:"f"`
	ID      int    `json:"i"`
	Level   int    `json:"l"`
	Name    string `json:"n"`
	SteamID string `json:"s"`
	UserID  int    `json:"u"`
}

// CrashList is a list of crash games, should only be triggered once but due to lack of documentation I cannot say for sure
type CrashList struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data CrashListData `json:"data"`
}

// Game is the current crash game (used in the CrashList payload)
type Game struct {
	Bets      []CrashMultipleBetsData `json:"bets"`
	Elapsed   int                     `json:"elapsed"`
	ID        int                     `json:"id"`
	State     int                     `json:"state"`
	TimeStart time.Time               `json:"timeStart"`
	Timer     int                     `json:"timer"`
}

// CrashListDataHistory contains information about previous games of crash, the game id and crashpoint
type CrashListDataHistory struct {
	CrashPoint float64 `json:"crashPoint"`
	ID         int     `json:"id"`
}

// CrashListDataSettings are the settings for the current crash
type CrashListDataSettings struct {
	Disabled bool `json:"disabled"`
	MaxValue int  `json:"maxValue"`
	MaxWin   int  `json:"maxWin"`
	MinValue int  `json:"minValue"`
}

// CrashListData is the data for the CrashList payload, includes previous crash games and the current active crash game
type CrashListData struct {
	Game     Game                   `json:"game"`
	History  []CrashListDataHistory `json:"history"`
	Settings CrashListDataSettings  `json:"settings"`
}

// CrashEnd is used when a crash game is done, includes information about the game
type CrashEnd struct {
	Room string       `json:"room"`
	Type string       `json:"type"`
	Data CrashEndData `json:"data"`
}

// CrashEndData contains information about the ended crash game like the game id and crash point
type CrashEndData struct {
	CrashPoint float64 `json:"crashPoint"`
	ID         int     `json:"id"`
	State      int     `json:"state"`
	Timer      int     `json:"timer"`
}

// CrashCashOut is used when someone cashes out of a crash game, taking their profit and leaving
type CrashCashOut struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data CrashCashOutData `json:"data"`
}

// CrashCashOutData contains information like how much money the person leaving made and I think their rustchance user id. The Amount isn't their profit but instead their bet + their profit. You can reverse this by dividing the Amount by the CrashPoint as that's the reverse of how they calculate this value
type CrashCashOutData struct {
	Amount    int     `json:"amount"`
	CashoutAt float64 `json:"cashoutAt"`
	ID        int     `json:"id"`
}

// CoinflipUpdateGame is used for when a coinflip game is updating status
type CoinflipUpdateGame struct {
	Room string                 `json:"room"`
	Type string                 `json:"type"`
	Data CoinflipUpdateGameData `json:"data"`
}

// CoinflipUpdateGameData is the data for when a coinflip game updates, rather that be a user joining or the game finishing
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

// Side represents either red or blue side on a coinflip match
type Side struct {
	Avatar  string  `json:"avatar"`
	ID      int     `json:"id"`
	Items   [][]int `json:"items"`
	Level   int     `json:"level"`
	Name    string  `json:"name"`
	Steamid string  `json:"steamid"`
}

// CoinflipNewGame is the event for coinflip_newgame
type CoinflipNewGame struct {
	Room string              `json:"room"`
	Type string              `json:"type"`
	Data CoinflipNewGameData `json:"data"`
}

// CoinflipNewGameData is the data for a new game of coinflip, this contains the game ID, how much money is on the line, etc.
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

// CoinflipList is a list of coinflip games, finished or not. This event should only fire on startup once but with no real documentation from rustchance there's no way to tell except time
type CoinflipList struct {
	Room string           `json:"room"`
	Type string           `json:"type"`
	Data CoinflipListData `json:"data"`
}

// CoinflipListData is a slice of CoinflipNewGameData because the data there is also applicable to this event
type CoinflipListData struct {
	Games []CoinflipNewGameData `json:"games"`
}

// CoinflipGameStatus updates the game status for a coinflip game, the data is CoinflipGameStatusData
type CoinflipGameStatus struct {
	Room string                 `json:"room"`
	Type string                 `json:"type"`
	Data CoinflipGameStatusData `json:"data"`
}

// CoinflipGameStatusData is the data for a coinflip game, this includes the ID of the game, serial/seed information, and which side won as well as if the game is finished
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

// CoinflipDeleteGame is the event for when a coinflip game is deleted, the Data is the ID for the game
type CoinflipDeleteGame struct {
	Room string `json:"room"`
	Type string `json:"type"`
	Data int    `json:"data"`
}

// ChatStats is the event for chat_stats in the socket
type ChatStats struct {
	Room string        `json:"room"`
	Type string        `json:"type"`
	Data ChatStatsData `json:"data"`
}

// ChatStatsData includes the steam stability status and how many users are online in the current chat
type ChatStatsData struct {
	Online      int `json:"online"`
	SteamStatus int `json:"steamStatus"`
}

// ChatRooms is the struct for the event we use to see when to set our current room, when this event fires you should chance your room to "en", "tr", or "ru". The data is the avalible rooms we can join
type ChatRooms struct {
	Room string   `json:"room"`
	Type string   `json:"type"`
	Data []string `json:"data"`
}

// ChatMessage is the data for the chat_message event through the socket
type ChatMessage struct {
	Room string          `json:"room"`
	Type string          `json:"type"`
	Data ChatMessageData `json:"data"`
}

// ChatMessageProfile is the profile that is making a chat message, this includes things like steamID, steam avatar, rust chance id, etc.
type ChatMessageProfile struct {
	UnderscoreID string `json:"_id"`
	ID           int    `json:"id"`
	Steamid      string `json:"steamid"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Rank         int    `json:"rank"`
	Level        int    `json:"level"`
}

// ChatMessageData is the expected data field for the chat_message socket event
type ChatMessageData struct {
	Profile ChatMessageProfile `json:"profile"`
	Content string             `json:"content"`
	ID      string             `json:"id"`
	Time    int                `json:"time"`
}

// AccountLeaderboard is the information about an account leaderboard
type AccountLeaderboard struct {
	Ranked   bool `json:"ranked"`
	Tickets  int  `json:"tickets"`
	Position int  `json:"position"`
}

// TicketsLeaderboard is the data for the current users in the tickets leaderboard
type TicketsLeaderboard struct {
	Success bool                     `json:"success"`
	Result  TicketsLeaderboardResult `json:"result"`
}

// User represents a user in the ticket leaderboard
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Tickets  int    `json:"tickets"`
}

// TicketsLeaderboardResult contains data for the tickets leaderboard
type TicketsLeaderboardResult struct {
	Users   []User      `json:"leaderboard"`
	Rewards interface{} `json:"rewards"`
}
