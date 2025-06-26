package model

type Item struct {
	ID            string
	Name          string
	Description   string
	ImageFilename string
	IsActive      bool
}

func MockItem() Item {
	return Item{
		Name:          "Шашлык",
		Description:   "Очень вкусный шашлык ",
		ImageFilename: "image.jpg",
		IsActive:      false,
	}
}
