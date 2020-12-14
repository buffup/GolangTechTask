package handlers

type CreateBuff struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type Buff struct {
	ID       uint64   `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}
