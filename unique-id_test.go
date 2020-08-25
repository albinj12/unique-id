package uniqueid

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// test for checking the length of number id
func TestLengthnumberID(t *testing.T) {
	length := rand.Intn(100)
	id, err := Generateid("n", length)
	if err != nil {
		t.Error("error while generating id.", err)
	}
	if (len(id)) != length {
		t.Errorf("got %d, want %d", len(id), length)
	}
}

// test for checking the length of letter id
func TestLengthletterID(t *testing.T) {
	length := rand.Intn(100)
	id, err := Generateid("l", length)
	if err != nil {
		t.Error("error while generating id.", err)
	}
	if (len(id)) != length {
		t.Errorf("got %d, want %d", len(id), length)
	}
}

// test for checking the length of alphanumeric id
func TestLengthalphanumericID(t *testing.T) {
	length := rand.Intn(100)
	id, err := Generateid("l", length)
	if err != nil {
		t.Error("error while generating id.", err)
	}
	if (len(id)) != length {
		t.Errorf("got %d, want %d", len(id), length)
	}
}

// test for checking type
func TestType(t *testing.T) {
	length := rand.Intn(100)
	id, err := Generateid("n", length)
	if err != nil {
		t.Error("error while generating id.", err)
	}
	idarray := strings.Split(id, "")
	for _, v := range idarray {
		_, err := strconv.Atoi(v)
		if err != nil {
			t.Error("id is not number only")
		}
	}
}

// test for checking uniquness of number ID
func TestUniquenumberID(t *testing.T) {
	totalIDs := 100
	length := rand.Intn(100)
	generated := make(map[string]bool)
	for i := 0; i < totalIDs; i++ {
		id, err := Generateid("n", length)
		if err != nil {
			t.Error("error while generating id.", err)
		}
		if generated[id] {
			t.Error("id already created.")
		}
		generated[id] = true
	}
}

//test for checking uniquness of letter ID
func TestUniqueletterID(t *testing.T) {
	totalIDs := 100
	length := rand.Intn(100)
	generated := make(map[string]bool)
	for i := 0; i < totalIDs; i++ {
		id, err := Generateid("l", length)
		if err != nil {
			t.Error("error while generating id.", err)
		}
		if generated[id] {
			t.Error("id already created.")
		}
		generated[id] = true
	}
}

//test for checking uniquness of alphanumeric ID
func TestUniquealphanumericID(t *testing.T) {
	totalIDs := 100
	length := rand.Intn(100)
	generated := make(map[string]bool)
	for i := 0; i < totalIDs; i++ {
		id, err := Generateid("a", length)
		if err != nil {
			t.Error("error while generating id.", err)
		}
		if generated[id] {
			t.Error("id already created.")
		}
		generated[id] = true
	}
}
