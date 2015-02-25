package main

//#include <unistd.h>
//#include <sys/types.h>
import "C"
import "fmt"

func main() {
	pid := C.fork()

	if pid == 0 {
		childProcess()
	} else {
		parentProcess()
	}
}

func childProcess() {
	for i := 1; i <= 200; i++ {
		fmt.Println("This line is from child, value = ", i)
	}
	fmt.Println(" *** Child process is done *** ")
}

func parentProcess() {
	for i := 1; i <= 200; i++ {
		fmt.Println("This line is from parent, value = ", i)
	}
	fmt.Println(" *** Parent process is done *** ")
}
