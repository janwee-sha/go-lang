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
	i := 0
	fmt.Printf("address of int : %p\n", &i)
	fmt.Printf("value of int : %d\n",i)
	updatePassingInt(i)
	fmt.Printf("address of int : %p\n", &i)
	fmt.Printf("value of int : %d\n",i)

	fmt.Println("---------------------------------------")

	fmt.Printf("address of nums : %p\n", &nums)
	var str string
	for i,num := range nums{
		str+=strconv.Itoa(num)
		if i< len(nums)-1{
			str+=","
		}
	}
	fmt.Println("value of int array :"+str)
	updatePassingArr(nums)
	fmt.Printf("address of nums : %p\n", &nums)
	var str2 string
	for i,num := range nums{
		str2+=strconv.Itoa(num)
		if i< len(nums)-1{
			str2+=","
		}
	}
	fmt.Println("value of int array :"+str2)
}

func updatePassingArr(nums []int) {
	fmt.Printf("address of int array : %p\n", &nums)
	nums[1] = 3
	var str string
	for i,num := range nums{
		str+=strconv.Itoa(num)
		if i< len(nums)-1{
			str+=","
		}
	}
	fmt.Println("value of int array :"+str)
}

func updatePassingInt(i int){
	fmt.Printf("address of int : %p\n", &i)
	i=3
	fmt.Printf("value of int : %d\n",i)
}
