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

func Generateid(params ...interface{}) (string, error) {

	rand.Seed(time.Now().UTC().UnixNano())

	var size int
	switch len(params) {
	case 0:
		return "", errors.New("type is required")
	case 1:
		size = 16
	case 2:
		size = params[1].(int)
	}

	y := strconv.Itoa(time.Now().Year())[2:]
	m := strconv.Itoa(int(time.Now().Month()))
	d := strconv.Itoa(time.Now().Day())
	h := strconv.Itoa(time.Now().Hour())
	t := strconv.Itoa(time.Now().Minute())
	s := strconv.Itoa(time.Now().Second())[:1]
	n := strconv.Itoa(time.Now().Nanosecond())[2:4]
	p := strconv.Itoa(os.Getpid())[:3]

	if len(m) == 1 {
		m = "0" + m
	}
	if len(d) == 1 {
		d = "0" + d
	}
	if len(h) == 1 {
		h = "0" + h
	}
	if len(t) == 1 {
		t = "0" + t
	}
	uid := n + p + s + t + h + d + m + y
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
	case "i":
		return uid, nil
	case "a":
		alphabets := []string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "f", "F", "g", "G", "h", "H", "i", "I", "j", "J", "k", "K", "l", "L", "m", "M", "n", "N", "o", "O", "p", "P", "q", "Q", "r", "R", "s", "S", "t", "T", "u", "U", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z"}

		uidarray := strings.Split(uid, "")
		stringuidarray := make([]string, size)
		for i := 0; i < len(uidarray); i++ {
			j, _ := strconv.Atoi(uidarray[i])
			j += rand.Intn(42)
			stringuidarray[i] = alphabets[j]
		}
		stringuid := strings.Join(stringuidarray, "")
		return stringuid, nil
	case "n":
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
		alphauid := strings.Join(uidarray, "")
		return alphauid, nil
	default:
		return "", errors.New("invalid type specified")
	}
}
