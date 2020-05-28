package service

import (
	"oe_server/doa"
)

// SignIn 登录
func SignIn(userName, passWord string) (doa.BlogAuth, error) {
	ba := doa.BlogAuth{}
	err := ba.GetBlogAuth(userName, passWord)
	return ba, err
}
