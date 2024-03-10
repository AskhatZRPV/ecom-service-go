package category

type Category struct {
	ID          int
	Title       string
	Description string
}

func New(title string, description string) *Category {
	return &Category{
		Title:       title,
		Description: description,
	}
}

func Update(id int, title string, description string) *Category {
	return &Category{
		ID:          id,
		Title:       title,
		Description: description,
	}
}
