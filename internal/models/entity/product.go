package entity

type Product struct {
	Id       int32    `json:"id" db:"id"`
	Category int32    `json:"category" db:"category"`
	Name     string   `json:"name" db:"name"`
	Color    string   `json:"color" db:"color"`
	Price    int32    `json:"price" db:"price"`
	Size     int32    `json:"size" db:"size"`
	Quantity int32    `json:"quantity" db:"quantity"`
	Images   []string `json:"images" db:"images"`
}
