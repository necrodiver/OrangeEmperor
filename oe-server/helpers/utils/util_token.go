package utils

import (
	"oe_server/helpers/setting"
	"strings"
	"time"
)

// CreateToken 创建token
func CreateToken(id int) (string, error) {
	s1 := string(setting.TOKEN_TIMEOUT * time.Second)
	var sb strings.Builder
	sb.WriteString(s1)
	sb.WriteString("-")
	sb.WriteString(string(id))
	sb.WriteString("-")
	sb.WriteString(time.Now().Format("2006_01_02 15:04:05"))
	tokenStr := []byte(sb.String())
	return RSAEncrypt(tokenStr)
}
