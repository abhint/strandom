package strandom_test

import (
	"fmt"
	"testing"

	"github.com/abhint/strandom"
)

func TestRandomString(t *testing.T) {
	randomString := strandom.RandomString(10)
	if len(randomString) != 10 {
		t.Errorf("Expected string length 10, but got %d", len(randomString))
	}
	fmt.Println(randomString)
}

func BenchmarkRandomString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strandom.RandomString(1)
	}
}
