package players

// this will be mongodb model for player
type Player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Players struct {
	Players []Player `json:"players"`
}
