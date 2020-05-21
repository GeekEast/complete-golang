package G12_error_handling

import (
	"errors"
	"fmt"
)


func f() (int,*MyError) {
	return 1, &MyError{}  // must be pointer
}

func g() (int,error) {
	return 1, errors.New("general error")
}

type MyError struct{}
func (e MyError) Error() string {
	return "Hello Error!"
}

func TestMakeError() {
	_,err := f()
	_,e := g()
	fmt.Println(err)
	fmt.Println(e)
}