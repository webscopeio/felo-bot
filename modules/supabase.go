package modules

import (
	"encoding/json"

	"github.com/supabase-community/supabase-go"
)

type Supabase struct {
	SUPABASE_KEY string
	SUPABASE_URL string
}

func (s *Supabase) Init(options *supabase.ClientOptions) (*supabase.Client, error) {
	client, err := supabase.NewClient(s.SUPABASE_URL, s.SUPABASE_KEY, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type DB = supabase.Client

type User struct {
	Id   string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Name string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Rating struct {
	Id string `json:"id" db:"id"`
	Game string `json:"game" db:"game"`
	Elo int `json:"elo" db:"elo"`
	Streak int `json:"streak" db:"streak"`
	Played int `json:"played" db:"played"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type Match struct {
	Id string `json:"id" db:"id"`
	PlayersHome []string `json:"players_home" db:"players_home"`
	PlayersAway []string `json:"players_away" db:"players_away"`
	ScoreHome int `json:"score_home" db:"score_home"`
	ScoreAway int `json:"score_away" db:"score_away"`
	CreatedAt string `json:"created_at" db:"created_at"`
}


type QueryResult[T User | Rating | Match] struct {
	data []T
	count int64
}

func GetAllUsers(db *DB) (*QueryResult[User], error) {
	var users []User
	data, count, err := db.From("users").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &users)
	return &QueryResult[User]{
		data: users,
		count: count,
	}, nil
}

func GetUserById(db *DB, id string) (*User, error) {
	var user User
	data, _, err := db.From("users").Select("*", "exact", false).Eq("id", id).Execute()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &user)
	return &user, nil
} 
