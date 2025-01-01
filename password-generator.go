package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	passwordhandler "github.com/sahay-shashank/password-generator-golang/handler"
)

func applicationHandler(option *string, lengthptr *int) {
	opt := strings.ToLower(*option)
	switch opt {
	case "web":
		passwordhandler.Web()
	case "cli":
		passwordhandler.Cli(lengthptr)
	}

}

func main() {
	option := flag.String("option", "", "The Option for usage")
	lengthptr := flag.Int("length", 7, "Length of the password")

	flag.Parse()
	if *option == "" {
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, `
			Usage %s -option=[cli | web] -length=[length in integer greater than equal to 7]

			If option is web, length will not be parsed in the cli. The length should be 
			provided in the url as an argument of 'length'.

			'length' is defaulted to 7 in both the options.

			Character set includes:
			- abcdefghijklmnopqrstuvwxyz
			- ABCDEFGHIJKLMNOPQRSTUVWXYZ
			- 0123456789
			- !@#$%

			`, os.Args[0])
			flag.PrintDefaults()
		}
		os.Exit(1)
	}

	applicationHandler(option, lengthptr)
}
