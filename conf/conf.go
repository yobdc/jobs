package conf

import (
	ini "gopkg.in/ini.v1"
	"log"
)

// Props 保存配置文件设置变量
type Props struct {
	config *ini.File
}

var conf Props

func init() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
	conf.config = config
	log.Println(config)
}
