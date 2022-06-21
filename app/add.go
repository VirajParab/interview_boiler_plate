package main

type number struct {
	n int
}

func (num number)add(otherNumber number) number {
	return number{
		n: otherNumber.n + num.n,
	}
}
