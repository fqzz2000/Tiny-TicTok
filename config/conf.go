package config
import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)


type Mysql struct {
	Host 		string
	Port 		int 
	Database 	string
	Username 	string
	Password 	string
	Charset 	string
	ParseTime 	bool `toml:"parse_time"`
	loc 	string
}

type Config struct {
	DB Mysql
}

var Info Config


func init() {
	if _, err := toml.DecodeFile("D:\\GOLandPRo\\douyin_pro\\byte_douyin_project\\config\\config.toml", &Info); err != nil {
		panic(err)
	}
	//去除左右的空格
	strings.Trim(Info.DB.Host, " ")
}

func 