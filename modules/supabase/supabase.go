package supabase

import "github.com/supabase-community/supabase-go"

type Client struct {
	SUPABASE_KEY string
	SUPABASE_URL string
}

func (s *Client) Init(options *supabase.ClientOptions) (*supabase.Client, error) {
	client, err := supabase.NewClient(s.SUPABASE_URL, s.SUPABASE_KEY, options)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type DB = supabase.Client

type QueryResult[T User | Rating | Match] struct {
	data []T
	count int64
}