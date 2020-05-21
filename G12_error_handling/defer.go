package G12_error_handling

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func TestDefer() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f) // run at the end of main
	writeFile(f)
}


func c() (i int) {
	defer func() { i++ }()
	return 1
}

func DeferSideEffect() {
	println(c()) // 2
}


