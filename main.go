package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get user input as length of prime list
	userInput := bufio.NewScanner(os.Stdin)

	// when user enters the number
	for userInput.Scan()  {
		// convert input string to integer
		listNumber, _ := strconv.Atoi(userInput.Text())

		// create a fibbonaci array
		var fib [] int

		// populate array with n length
		for i := 0; i < listNumber; i++ {
			// value in index 0 to 1 has the same value as its index
			if i <= 1 {
				fib = append(fib, i)
			} else {
				// append new fibbonaci number to fib array
				number := fib[i-1] + fib[i-2]
				fib = append(fib, number)
			}
		}

		// print fib array
		fmt.Println(fib)

	}
}
