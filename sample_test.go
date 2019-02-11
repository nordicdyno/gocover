package gocover_test

import (
	"testing"

	"github.com/nordicdyno/gocover"
)

func TestConditionalAdd(t *testing.T) {
	result1 := gocover.ConditionalAdd(100, false)
	if result1 != 100+5 {
		t.Errorf("got unexpected result: %v", result1)
	}

	// result2 := gocover.ConditionalAdd(100, true)
	// if result1 != 100+500 {
	// 	t.Errorf("got unexpected result: %v", result2)
	// }
}
