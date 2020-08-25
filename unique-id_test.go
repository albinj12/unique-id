package uniqueid

import (
	"math/rand"
	"testing"
)

// test for checking the length of generated id
func TestLength(t *testing.T) {
	length := rand.Intn(100)
	id, err := Generateid("n", length)
	if err != nil {
		t.Error("error while generating id.", err)
	}
	if (len(id)) != length {
		t.Errorf("got %d, want %d", len(id), length)
	}
}

// test for checking uniquness
func TestUnique(t *testing.T) {
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
