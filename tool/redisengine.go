package tool

import (
	"fmt"
	_ "github.com/go-redis/redis"
	 "github.com/garyburd/redigo/redis"
)

/*
此处添加了过期时间，在token中也添加了过期时间，到时候验证时选择其中之一即可
*/

// 初始化并连接redis 默认使用0号数据库
func InitRedis()(redis.Conn, error){
	redisconf := GetConfig().Redisdata
	address := redisconf.Address + ":" +redisconf.Port
	rdb,err := redis.Dial("tcp",address)
	if err != nil {
		return nil, err
	}
	if _, err := rdb.Do("AUTH", redisconf.Password); err != nil {
		rdb.Close()
		return nil, err
	}
	/*
	if _, err := rdb.Do("SELECT", "0"); err != nil {
		rdb.Close()
		return nil, err
	}
	*/
	return rdb, nil
}

func Setjwt(useraccount string,token string) {
	//keyoftoken := useraccount
	rdb,err:= InitRedis()
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	// 添加token到redis中去
	defer rdb.Close()
	_, err = rdb.Do("Set", useraccount, token)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置过期时间
	_, err = rdb.Do("expire", useraccount, 60*60*3)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func Getjwt(useraccount string, token string) (bool,error) {
	rdb,err:= InitRedis()
	if err != nil {
		return false,err
	}
	// 添加token到redis中去
	defer rdb.Close()
	// 返回byte类型
	result, err := rdb.Do("Get", useraccount)
	if err != nil {
		return false, err
	}
	if result == nil {
		return false,nil
	}
	byteresult := result.([]byte)
	strresult := string(byteresult)
	if strresult == token {
		return true, nil
	} else {
		return false,nil
	}
}