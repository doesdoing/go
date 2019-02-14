package check

import "fmt"

/*
CHECK...
*/
func Check(a error) {
	if a != nil {
		fmt.Println(a)
	}
}
