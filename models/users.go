package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (user *User) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", user.ID)
	return db.QueryRow(statement).Scan(&user.Name, &user.Age)
}

func (user *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", user.Name, user.Age, user.ID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", user.ID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(name, age) Values('%s', '%d')", user.Name, user.Age)
	_, err := db.Exec(statement)

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetUsers(db *sql.DB, start, count int) ([]User, error) {
	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
