package set1

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestSolveQ4(t *testing.T) {
	file, err := os.Open("question4_data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	output, err := SolveQ4(scanner)
	if err != nil {
		t.Errorf("Got an unexpected error %v", err)
	}

	expectedOutput := []byte{78, 111, 119, 32, 116, 104, 97, 116, 32, 116, 104, 101, 32, 112, 97, 114, 116, 121, 32,
		105, 115, 32, 106, 117, 109, 112, 105, 110, 103, 10}
		
	if output != string(expectedOutput) {
		t.Errorf("Expected output: %v obtained output: %v", expectedOutput, output)
	}
}
