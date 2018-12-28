package main

import "fmt"

func isPalindrome(x int) bool {

	//排除小于0，和最后以为是0的情况
	if x < 0 || (x%10 == 0) && x != 0 {
		return false
	}
	//初始化
	temp := 0
	for x > temp {
		temp = temp*10 + x%10
		x /= 10
	}
	//判断
	return x == temp || x == temp/10
}
func main() {
	if isPalindrome(1241) {
		fmt.Println("是回文数")
	} else {
		fmt.Println("不是回文数")
	}

}
