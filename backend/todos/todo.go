package todos

type Todo struct {
	ID uint64  `json:"id"`
	Description string `json:"description"`
	UserEmail string `json:"useremail"`
	Complete bool `json:"complete"`
}

