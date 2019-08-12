package fuzzing

import (
	fuzz "github.com/google/gofuzz"
	"testing"
)

func TestAdder_Test(t *testing.T) {

	fezzes := fuzz.New()

	var num int
	fezzes.Fuzz(&num)

	adder := Adder{}

	adder.Test(num)

}
