package fuzzing

type Add interface {
	Test(int) int
}

type Adder struct {
	data int
}

func (a *Adder) Test(n int) int {
	a.data += n
	return a.data
}
