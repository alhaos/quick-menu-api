package model

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func MockCategory() Category {
	return Category{
		Name:        "Шаурма",
		Description: "Разная шаурма",
		Image:       "",
	}
}
