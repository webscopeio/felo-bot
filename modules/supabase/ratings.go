package supabase

type Rating struct {
	Id string `json:"id" db:"id"`
	Game string `json:"game" db:"game"`
	Elo int `json:"elo" db:"elo"`
	Streak int `json:"streak" db:"streak"`
	Played int `json:"played" db:"played"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}