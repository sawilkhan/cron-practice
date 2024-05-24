package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	loc , err := time.LoadLocation("EST")
	if err != nil{
		panic(err)
	}
	s := gocron.NewScheduler(loc)

	//s.Every(1).Second().Do(task)

	s.LimitRunsTo(5).Every(1).Second().Do(task)
	s.StartBlocking()
	fmt.Println("After task statement")
}

func task(){
	fmt.Println("Kaam")
}