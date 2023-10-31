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

// GetAll returns a slice of all users, sorted by last name
func (u *PostgresTestRepository) GetAll() ([]*User, error) {
	users := []*User{}

	return users, nil
}

// GetByEmail returns one user by email
func (u *PostgresTestRepository) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "email@me.com",
		FirstName: "FirstN",
		LastName:  "LastN",
		Password:  "101010",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *PostgresTestRepository) GetOne(id int) (*User, error) {
	user := User{
		ID:        1,
		Email:     "email@me.com",
		FirstName: "FirstN",
		LastName:  "LastN",
		Password:  "101010",
		Active:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *PostgresTestRepository) Update(user User) error {

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *PostgresTestRepository) DeleteByID(id int) error {

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *PostgresTestRepository) Insert(user User) (int, error) {

	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *PostgresTestRepository) ResetPassword(password string, user User) error {

	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *PostgresTestRepository) PasswordMatches(plainText string, user User) (bool, error) {

	return true, nil
}
