package cmd

import (
	"flag"
	"fmt"
	"io"
)

func HandleCmdB(stdOut io.Writer, s []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(stdOut)
	fs.StringVar(&v, "verb", "argument-value", "argument 1")
	err := fs.Parse(s)

	if err != nil {
		return err
	}
	fmt.Fprint(stdOut, "Executing b\n")
	return nil
}

func HandleCmdA(stdOut io.Writer, s []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-a", flag.ContinueOnError)
	fs.SetOutput(stdOut)
	fs.StringVar(&v, "verb", "argument-value", "argument 1")
	err := fs.Parse(s)

	if err != nil {
		return err
	}
	fmt.Fprint(stdOut, "Executing a\n")
	return nil

}
