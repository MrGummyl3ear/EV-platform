package repository

type Repository struct {
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
