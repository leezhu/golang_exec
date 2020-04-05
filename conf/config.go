//.package conf参照写的读取数据库配置模块
package conf

import (
	"errors"
	"path/filepath"
	"strings"
)

var AppConfg Config

//配置接口,
type Config interface {
	GetString(string) string
	GetInt(string) (int, error)
	GetInt64(string) (int64, error)
	GetFloat(string) (float64, error)
	GetBool(string) (bool, error)
}

//init函数是在包调用的时候初始化，可以有多个
func init() {
	AppConfg, _ = NewConfig("init", "config/app.config")
}

//工厂模式，返回实例。大写开头的函数名是可以被外部调用
func NewConfig(adapter, filename string) (Config, error) {
	path, err := GetCurrentPath(filename)
	if err != nil {
		return nil, err
	}
	switch adapter {
	case "init":
		return GetIniConfig(path)
	default:
		return nil, errors.New("adapter not found")
	}
	return nil, nil
}

//获取当前文件的绝对路径
func GetCurrentPath(filename string) (path string, err error) {
	path, err = filepath.Abs(filename) //获取绝对路径，就当前目录pwd加上文件路径
	if err != nil {
		return "", errors.New("get abs path error")
	}
	path = strings.Replace(path, "\\", "/", -1) //替换斜杠，主要是出现在windows环境下的路径
	path = strings.Replace(path, "\\\\", "/", -1)
	return path, nil
}
