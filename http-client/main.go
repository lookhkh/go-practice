package main

import (
	"fmt"
	"os"

	"github.com/lookhkh/data-downlodaer/pkgquery"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must Specify URL")
		os.Exit(1)
	}

	body, err := pkgquery.FetchPackageData(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%v\n", body)

}
