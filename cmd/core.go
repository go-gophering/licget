package cmd

import (
	"github.com/go-gophering/licget/licget"
	"os"
)

type arguments struct {
	license_name string
	list_all     bool
}

func parse_args(args []string) arguments {
	var list_all bool
	var license_name string
	switch args[0] {
	case "all":
		list_all = true
	case "list":
		list_all = true
	case "":
		println("Error, an argument is required.")
	default:
		license_name = args[0]
	}
	return arguments{
		license_name: license_name,
		list_all:     list_all,
	}
}

func Execute() {
	args := parse_args(os.Args[1:])
	licget.Run(args.license_name, args.list_all)
}
