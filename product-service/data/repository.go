package data

type Repository interface {
	GetAll() ([]*Product, error)
	GetByName(name string) (*Product, error)
	GetOne(id int) (*Product, error)
	Update(product Product) error
	DeleteByID(id int) error
	Insert(product Product) (int, error)
}
