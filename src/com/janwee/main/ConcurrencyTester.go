package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(src string,c chan string){
	for {
		c <- fmt.Sprintf("%s: %v", src, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consume(c chan string){
	for {
		str := <- c
		fmt.Println(str)
	}
}

func main(){
	c1 := make(chan string)
	go produce("cat",c1)
	go produce("dog",c1)
	consume(c1)
}

