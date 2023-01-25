package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)


type Mysql struct {
	Host 		string
	Port 		int 
	Database 	string
	Username 	string
	Password 	string
	Charset 	string
	ParseTime 	bool `toml:"parse_time"`
	Loc 	string
}

type Path struct {
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB Mysql `toml:"mysql"`
	DBTest Mysql `toml:"mysqltest"`
	Path `toml:"path"`
}

var Info Config


func init() {
	if _, err := toml.DecodeFile("D:\\DKU\\byteCamp\\Tiny-TicTok\\config\\conf.toml", &Info); err != nil {
		panic(err)
	}
	//去除左右的空格
	strings.Trim(Info.DB.Host, " ")
}

func DBConnectString() string{
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Info.DB.Username, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime, Info.DB.Loc)
	log.Println(arg)
	return arg
}

func DebugDBConnectString() string{
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
	Info.DBTest.Username, Info.DBTest.Password, Info.DBTest.Host, Info.DBTest.Port, Info.DBTest.Database,
	Info.DBTest.Charset, Info.DBTest.ParseTime, Info.DBTest.Loc)
log.Println(arg)
return arg
}