package tool

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

// 系统的全部配置结构
type Config struct {
	Apphost string `json:"Apphost"`
	Port string `json:"Appport"`
	Database Databaseconf `json:"database"`
	Redisdata Redisconf `json:"redis"`
}

// 数据库的结构
type Databaseconf struct {
	Type string `json:"Type"`
	User string `json:"User"`
	Password string `json:"Password"`
	Host string `json:"Host"`
	Port string `json:"Port"`
	DBName string `json:"DBName"`
	Charset string `json:"Charset"`
	Showsql bool `json:"Showsql"`
}

// Redis数据库结构
type Redisconf struct {
	Address string `json:"Address"`
	Port string `json:"Port"`
	Password string `json:"Password"`
	Poolsize int64 `json:"Poolsize"`
}

// 定义一个全局变量表示得到的config对象
var _config *Config = nil

// 解析config
func ParseConfig(path string) (*Config, error) {
	jsonfile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonfile.Close()

	// 读取文件并解析成json格式
	reader := bufio.NewReader(jsonfile)
	bytejsonfile, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	// 将配置里的对象直接匹配上面的Config struct
	err = json.Unmarshal(bytejsonfile, &_config)
	return _config, nil
}

// 返回config对象
func GetConfig() (*Config) {
	return _config
}
