package config

import (
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
)

var File *goconfig.ConfigFile

const Name = "/conf/conf.ini"
func init() {

	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file := currentDir + Name

	if !fileExist(file) {
		panic(errors.New("配置文件不存在"))
	}

	if len(os.Args) > 1 {
		if os.Args[1] != "" {
			file = os.Args[1] + Name
		}
	}

	File, err = goconfig.LoadConfigFile(file)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(File)

}
func Test()  {
	fmt.Println("aa")
}

func fileExist(str string) bool {
	_,err := os.Stat(str)
	return err == nil || os.IsExist(err)
}
