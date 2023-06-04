package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

var NoNameError = errors.New("No name! Enter Your name")
var InvalidArgs = errors.New("Invalid Number of Args")
var ArgsLessThanZero = errors.New("Must speckify a number grater than 0")

func validateArgs(c config) error {
	if !c.printUsage && !(c.numTimes > 0) {
		return ArgsLessThanZero
	}
	return nil
}

func parseArgs(args []string) (config, error) {

	var numTimes int
	var err error

	c := config{}

	if len(args) != 1 {
		return c, InvalidArgs
	}
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}

	c.numTimes = numTimes

	return c, err
}

func GetName(r io.Reader, w io.Writer) (string, error) {

	msg := "Your Name Please? Press Enter When Done.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", NoNameError
	}

	return name, nil

}

func printUsage(w io.Writer) {
	usage := fmt.Sprintf("Usage : %s [-h --help] a  num of Times", os.Args[0])
	fmt.Fprintf(w, usage)
}

func runCmd(r io.Reader, w io.Writer, c config) error {

	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := GetName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Noce to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintln(w, msg)
	}
}

func main() {

	fmt.Printf("Hello World %v\n", "hi")

	c, err := parseArgs(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = validateArgs(c)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	// name, err := GetName(os.Stdin, os.Stdout)
	// if err != nil {
	// 	fmt.Fprintf(os.Stdout, "Erorr : %v\n", err)
	// 	os.Exit(1)
	// }

}
