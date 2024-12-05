package exercises

import "fmt"

/*
	write two functions. the `updateslice` function takes in a
	`[]string` and a `string`. it sets the last position in the
	passed-in slice to the passed-in `string`. at the end of
	`updateslice`, print the slice after making the change. the
	`growslice` function also takes in a `[]string` and a `string`. it
	appends the `string` onto the slice. at the end of `growslice`,
	print the slice after making the change. call these functions from
	`main`. print out the slice before each function is called and
	after each function is called. do you understand why some changes
	are visible in `main` and why some changes are not?
*/

func UpdateSlice(slice []string, s string) {
	slice[len(slice)-1] = s
	fmt.Println("post-update in func", slice)
}

func GrowSlice(slice []string, s string) {
	slice = append(slice, s)
	fmt.Println("post-append in func", slice)
}

func Exercise_6_2() {
	testSlice := []string{"TEST STRING 01", "TEST STRING 02", "TEST STRING 03"}

	fmt.Println("Before update:", testSlice)
	UpdateSlice(testSlice, "UPDATED")
	fmt.Println("After update:", testSlice)

	GrowSlice(testSlice, "APPENDED")
	fmt.Println("After append:", testSlice)
}
