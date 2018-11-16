package main

import (
	"fmt"
	"utils"
)

func main() {
	num := utils.GetRandowInt(10, 20)
	fmt.Println(num)
	name := utils.GetRandowName()
	fmt.Println(name)
	//var scoreMap = map[string]int{"aa": 1, "bb": 2}
	var data = []map[string]string{
		{"name":"aa","score":"1"},
		{"name":"bb","score":"2"},
	}
	utils.InsertMany2Mysql("insert into score(name,score) values ", data)
	//for name, score := range scoreMap {
	//	valueSlice := []interface{}{name,score}
	//	e := utils.Insert2Mysql("insert into score(name,score) values(?,?)", valueSlice)
	//	utils.HandlerError(e, `Write2Mysql("score", argsMap)`)
	//}
	/*考试成绩*/
	//type ExamScore struct {
	//	Id    int    `db:"id"`
	//	Name  string `db:"name"`
	//	Score int    `db:"score"`
	//}
	//scores := make([]ExamScore, 0)
	//argsMap := make(map[string]interface{})
	//argsMap["name"] = "aa"
	//err := utils.QueryFromMysql( "select * from score where id > 3",&scores)
	//if err != nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(scores)
	//}
}
