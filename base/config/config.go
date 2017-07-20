package config

import (
	"github.com/astaxie/beego/config"
	"os"
	"path/filepath"
)

type Path struct {
	RootPath string
}

var (
	Ini config.Configer
)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "config.ini")
	cnf, isValid := config.NewConfig("ini", appConfigPath)
	if isValid != nil {
		println("fail to open file : config.ini " + isValid.Error())
		return
	}
	//配置对象
	Ini = cnf
}
