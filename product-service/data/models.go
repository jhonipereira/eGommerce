package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type PostgresRepository struct {
	Conn *sql.DB
}

func NewPostgresRepository(pool *sql.DB) *PostgresRepository {
	db = pool
	return &PostgresRepository{
		Conn: pool,
	}
}

// Product is the structure which holds one product from the database.
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Photos      string    `json:"photos,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GetAll returns a slice of all products, sorted by name
func (u *PostgresRepository) GetAll() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, photos, created_at, updated_at
	from products order by name`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Photos,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

// GetByName returns one product by name
func (u *PostgresRepository) GetByName(name string) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, photos, created_at, updated_at from products where name = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, name)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Photos,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// GetOne returns one product by id
func (u *PostgresRepository) GetOne(id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, photos, created_at, updated_at from products where id = $1`

	var product Product
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Photos,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// updates one product in the database, using the information
// stored in the receiver u
func (u *PostgresRepository) Update(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update products set
		name = $1,
		description = $2,
		photos = $3,
		updated_at = $4
		where id = $5
	`

	_, err := db.ExecContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Photos,
		time.Now(),
		product.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one product from the database, by ID
func (u *PostgresRepository) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from products where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new product into the database, and returns the ID of the newly inserted row
func (u *PostgresRepository) Insert(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into products (name, description, photos, created_at, updated_at)
		values ($1, $2, $3, $4, $5) returning id`

	err := db.QueryRowContext(ctx, stmt,
		product.Name,
		product.Description,
		product.Photos,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}
