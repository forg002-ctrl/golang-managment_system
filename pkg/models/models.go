package models

func newAccount(id int, username string) *Account {
	return &Account{
		ID: id,
		Username: username,
	};
}

type Account struct {
	ID 	int `json:"id"`
	Username string `json:"username"` 
}