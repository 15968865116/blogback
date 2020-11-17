package tool

import (
	"finalgo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	// 一种结构 数据库结构
	*xorm.Engine
}

// 定义一个全局的orm 数据库引擎 用以其他地方的调用
var DBengine *Orm

// 连接并操作数据库
func OrmEngine(_config *Config) (*Orm, error){
	//连接参数
	database := _config.Database
	conn := database.User+":"+database.Password+"@tcp("+database.Host+":"+database.Port+")/"+database.DBName+"?charset="+database.Charset
	//连接数据库 获取数据库engine信息，就是获取了一个对象
	engine,err := xorm.NewEngine(database.Type,conn)
	if err != nil {
		return nil, err
	}
	engine.ShowSQL(database.Showsql)
	//新建一个User数据表
	err = engine.Sync2(new(model.User))
	// 新建一个Blog表
	err = engine.Sync2(new(model.Blog))
	err = engine.Sync2(new(model.Comment))
	orm := new(Orm)
	orm.Engine = engine
	DBengine = orm
	return orm,nil
}
