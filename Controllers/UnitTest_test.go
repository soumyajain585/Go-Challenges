package Controllers

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// Commands

// To test code - go test .
// To check coverage - go test -cover .
// To create cover file - go test -coverprofile=coverage.out
// To check cover code in html - go tool cover -html=coverage.out

func Test_isPrime01(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 7)
	}

	if msg != "7 is a prime number!" {
		t.Error("wrong message returned:", msg)
	}
}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition!"},
		{name: "one", input: "1", expected: "1 is not prime, by definition!"},
		{name: "two", input: "2", expected: "2 is a prime number!"},
		{name: "three", input: "3", expected: "3 is a prime number!"},
		{name: "negative", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "1.1", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "επτά", expected: "Please enter a whole number!"},
	}
	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, tt.expected) {
			t.Errorf("%s: expected %s but got %s", tt.name, tt.expected, res)
		}
	}
}
func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
