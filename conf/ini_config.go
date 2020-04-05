//.package conf

package conf

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type IniConfig struct {
	ConfigMap map[string]string //存储配置字段映射
	strcet    string
}

//配置类型
// [log]
// level = 1
// adapter = file

// [db]
// host = 127.0.0.1
// port = 3306
// user = root
// password = 654321
// name = chain_book
// timezone = Asia/Shanghai
// prefix = bo_
//疑问，返回的结构体指针和结构体有什么区别
//整体的思路就是每一行文本的获取，需要注意处理每一行最后文本值的注释。
func GetIniConfig(filename string) (*IniConfig, error) {
	middle := "."
	config := new(IniConfig)                   //默认值构造
	config.ConfigMap = make(map[string]string) //初始化空值
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	read := bufio.NewReader(file) //文件打开内容的缓存区
	for {                         //无限循环
		b, _, err := read.ReadLine()
		if err != nil {
			if err == io.EOF { //到达末尾才退出
				break
			}
			return nil, err
		}
		str := strings.TrimSpace(string(b))
		//配置文件中的注释
		if strings.Index(str, "#") == 0 { //注释不处理
			continue
		}
		//配置文件中的前缀处理
		n1 := strings.Index(str, "[")
		n2 := strings.LastIndex(str, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			config.strcet = strings.TrimSpace(str[n1+1 : n2]) //记录配置头
			continue
		}
		if len(config.strcet) < 1 {
			continue
		}
		//
		eqIndex := strings.Index(str, "=")
		eqLeft := strings.TrimSpace(str[0:eqIndex])
		if eqIndex < 0 {
			continue
		}
		eqRight := strings.TrimSpace(str[eqIndex+1:])
		if len(eqLeft) < 1 {
			continue
		}
		pos := strings.Index(eqRight, "\t#") //查找右侧的注释
		val := eqRight
		if pos > 1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}
		pos = strings.Index(eqRight, " #") //空格的注释
		if pos > 1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}
		pos = strings.Index(eqRight, "\t//") //查找右侧的注释
		if pos > 1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}
		pos = strings.Index(eqRight, " //") //空格的注释
		if pos > 1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}
		if len(val) < 1 {
			continue
		}
		key := config.strcet + middle + eqLeft //组装map.key
		config.ConfigMap[key] = strings.TrimSpace(val)
	}
	return config, nil

}

//获取指定的key值
func (self *IniConfig) Get(key string) string {
	v, ok := self.ConfigMap[key]
	if ok {
		return v
	}
	return ""
}

//实现接口的所有方法
func (self *IniConfig) GetString(key string) string {
	return self.Get(key)
}

func (self *IniConfig) GetInt(key string) (int, error) {
	return strconv.Atoi(self.Get(key)) //将字符转成整形
}

func (self *IniConfig) GetInt64(key string) (int64, error) {
	return strconv.ParseInt(self.Get(key), 10, 64)
}

func (self *IniConfig) GetFloat(key string) (float64, error) {
	return strconv.ParseFloat(self.Get(key), 64)
}

func (self *IniConfig) GetBool(key string) (bool, error) {
	return strconv.ParseBool(self.Get(key))
}
