package repository

// Repository consists bunch of different data storage
// that the project might be using

type Repository struct {
}

// New: Instantiation of all the repository

func New() *Repository {
	return &Repository{}
}
