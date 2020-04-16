//.package utils

//起名

package utils

var (
	//姓氏
	familyNames    = []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "楚", "卫", "蒋", "沈", "韩", "杨", "张", "欧阳", "东门", "西门", "上官", "诸葛", "司徒", "司空", "夏侯"}
	middleNamesMap = map[string][]string{} //string: []映射
	lastNames      = []string{"春", "夏", "秋", "冬", "风", "霜", "雨", "雪", "木", "禾", "米", "竹", "山", "石", "田", "土", "福", "禄", "寿", "喜", "文", "武", "才", "华"}
)

//初始化姓氏和对应的辈分
func init() {
	for _, x := range familyNames {
		if x != "欧阳" {
			middleNamesMap[x] = []string{"德", "惟", "守", "世", "令", "子", "伯", "师", "希", "与", "孟", "由", "宜", "顺", "元", "允", "宗", "仲", "士", "不", "善", "汝", "崇", "必", "良", "友", "季", "同"}
		} else {
			middleNamesMap[x] = string{"宗", "的", "永", "其", "光"}
		}
	}
}

//获得随机姓名
func GetRandomName() (name string) {
	familyName := familyNames[GetRandomName(0, len(familyNames)-1)]
	middlerName := middlerName[familyName][GetRandomInt(0, len(middleNamesMap[familyName])-1)]
	lastNames := lastNames[GetRandomInt(0, len(lastNames)-1)]
	return familyName + middlerName + lastNames
}
