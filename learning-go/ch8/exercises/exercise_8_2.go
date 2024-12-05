package exercises

import (
	"fmt"
	"strconv"
)

/*
Define a generic interface called Printable that matches a type that
implements fmt.Stringer and has an underlying type of int or float64.
Define types that meet this interface. Write a Printable and prints
its value to the screen using fmt.Println
*/

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintablePhone int

type RoundedFloat float64

func (pp PrintablePhone) String() string {
	str := strconv.Itoa(int(pp))
	fmtStr := fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:])
	return fmtStr
}

func (rf RoundedFloat) String() string {
	fmtStr := fmt.Sprintf("%.2f (rounded to two decimal places)", rf)
	return fmtStr
}

func Print[P Printable](p P) {
	fmt.Println(p)
}

func Exercise_8_2() {
	var pp PrintablePhone = 4081234567
	Print(pp)

	var rf RoundedFloat = 123456789.123456
	Print(rf)
}
