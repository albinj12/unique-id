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
		size = 16
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
	case "a":
		// generate number id
		alphabets := []string{"a","A","b","B","c","C","d","D","e","E","f","F","g","G","h","H","i","I","j","J","k","K","l","L","m","M","n","N","o","O","p","P","q","Q","r","R","s","S","t","T","u","U","v","V","w","W","x","X","y","Y","z","Z"}
		y := strconv.Itoa(time.Now().Year())[2:]
		m := strconv.Itoa(int(time.Now().Month()))
		d := strconv.Itoa(time.Now().Day())
		h := strconv.Itoa(time.Now().Hour())
		t := strconv.Itoa(time.Now().Minute())
		s := strconv.Itoa(time.Now().Second())
		n := strconv.Itoa(time.Now().Nanosecond())[2:4]
		p := strconv.Itoa(os.Getpid())[2:]

		uid := n+p+s+t+h+d+m+y
		if size > 16 {
			narray := []string{"1","2","3","4","5","6","7","8","9","0"}
			b := make([]string, size - 16)
			for i := range b {
				b[i] = narray[rand.Intn(len(narray))]
			}
			uid = uid + strings.Join(b,"")
		}
		if size < 16 {
			uid = uid[:size]
		}
		uidarray := strings.Split(uid,"")
		stringuidarray := make([]string, size)
		for i := 0; i < len(uidarray); i++ {
			j,_ := strconv.Atoi(uidarray[i])
			j += rand.Intn(42)
			stringuidarray[i] = alphabets[j]
		}
		stringuid := strings.Join(stringuidarray,"")
		return stringuid[:size],nil
	default:
		return "",errors.New("invalid type specified")
	}
}

