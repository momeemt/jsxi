package main

type (
	Post struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		Description string `json:"description"`
		UserId      int    `json:"user_id"`
	}

	Response struct {
		Posts []Post `json:"posts"`
	}
)
