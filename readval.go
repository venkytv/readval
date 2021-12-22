package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

const Usage = `usage: %s [OPTIONS] prompt...
Prints prompts, reads input from TTY, and echoes it back on stdout.

  -s, --silent		Do not echo input
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0])
	}

	var silent bool
	flag.BoolVar(&silent, "s", false, "do not echo input")
	flag.BoolVar(&silent, "silent", false, "do not echo input")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	in, err := os.Open("/dev/tty")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0700)
	if err != nil {
		log.Fatal(err)
	}

	prompt := strings.Join(flag.Args(), " ") + ": "
	fmt.Fprint(out, prompt)

	var val string
	if silent {
		v, err := term.ReadPassword(int(in.Fd()))
		fmt.Fprint(out, "\n")
		if err != nil {
			log.Fatal(err)
		}
		val = string(v) + "\n"
	} else {
		reader := bufio.NewReader(in)
		val, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print(val)
}
