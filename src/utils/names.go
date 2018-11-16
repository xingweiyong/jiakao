package utils

/*
随机生成考生名字
名字由first middle last 三部分组成
*/

var (
	//姓氏
	familyNames = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳", "东门", "西门", "上官", "诸葛", "司徒", "司空", "夏侯"}
	//辈分
	middleNameMap = map[string][]string{}
	//名字
	lastNames = []string{"春", "夏", "秋", "冬", "风", "霜", "雨", "雪", "木", "禾", "米", "竹", "山", "石", "田", "土", "福", "禄", "寿", "喜", "文", "武", "才", "华"}
)

func init()  {
	for _,x := range familyNames{
		if x != "欧阳"{
			middleNameMap[x] = []string{"德", "惟", "守", "世", "令", "子", "伯", "师", "希", "与", "孟", "由", "宜", "顺", "元", "允", "宗", "仲", "士", "不", "善", "汝", "崇", "必", "良", "友", "季", "同"}
		}else {
			middleNameMap[x] = []string{"宗", "的", "永", "其", "光"}
		}
	}
}

func GetRandowName() (name string) {
	familyName := familyNames[GetRandowInt(0,len(familyNames))]
	middleName := middleNameMap[familyName][GetRandowInt(0,len(middleNameMap[familyName]))]
	lastName := lastNames[GetRandowInt(0,len(lastNames))]
	return familyName+middleName+lastName
}

