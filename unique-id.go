package uniqueid

import (
	"errors"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Generateid returns unique ID based on the parameters provided
func Generateid(params ...interface{}) (string, error) {

	rand.Seed(time.Now().UTC().UnixNano())

	var size int
	prefix := ""
	switch len(params) {
	case 0:
		return "", errors.New("type is required")
	case 1:
		size = 16
	case 2:
		size = params[1].(int)
	case 3:
		if params[1] == nil{
			size = 16
		}else {
			size = params[1].(int)
		}
		prefix = params[2].(string)
	}

	//fmt.Println(prefix)

	y := strconv.Itoa(time.Now().Year())[2:]
	m := "00" + strconv.Itoa(int(time.Now().Month()))
	d := "00" + strconv.Itoa(time.Now().Day())
	h := "00" + strconv.Itoa(time.Now().Hour())
	t := "00" + strconv.Itoa(time.Now().Minute())
	s := strconv.Itoa(time.Now().Second())[:1]
	n := strconv.Itoa(time.Now().Nanosecond())[2:4]
	p := strconv.Itoa(os.Getpid())[:3]

	uid := n + p + s + t[len(t)-2:] + h[len(h)-2:] + d[len(d)-2:] + m[len(m)-2:] + y
	if size > 16 {
		narray := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
		b := make([]string, size-16)
		for i := range b {
			b[i] = narray[rand.Intn(len(narray))]
		}
		uid = uid + strings.Join(b, "")
	}
	if size < 16 {
		uid = uid[:size]
	}

	switch params[0] {
	case "n":
		return prefix+uid, nil
	case "l":
		alphabets := []string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "f", "F", "g", "G", "h", "H", "i", "I", "j", "J", "k", "K", "l", "L", "m", "M", "n", "N", "o", "O", "p", "P", "q", "Q", "r", "R", "s", "S", "t", "T", "u", "U", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z"}

		uidarray := strings.Split(uid, "")
		for i := 0; i < len(uidarray); i++ {
			j, _ := strconv.Atoi(uidarray[i])
			j += rand.Intn(42)
			uidarray[i] = alphabets[j]
		}
		uid = strings.Join(uidarray, "")
		return prefix+uid, nil
	case "a":
		alphabets := []string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "f", "F", "g", "G", "h", "H", "i", "I", "j", "J", "k", "K", "l", "L", "m", "M", "n", "N", "o", "O", "p", "P", "q", "Q", "r", "R", "s", "S", "t", "T", "u", "U", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z"}
		totalletters := math.Floor(0.7 * float64(size))
		randarray := make([]int, int(totalletters))
		for i := 0; i < int(totalletters); i++ {
			randarray[i] = rand.Intn(size)
		}
		uidarray := strings.Split(uid, "")
		for _, v := range randarray {
			j, _ := strconv.Atoi(uidarray[v])
			j += rand.Intn(42)
			uidarray[v] = alphabets[j]
		}
		uid := strings.Join(uidarray, "")
		return prefix+uid, nil
	default:
		return "", errors.New("invalid type specified")
	}
}
