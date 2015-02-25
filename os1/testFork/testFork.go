package main

//#include <unistd.h>
//#include <sys/types.h>
import "C"
import "fmt"

func main() {
	testFork()
}

func testFork() {
	fmt.Println("TESTING FORK()")

	pid := C.fork()

	if pid > 0 {
		fmt.Println(" Parent's process PID: ", pid)
	} else {
		fmt.Println(" Child's  process PID: ", pid)
	}
	fmt.Println("END")
}
