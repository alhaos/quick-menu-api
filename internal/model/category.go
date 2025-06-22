package model

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ClientID    string `json:"clientID"`
	Description string `json:"description"`
}
