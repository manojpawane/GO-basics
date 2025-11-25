package main

import (
	"fmt"
	"time"
)


func main(){
	t := time.Now()
	switch  {
	case t.Hour() < 12: fmt.Print("Good morning")
	case t.Hour() < 17: fmt.Print("Good afternoon")
	default : fmt.Print("Good evening")
	}
}