package doa

import "database/sql"

// BaseDoa 父级DOA操作
type BaseDoa struct {
}

// OpenDB 打开数据库
func (b BaseDoa) OpenDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "./db/blog.db")
	if err != nil {
		panic(err)
	}
	return DB
}

// CloseDB 关闭数据库
func (b BaseDoa) CloseDB(DB *sql.DB) error {
	if DB != nil {
		err := DB.Close()
		return err
	}
	return nil
}
