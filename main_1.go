package main

import (
	"fmt"
	"utils"
)

func main() {
	//conn,err := utils.GetRedisConn()
	//if err != nil {
	//	conn.Do("set","aaa",17)
	//}
	//utils.CloseRedis()
	/*二级缓存查询成绩*/
	//conn, err := utils.GetRedisConn()
	//if err != nil {
	//	fmt.Println("GetRedisConn error!")
	//}
	//var name  = "aaa"
	//score, err := conn.Do("get", name)
	////score,err= redis.Int(score,err)
	//fmt.Println(err,score)
	//var name  = "aa"
	type ExamScore struct {
		Id    int    `db:"id"`
		Name  string `db:"name"`
		Score int    `db:"score"`
	}
	scores := make([]ExamScore, 0)
	var name  = "aa"
	sql := fmt.Sprintf("select * from score where name=%q", name)
	utils.QueryFromMysql(sql, &scores)
	fmt.Println(scores)
}
