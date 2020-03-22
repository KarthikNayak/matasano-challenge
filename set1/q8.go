package set1

import (
	"bufio"
	"os"
)

func Q8(filename string) bool {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		data := fscanner.Text()
		for i := 0; i <= len(data); i += 16 {
			for j := i; j <= len(data); j += 16 {
				if i == j {
					return true
				}
			}
		}
	}
	return false
}
