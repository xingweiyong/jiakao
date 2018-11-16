package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"utils"
)

var (
	chNames = make(chan string, 100)
	examers = make([]string, 0)

	//信号量，5条车道
	chLanes = make(chan int, 5)
	//违纪者
	chFouls = make(chan string, 100)
	//考试成绩
	scoreMap = make(map[string]int)
)

type ExamScore struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Score int    `db:"score"`
}


func Patrol() {
	//Ticker 间隔时间重复, Timer one time
	ticker := time.NewTicker(1 * time.Second)
	fmt.Println("老鹰开始巡考...")
	//select 多路复用
	for {
		select {
		case name := <-chFouls:
			fmt.Printf("%s,抓到一只小鸡！\n", name)
		default:
			fmt.Println("考场纪律良好。")
		}
		<-ticker.C
	}
}

func TakeExam(name string) {
	chLanes <- 123
	fmt.Printf("%s，正在考试...", name)
	examers = append(examers, name)
	score := utils.GetRandowInt(0, 100)
	scoreMap[name] = score
	if score < 10 {
		score = 0
		chFouls <- name
	}
	<-time.After(400 * time.Millisecond)
	<-chLanes
}

/*二级缓存查询成绩*/
func QueryScore(name string) {
	conn, err := utils.GetRedisConn()
	if err != nil {
		fmt.Println("GetRedisConn error!")
	}
	score, err := conn.Do("get", name)
	if err != nil || score == nil {
		// 从MySQL查询
		score := make([]ExamScore, 0)
		sql := fmt.Sprintf("select * from score where name=%q", name)
		err := utils.QueryFromMysql(sql, &score)
		if err == nil {
			if len(score) >0{
				fmt.Printf("Mysql成绩：%s,%v", score[0].Name, score[0].Score)
				// 把成绩同步到redis
				conn.Do("set",score[0].Name, score[0].Score)
			} else {
				fmt.Printf("%s，还未参加考试，暂无成绩！",name)
			}
		} else {
			fmt.Printf("mysql 查询错误！")
		}
	} else {
		//redis
		score, err = redis.Int(score, err)
		if err == nil {
			fmt.Printf("Redis成绩：%s,%d", name, score)
		} else {
			fmt.Printf("redis 类型转换错误！")
		}
	}
}

func WriteScore2Mysql(scoreMap map[string]int)  {
	for name,score := range scoreMap{
		valueSlice := []interface{}{name,score}
		e := utils.Insert2Mysql("insert into score(name,score) values(??)",valueSlice)
		utils.HandlerError(e,"WriteScore2Mysql error!!")
	}
	fmt.Println("成绩录入完毕")
}


func main() {
	QueryScore("bb")
}
