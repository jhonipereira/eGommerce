package data

import (
	"database/sql"
	"time"
)

type PostgresTestRepository struct {
	Conn *sql.DB
}

func NewPostgresTestRepository(db *sql.DB) *PostgresTestRepository {
	return &PostgresTestRepository{Conn: db}
}

func (u *PostgresTestRepository) GetAll() ([]*Product, error) {
	products := []*Product{}

	return products, nil
}

func (u *PostgresTestRepository) GetByName(name string) (*Product, error) {
	product := Product{
		ID:          1,
		Name:        "very nice product",
		Description: "lorem ipsum",
		Photos:      "https://www.shutterstock.com/pt/image-vector/happy-thanksgiving-retro-greeting-card-ribbon-1828500257",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &product, nil
}

func (u *PostgresTestRepository) GetOne(id int) (*Product, error) {
	product := Product{
		ID:          1,
		Name:        "very nice product for black friday",
		Description: "happy thanksgiving",
		Photos:      "https://www.shutterstock.com/pt/image-vector/happy-thanksgiving-retro-greeting-card-ribbon-1828500257",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return &product, nil
}

func (u *PostgresTestRepository) Update(product Product) error {

	return nil
}

func (u *PostgresTestRepository) DeleteByID(id int) error {

	return nil
}

func (u *PostgresTestRepository) Insert(product Product) (int, error) {

	return 3, nil
}
