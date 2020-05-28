package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// 配置文件读取
var (
	Cfg          *ini.File
	RunMode      string
	HTTP_PORT    string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize        int
	AES_KEY         string
	RSA_PUBLIC_KEY  string
	RSA_PRIVATE_KEY string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadBase()
	loadServer()
	loadApp()
	loadEncodeKeys()
}

// LoadBase 读取配置运行模式
func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 读取数据库
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTP_PORT = sec.Key("HTTP_PORT").MustString("8082")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 读取APP
func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

// LoadApp 读取APP
func loadEncodeKeys() {
	sec, err := Cfg.GetSection("encodekey")
	if err != nil {
		log.Fatalf("Fail to get section 'encodekey': %v", err)
	}
	AES_KEY = sec.Key("AES_KEY").MustString("")
	RSA_PUBLIC_KEY = sec.Key("RSA_PUBLIC_KEY").MustString("")
	RSA_PRIVATE_KEY = sec.Key("RSA_PRIVATE_KEY").MustString("")
}
