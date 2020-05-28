package doa

import "time"

// BlogAuth 用户模型
type BlogAuth struct {
	Id           int
	UserName     string
	PassWord     string
	CreatedID    int
	CreatedTime  time.Time
	ModifiedID   int
	ModifiedTime time.Time
	State        int
	DeletedOn    int
}

// BlogAuthList 用户数组
type BlogAuthList struct {
	list []BlogAuth
}

// GetBlogAuth 获取单个用户数据
func (b *BlogAuth) GetBlogAuth(userName, passWord string) error {
	base := BaseDoa{}
	db := base.OpenDB()
	db.QueryRow("select * from blog_auth where username = ? and password = ?", userName, passWord).Scan(&b.Id, &b.UserName, &b.PassWord, &b.CreatedID, &b.CreatedTime, &b.ModifiedID, &b.ModifiedTime, &b.State, &b.DeletedOn)
	err := base.CloseDB(db)
	return err
}
