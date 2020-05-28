package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

// User User结构体
type User struct {
	ID          int    `json:"id"`
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	CreatedTime string `json:"createdTime"`
}

// StartDb 启动数据库并处理
func StartDb() []User {
	db, err := sql.Open("sqlite3", "./db/blog.db")
	checkErr(err)
	rows := GetUserList(db)
	var (
		id           int
		username     string
		password     string
		createdID    int
		createdTime  time.Time
		modifiedID   int
		modifiedTime time.Time
		state        int
		deletedOn    int
	)
	var userList []User
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &createdID, &createdTime, &modifiedID, &modifiedTime, &state, &deletedOn)
		if err != nil {
			checkErr(err)
		}
		userInfo := User{
			ID:          id,
			UserName:    username,
			PassWord:    password,
			CreatedTime: transFormat(&createdTime),
		}
		log.Println(id, username, password, createdID, transFormat(&createdTime), modifiedID, transFormat(&modifiedTime), state, deletedOn)
		userList = append(userList, userInfo)
	}
	defer rows.Close()
	return userList
}
func transFormat(t *time.Time) string {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return ""
}

// GetUserList 获取用户列表
func GetUserList(db *sql.DB) *sql.Rows {
	rows, err := db.Query("select * from blog_auth")
	checkErr(err)
	return rows
}

// checkErr 错误程序终止运行
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
