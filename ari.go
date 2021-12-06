package main

import (
	"ari/par"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "%s <input file>\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	in := flag.Arg(0)
	if in == "" {
		Usage()
	}

	infile, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(infile)
	for sc.Scan() {
		line := sc.Text()
		ast := par.Parse(line)
		// ast.PrintTree() // for debugging
		fmt.Println(line, "=", ast.Interpret())
	}
}
