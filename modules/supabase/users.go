package supabase

import "encoding/json"

type User struct {
	Id        string `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Name      string `json:"name" db:"name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func GetAllUsers(db *DB) (*QueryResult[User], error) {
	var users []User
	data, count, err := db.From("users").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &users)
	return &QueryResult[User]{
		data:  users,
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
