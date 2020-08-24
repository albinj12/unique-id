package uniqueid

import (
	"errors"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Generateid(params ...interface{}) (string,error) {

	var size int
	switch len(params) {
	case 0:
		return "",errors.New("type is required")
	case 1:
		size = 15
	case 2:
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
		if size > 15 {
			narray := []string{"1","2","3","4","5","6","7","8","9","0"}
			b := make([]string, size - 15)
			for i := range b {
				b[i] = narray[rand.Intn(len(narray))]
			}
			uid = uid + strings.Join(b,"")
		}

		return uid[:size],nil
	default:
		return "",errors.New("invalid type specified")
	}
}

