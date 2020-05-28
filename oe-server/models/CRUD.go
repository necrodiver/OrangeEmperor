package models

import (
	"database/sql"
	"fmt"
	"time"

	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

// logrus 日志记录

// People - database fields
type People struct {
	id           int
	username     string
	password     string
	createdID    int
	createdTime  time.Time
	modifiedID   int
	modifiedTime time.Time
	state        int
	deletedOn    int
}

type appContext struct {
	db *sql.DB
}

func connectDB(driverName string, dbName string) (*appContext, string) {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return nil, err.Error()
	}
	if err = db.Ping(); err != nil {
		return nil, err.Error()
	}
	return &appContext{db}, ""
}

// Create
func (c *appContext) Create() {
	stmt, err := c.db.Prepare("INSERT INTO blog_auth(username,password) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec("zhangsan", "123456")
	if err != nil {
		fmt.Printf("add error: %v", err)
		return
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted id is ", lastID)
}

// Read
func (c *appContext) Read() {
	rows, err := c.db.Query("SELECT * FROM users where id<?", 30)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := new(People)
		err := rows.Scan(&p.id, &p.username, &p.password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(p.id, p.username, p.password)
	}
}

// UPDATE
func (c *appContext) Update() {
	stmt, err := c.db.Prepare("UPDATE users SET age = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(10, 1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("update affect rows is ", affectNum)
}

// DELETE
func (c *appContext) Delete() {
	stmt, err := c.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete affect rows is ", affectNum)
}

// Mysqlite3 - sqlite3 CRUD
func main() {
	c, err := connectDB("sqlite3", "./db/blog.db")
	if err != "" {
		print(err)
	}

	c.Create()
	fmt.Println("add action done!")

	c.Read()
	fmt.Println("get action done!")

	c.Update()
	fmt.Println("update action done!")

	c.Delete()
	fmt.Println("delete action done!")
}
