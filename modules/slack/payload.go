package slack

type Block struct {
	BlockId string `json:"action_id"`
}

var t = Block{
	BlockId: "test",
}