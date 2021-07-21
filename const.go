package wrapper

const (
	// SocketURL Is the url for the websocket that rustchance.com uses
	SocketURL = "wss://rustchance.com/feed"
	// AccountLeaderboardURL is the url for the account leader board
	AccountLeaderboardURL = "https://rustchance.com/api/account/leaderboard"
	// TicketsLeaderboardURL is the url for the tickets leaderboard data
	TicketsLeaderboardURL = "https://rustchance.com/api/bonuses"
	// AccountEarningsURL is the url to get the total account earning of an account
	AccountEarningsURL = "https://rustchance.com/api/account/stats/all"
	// AccountProfileURL is the url to get the account profile, right now we use this to get the json of user information
	AccountProfileURL = "https://rustchance.com/profile"
	// FaucetClaimURL is the url to claim the free 3 cent faucet
	FaucetClaimURL = "https://rustchance.com/api/account/faucet"
	// ProvefairSerialURL is the url to check the validity of a "provably fair" action
	ProvefairSerialURL = "https://rustchance.com/api/serial/"
)
