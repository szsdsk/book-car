package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	SslMode    string
)

func init() {
	file, err := ini.Load("src/config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadData(file)
}

// LoadData 读取配置文件，更好的连接数据库。
func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPost").MustString("5432")
	DbUser = file.Section("database").Key("DbUser").MustString("postgres")
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").MustString("acs")
	SslMode = file.Section("database").Key("SslMode").MustString("disable")
}
