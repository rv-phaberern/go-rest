package main

// User ...
type User struct {
	UserID    int    `db:"UserID" json:"userID"`
	UserName  string `db:"UserName" json:"userName"`
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
}
