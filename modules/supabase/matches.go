package supabase

type Match struct {
	Id string `json:"id" db:"id"`
	PlayersHome []string `json:"players_home" db:"players_home"`
	PlayersAway []string `json:"players_away" db:"players_away"`
	ScoreHome int `json:"score_home" db:"score_home"`
	ScoreAway int `json:"score_away" db:"score_away"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
