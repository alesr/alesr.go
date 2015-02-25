package main

//#include <unistd.h>
//#include <sys/types.h>
import "C"
import "fmt"
import "os"

func main() {

	fact := 1

	for i := 5; i > 1; i-- {

		fact = fact * i
		fmt.Println("Factorial -> ", fact)

		x := C.fork()

		if x > 0 {
			os.Exit(0)
		}
	}
}
