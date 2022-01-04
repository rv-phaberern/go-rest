package main

import "github.com/jmoiron/sqlx"

// Store will serve to execute queries
type Store struct {
	db *sqlx.DB
}

// NewStore returns a new instance of a store
func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// GetUsers ...
func (s *Store) GetUsers() ([]*User, error) {
	const query = `
		SELECT 
			UserID,
			UserName,
			FirstName,
			LastName
		FROM
			Test.User
	`
	users := make([]*User, 0)
	err := s.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID ...
func (s Store) GetUserByID(id int) (*User, error) {
	const query = `
		SELECT 
			UserID,
			UserName,
			FirstName,
			LastName
		FROM
			Test.User
		WHERE 
			UserID = ?
	`
	user := User{}
	err := s.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// AddUser ...
func (s *Store) AddUser(user *User) error {
	const query = `
		INSERT INTO 
			Test.User(UserName, FirstName, LastName)
		VALUES
			(?, ?, ?)
	`
	res, err := s.db.Exec(query, user.UserName, user.FirstName, user.LastName)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	user.UserID = int(id)

	return nil
}

// UpdateUser ...
func (s *Store) UpdateUser(user *User) error {
	const query = `
		UPDATE
			Test.User
		SET 
			UserName = ?,
			FirstName = ?,
			LastName = ?
		WHERE
			UserID = ?
	`
	_, err := s.db.Exec(query, user.UserName, user.FirstName, user.LastName, user.UserID)
	if err != nil {
		return err
	}

	return nil
}
