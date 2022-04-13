package models

type (
	// PlayerBSON ...
	User struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)
