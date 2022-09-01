package main

import (
"fmt" 
"time"
)
/*
func main() {
	fmt.Println("vim-go")
}
*/
func f() {
	fmt.Println("f function")
}
func oldmain(){
	go f()
	time.Sleep(1*time.Second)
	fmt.Println("main function")
}

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	c:= make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x,y,x+y)
}
