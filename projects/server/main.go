package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID           int
	Name         string
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5436 user=postgres dbname=postgres password=1234 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// err = insertUser(db, User{
	// 	Name:     "John Doe",
	// 	Email:    "john@doe.com",
	// 	Password: "password",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

	err = updateUser(db, 1, User{
		Name:  "Grisha",
		Email: "grisha@gmail.com",
	})

	users, err = getUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)

}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUserById(db *sql.DB, id int) (User, error) {
	var u User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", 2).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)

	if err != nil {
		return u, err
	}

	return u, nil
}

func insertUser(db *sql.DB, u User) error {
	_, err := db.Exec("insert into users (name, email, password) values ($1, $2, $3)",
		u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func updateUser(db *sql.DB, id int, newUser User) error {
	_, err := db.Exec("update users set name=$1, email=$2 where id=$3", newUser.Name, newUser.Email, id)

	if err != nil {
		return err
	}

	return nil
}
