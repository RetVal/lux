package lux

import (
	"log"
)

//D is simply a alias to log.Println(... interface{}). Helps to debug without having to import "log" all the time.
func D(x ...interface{}) {
	log.Println(x)
}
