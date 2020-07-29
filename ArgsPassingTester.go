package main

import (
	"fmt"
	"strconv"
)

/*
an arguments passing tester
 */
func main() {
	var nums=[]int{0,1,2}
	fmt.Printf("address of nums : %p\n", &nums)
	var str string="["
	for i,num := range nums{
		str+=strconv.Itoa(num)
		if i< len(nums)-1{
			str+=","
		}else{
			str+="]"
		}
	}
	fmt.Println("value of nums :"+str)
	argsPassingTest(nums)
	fmt.Printf("address of nums : %p\n", &nums)
	var str2 string="["
	for i,num := range nums{
		str2+=strconv.Itoa(num)
		if i< len(nums)-1{
			str2+=","
		}else{
			str2+="]"
		}
	}
	fmt.Println("value of nums :"+str2)
}

func argsPassingTest(nums []int) {
	nums[1] = 3
	fmt.Printf("address of nums : %p\n", &nums)
	var str string="["
	for i,num := range nums{
		str+=strconv.Itoa(num)
		if i< len(nums)-1{
			str+=","
		}else{
			str+="]"
		}
	}
	fmt.Println("value of nums :"+str)
}
