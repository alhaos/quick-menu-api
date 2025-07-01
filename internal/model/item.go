package model

type Item struct {
	ID          string `json:"ID"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Price       string `json:"price,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
}

func MockItem() Item {
	return Item{
		Name:        "Шашлык",
		Description: "Очень вкусный шашлык ",
		Image:       "image.jpg",
		IsActive:    false,
	}
}
