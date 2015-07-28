package lineReverse

import (
	"bufio"
	"log"
	"os"
	"testing"
)

var testCase = []struct {
	line, expected string
}{
	{"When you chase a dream, especially one with plastic chests, you sometimes do not see what is right in front of you.",
		".uoy fo tnorf ni thgir si tahw ees ton od semitemos uoy ,stsehc citsalp htiw eno yllaicepse ,maerd a esahc uoy nehW"},
}

func TestLineReverse(t *testing.T) {

	LineReverse("input.txt", "output.txt")

	file, err := os.Open("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var observed string
	for scanner.Scan() {
		observed = scanner.Text()
	}

	var expected = ".uoy fo tnorf ni thgir si tahw ees ton od semitemos uoy ,stsehc citsalp htiw eno yllaicepse ,maerd a esahc uoy nehW"
	if observed != expected {
		t.Errorf("Expected %s, got, %s.", expected, observed)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
