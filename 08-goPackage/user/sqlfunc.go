package user

import (
	"database/sql"
	"fmt"
	"log"
)

func insertUsers(db *sql.DB, user *User) (int, error) {
	row := db.QueryRow("INSERT INTO users (name,age) VALUES($1,$2) RETURNING id ",
		user.Name,
		user.Age,
	)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("Can't insert data", err)
		return 0, err
	}
	return id, nil
}
func queryOneRow(db *sql.DB, idRow int) (User, error) {
	stmt, err := db.Prepare("SELECT id, name, age FROM users where id =$1;")
	if err != nil {
		log.Fatal("can't prepare query one row  users statment !", err)
	}
	var user User
	row := stmt.QueryRow(idRow)

	err = row.Scan(&user.ID, &user.Name, &user.Age)

	if err != nil {
		return User{}, err
	}
	return user, err

}
func queryAll(db *sql.DB) ([]User, error) {
	stmt, err := db.Prepare("SELECT id, name, age FROM users;")
	if err != nil {
		log.Fatal("can't prepare query all  users statment !", err)
	}
	var users []User

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal("can't  query all  users  !", err)
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Fatal("can't Scan row into variable  :", err)

		}
		users = append(users, user)

		if err = rows.Err(); err != nil {
			return nil, err

		}
	}
	return users, nil
}
func updateUser(db *sql.DB, id int, name string) {
	stmt, err := db.Prepare("UPDATE users SET name=$2 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}
	if _, err := stmt.Exec(id, name); err != nil {
		log.Fatal("error execute update ", err)
	}
	fmt.Println("update success")

}
func deleteUserByID(db *sql.DB, idRow int) {
	stmt, err := db.Prepare("DELETE  FROM users where id =$1;")
	if err != nil {
		log.Fatal("can't prepare delete statement", err)
	}
	if _, err := stmt.Exec(idRow); err != nil {
		log.Fatal("can't execute delete statment", err)
	}

	fmt.Println("Delete success")

}
