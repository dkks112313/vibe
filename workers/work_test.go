package workers

import (
	"testing"
)

func Test(t *testing.T) {
	workers := Workers{Count: 10}
	result := workers.ReadFile("hello.txt")

	if len(result) == 0 {
		t.Error(result)
	}
}
