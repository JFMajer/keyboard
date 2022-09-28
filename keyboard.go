// Package keyboard reads user input from the keyboard.
package keyboard

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a floating point number from the keyboard.
// It returns number and any error encountered.
func GetFloat() (float64, error) {
	fmt.Print("Please input your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatal(err)
	}
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return grade, nil
}
