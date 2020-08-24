package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func Generateid(params ...interface{}) (string,error) {
	if len(params) == 0{
		errors.New("type is required")
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
		n := strconv.Itoa(time.Now().Nanosecond())[:4]

		uid := n+y+m+d+h+t+s
		return uid,nil
	default:
		return "",errors.New("invalid type specified")
	}
}

func main()  {
	id, _ := Generateid("i")
	fmt.Println(id)
}
