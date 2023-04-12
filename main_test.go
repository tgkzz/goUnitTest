package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
		message  string
	}{
		{0, false, "0 is not prime, by definition!"},
		{1, false, "1 is not prime, by definition!"},
		{-1, false, "Negative numbers are not prime, by definition!"},
		{2, true, "2 is a prime number!"},
		{3, true, "3 is a prime number!"},
		{4, false, "4 is not a prime number because it is divisible by 2!"},
		{5, true, "5 is a prime number!"},
		{6, false, "6 is not a prime number because it is divisible by 2!"},
		{7, true, "7 is a prime number!"},
		{8, false, "8 is not a prime number because it is divisible by 2!"},
		{9, false, "9 is not a prime number because it is divisible by 3!"},
		{10, false, "10 is not a prime number because it is divisible by 2!"},
	}

	for _, test := range tests {
		actual, message := isPrime(test.input)
		if actual != test.expected || message != test.message {
			t.Errorf("isPrime(%d) = (%v, %s), expected (%v, %s)", test.input, actual, message, test.expected, test.message)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "Please enter a whole number!"},
		{"0", "0 is not prime, by definition!"},
		{"1", "1 is not prime, by definition!"},
		{"2", "2 is a prime number!"},
		{"31", "31 is a prime number!"},
		{"4", "4 is not a prime number because it is divisible by 2!"},
		{"-2", "Negative numbers are not prime, by definition!"},
		{"4.0", "Please enter a whole number!"},
		{"q", ""},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		scanner := bufio.NewScanner(input)
		res, _ := checkNumbers(scanner)

		if !strings.EqualFold(res, e.expected) {
			t.Errorf("Expected: %s. But got: %s", e.expected, res)
		}
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	//perform test
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Incorrect intro text, got %s", string(out))
	}
}

func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error creating pipe: %v", err)
	}

	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = oldOut

	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("error reading from pipe: %v", err)
	}

	expected := "-> "
	if string(output) != expected {
		t.Errorf("unexpected output:\nexpected: %q\nactual:   %q", expected, output)
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
