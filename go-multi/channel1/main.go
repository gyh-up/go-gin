package main

import "fmt"

type Dog struct {
	Name string
	Color string
}
func main() {
	allChan := make(chan interface{}, 10)
	allChan <- Dog{"小黄", "黄色"}
	allChan <- 1
	allChan <- "很活跃"

	// ------- 类型断言
	dog := <-allChan
	fmt.Printf("%T", dog)
	a := dog.(Dog) // 获取到dog，但是没有获取到dog的属性，需要进行 类型断言
	fmt.Printf("%v", a.Color)

	// ------- channel 的循环遍历与关闭
	close(allChan)   // 不关闭channel 无法循环获取到channel的数据，但是关闭后不能继续写入
	//for v := range allChan {
	//	fmt.Println(v)
	//}

	for {
		val, ok := <-allChan
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
