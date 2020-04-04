//.package go_example

package main

import "fmt"
import "errors"
import "reflect"

//error是一个接口，可以自定义错误类型
// type Error interface{

// }

//内置的error返回
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can`t work with it")
	} else {
		return arg, nil
	}
}

//自定义错误类型
type argError struct {
	arg  int
	prob string
}

//实现error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

//用自定义的错误返回
func f2(arg int) (int, error) {
	if arg == 42 {
		//必须传结构体指针，否则会报错
		return -1, &argError{-1, "can`t work with it"}
	} else {
		return arg, nil
	}
}
func main() {
	for _, i := range []int{7, 42} {
		if _, e := f1(i); e != nil {
			//e类型是*errors.errorString
			fmt.Println("type of error:", reflect.TypeOf(e))
			fmt.Println("f1 work failed,error:", e)
		} else {
			fmt.Println("f1 work success")
		}
	}
	for _, i := range []int{7, 42} {
		if _, e := f2(i); e != nil {
			fmt.Println("f2 work failed,error:", e)
		} else {
			fmt.Println("f2 work success")
		}
	}
	_, e := f2(42)
	//这里e.(*argError)使用的是类型断言，判断是否属于这种类型
	if ae, ok := e.(*argError); ok {
		fmt.Println("ae.arg=", ae.arg, "ae.prob=", ae.prob)
	}
}
