//.package utils

//数据库工具

// 这份代码实现了二级缓存的核心功能，包括：

// 将全员考试成绩单（姓名/成绩键值对map）写入MySQL数据库
// 根据姓名从Redis缓存查询分数
// 将姓名与分数写入Redis缓存
// 还实现了一个通用的MySQL表查询方法：根据任意表名和查询条件map进行查询，并将结果送入指定的指针地址中，从而具有较高的复用价值
package utils

import (
	"sync"
)

var (
	//数据库读写锁,防止并发情况下的读写问题
	dbMutext sync.RWMutex
)

//通用的Mysql查询工具
//tableName 查询的表名
//argsMap 查询条件的集合
//dest 查询结果存储地址

func QueryFromMysql(tableName string, argsMap map[string]interface{}, dest interface{}) (err error) {
	fmt.Println("queryscorefrommysql..")
}
