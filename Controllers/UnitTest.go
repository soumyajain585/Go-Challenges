package Controllers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func UnitTest() {
	fmt.Println("Enter a number to check wheather it is prime or not :")
	doneChan := make(chan bool)
	go readUserInput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)
	fmt.Println("GoodBye!")
}

func readUserInput(i io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(i)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- false
			return
		}
		fmt.Println(res)
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
