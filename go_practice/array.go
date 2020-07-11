package main

import (
	"fmt"
	"strconv"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type student struct {
	Name string
	Age  int
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	peple People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showb")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, "---", a, b, ret)
	return ret
}

func main() {
	// fmt.Println(int('9'))
	// t := Teacher{}
	// t.ShowB()
	// a := 1
	// b := 2
	// defer calc("1", a, calc("10", a, b))
	// a = 0
	// defer calc("2", a, calc("20", a, b))
	// var c ByteCounter
	// c.Write([]byte("hello"))
	// fmt.Println(c)

	// c = 0
	// var name = "Dolly"
	// fmt.Fprintf(&c, "hello,%s", name)
	// fmt.Println(c)
	// q := []int{1, 2, 4}
	// months := [...]string{1: "January", 2: "1"}
	// check(q)
	// fmt.Println(q[0:3])
	// for k, v := range months {
	// 	fmt.Println(k, v)
	// }
	// var runes []rune
	// for _, r := range "Hello ,世界" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes)

	// var x []int = []int{1, 2, 3, 4}
	// fmt.Println(append(x, x...))

	// var ages = make([]string, 0, 10)
	// fmt.Printf("%q,%d,%d", ages, len(ages), cap(ages))
	// type Movie struct {
	// 	Title  string
	// 	Year   int  `json:"released"`
	// 	Color  bool `json:"color,omi"`
	// 	Actors []string
	// }
	// var movies = []Movie{
	// 	{Title: "cacc", Year: 123, Color: false, Actors: []string{"pau new"}},
	// 	{Title: "cacc1", Year: 123, Color: false, Actors: []string{"pau new"}},
	// }
	// data, err := json.MarshalIndent(movies, "", "	")
	// if err != nil {
	// 	log.Fatalf("Json marshaling failed:%s", err)
	// }
	// fmt.Printf("%s\n", data)

	//-----

	// copy()

	//------flag解析参数
	// var period = flag.Duration("perioad", 1*time.Second, "sleep period")
	// flag.Parse()
	// fmt.Printf("Sleeping for %v...", *period)
	// time.Sleep(*period)
	// fmt.Println()

	// m := make(map[string]*student)
	// stus := []student{
	// 	{Name: "alic1e", Age: 12},
	// 	{Name: "alice2", Age: 13},
	// 	{Name: "alice3", Age: 14},
	// }
	// for _, v := range stus {
	// 	m[v.Name] = &v
	// }
	// for k, v := range m {
	// 	fmt.Println("k=", k, "v=", v.Name)
	// }

	// runtime.GOMAXPROCS(1)
	// wg := sync.WaitGroup{}
	// wg.Add(20)
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("A:", i)
	// 		wg.Done()
	// 	}(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println("B:", i)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// fmt.Println(multiply("11", "11"))
	// c1 := f(0)
	// c2 := f(0)
	// fmt.Println(c1(), c2())
	// fmt.Println(c1(), c2())
	// fmt.Println(c1(), c2())
	// fmt.Println(c1(), c2())
	// fmt.Println(c1(), c2())
	// fmt.Println(c1(), c2())
	f2()
	s := []int{1, 2, 3, 4}
	c := make(chan int)
	go sum(s, c)
	go sum(s, c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func f2() *student {
	s := new(student)
	return s
}

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}
func check(a []int) {
	a[1] = 3
}

func multiply(num1 string, num2 string) string {
	//基本思路，最多110次相乘，然后进行字符相加的操作。
	sum := ""
	if len(num1) == 0 || len(num2) == 0 {
		return num1 + num2
	}
	zero := ""
	for i := len(num1) - 1; i >= 0; i-- {
		val := muti(int(num1[i]-'0'), num2)
		//加0
		val += zero
		zero += "0"
		sum = add(val, sum)
		// fmt.Println("i=", i)
		fmt.Println("sum===", sum)
	}
	return sum
}

//相乘
func muti(val int, num2 string) string {
	fmt.Println("val=", val, "num2=", num2)
	res := ""
	var temp int
	for i := len(num2) - 1; i >= 0; i-- {
		x := num2[i] - '0'
		mutiVal := int(x) * val
		sum := mutiVal + temp
		if sum/10 > 0 {
			temp = sum / 10
		} else {
			temp = 0
		}
		// fmt.Println()
		res = strconv.Itoa(sum%10) + res
		// fmt.Println("res1=", res, "sum=", sum)
	}
	if temp > 0 {
		res += string(temp)
	}
	// fmt.Println("res=", res)
	return res
}

//字符相加,从个位到顶
func add(x string, y string) string {
	// fmt.Println("x=", x, "y=", y)
	res := ""
	var temp int
	for i, j := len(x)-1, len(y)-1; (i >= 0) || (j >= 0); i-- {
		var val int
		if i >= 0 && j >= 0 {
			val = int((x[i] - '0')) + int(((y)[j] - '0'))
		} else if i < 0 {
			val = int((y)[j]) - '0'
		} else {
			val = int(x[i] - '0')
		}
		sum := val + temp
		if sum/10 > 0 {
			temp = 1
		} else {
			temp = 0
		}
		res = strconv.Itoa(sum%10) + res
		j--
	}
	if temp > 0 {
		res = "1" + res
	}
	return res
}

//12334*12222

//30 * 12333
// func
