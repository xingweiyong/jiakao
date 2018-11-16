package utils

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	//数据库读写锁
	dbMutex sync.RWMutex
	//mysql
	mysqlDB *sqlx.DB
	//redis
	redisPool *redis.Pool
)

/*全局数据库初始化*/
func InitMysql() {
	db, err := sqlx.Connect("mysql", "root:Bigdata@1234@tcp(47.92.243.212:3306)/test")
	HandlerError(err, "初始化全局数据库失败！")
	mysqlDB = db
}

/*关闭MySQL数据库*/
func CloseMysql() {
	if mysqlDB != nil {
		mysqlDB.Close()
	}
}

/*获得redis连接*/
func GetRedisConn() (redis.Conn, error) {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxActive:   100,
			MaxIdle:     10,
			IdleTimeout: 10,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", "47.92.243.212:1299")
				if err != nil {
					return nil, err
				}
				if _, err := c.Do("AUTH", "Bigdata@1234"); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
		}
	}
	return redisPool.Get(), nil
}

/*关闭redis连接*/
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}

/*mysql 增删查改*/
func QueryFromMysql(sql string, dest interface{}) (err error) {
	/*
		scores := make([]ExamScore, 0)
		argsMap := make(map[string]interface{})
		argsMap["name"] = "aa"
		err := utils.QueryFromMysql( "select * from score where id > 3",&scores)
	*/
	dbMutex.RLock()
	fmt.Println("QueryFromMysql...")
	InitMysql()
	defer CloseMysql()
	err = mysqlDB.Select(dest, sql)
	if err != nil {
		fmt.Println("QueryMysqlError...", err)
	}
	dbMutex.RUnlock()
	return err
}

func Insert2Mysql(sql string, valueSlice []interface{}) (err error) {
	/*
		for name, score := range scoreMap {
		valueSlice := []interface{}{name,score}
		e := utils.Insert2Mysql("insert into score(name,score) values(?,?)", valueSlice)
		utils.HandlerError(e, `Write2Mysql("score", argsMap)`)
	*/
	dbMutex.Lock()
	InitMysql()
	defer CloseMysql()
	_, err = mysqlDB.Exec(sql, valueSlice...)
	if err != nil {
		fmt.Println("Write2Mysql Error!", err)
	}
	dbMutex.Unlock()
	return err
}

func InsertMany2Mysql(sql string, data []map[string]string) (err error) {
	/*
		var data = []map[string]string{
			{"name":"aa","score":"1"},
			{"name":"bb","score":"2"},
		}
		utils.InsertMany2Mysql("insert into score(name,score) values ", data)
	*/
	dbMutex.Lock()
	InitMysql()
	defer CloseMysql()
	for k, v := range data {
		if k == 0 {
			temp := "("
			for _, t := range v {
				temp += (fmt.Sprintf("%q", t) + ",")
			}
			temp = temp[:len(temp)-1]
			temp += ")"
			sql += temp
		} else {
			temp := ",("
			for _, t := range v {
				temp += (fmt.Sprintf("%q", t) + ",")
			}
			temp = temp[:len(temp)-1]
			temp += ")"
			sql += temp
		}
	}
	_, err = mysqlDB.Exec(sql)
	if err != nil {
		fmt.Println("WriteMany2Mysql Error!", err)
	}
	dbMutex.Unlock()
	return err
}
