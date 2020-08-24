package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Generateid(params ...interface{}) (string,error) {
	//if len(params) == 0{
	//	errors.New("type is required")
	//}
	//if len(params) == 2 {
	//	size := params
	//}

	var size int

	switch len(params) {
	case 0:
		return "",errors.New("type is required")
	case 1:
		size = 15
	case 2:
		size = params[1].(int)
	case 3:
		size = params[1].(int)
	}

	switch params[0] {
	case "i":
		// generate number id
		y := strconv.Itoa(time.Now().Year())[2:]
		m := strconv.Itoa(int(time.Now().Month()))
		d := strconv.Itoa(time.Now().Day())
		h := strconv.Itoa(time.Now().Hour())
		t := strconv.Itoa(time.Now().Minute())
		s := strconv.Itoa(time.Now().Second())
		n := strconv.Itoa(time.Now().Nanosecond())[2:4]
		p := strconv.Itoa(os.Getpid())[2:]



		uid := n+p+s+t+h+d+m+y
		return uid[:size],nil
	default:
		return "",errors.New("invalid type specified")
	}
}

func main()  {
	id, _ := Generateid("i",6,nil)
	fmt.Println(id)
}
