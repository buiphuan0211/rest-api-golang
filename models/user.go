package models

type (
	// PlayerBSON ...
	PlayerBSON struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)
