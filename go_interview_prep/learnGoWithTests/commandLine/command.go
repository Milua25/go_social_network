package main

import (
	"flag"
	"fmt"
	"os"
)

func customUsage(){
	// Get the output destination (default is os.Stderr)
	w := flag.CommandLine.Output() 
	
	// Print a custom usage message preamble
	fmt.Fprintf(w, "Usage: %s [OPTIONS] <command>\\n", os.Args[0])
	fmt.Fprintf(w, "This is a custom help message for the application.\\n\\n")

	// Print the default values for all defined flags
	fmt.Fprintf(w, "Options:\\n")
	flag.PrintDefaults() 

	// Print a custom usage message postamble
	fmt.Fprintf(w, "\\nCommands:\\n")
	fmt.Fprintf(w, "  start    Start the service\\n")
	fmt.Fprintf(w, "  stop     Stop the service\\n")
}

func main() {
	//	args := os.Args[0]

	// fmt.Println(args)

	// for _, arg := range args {
	// 	fmt.Println(arg)
	// }
	// var name string
	// var age int
	// var male bool

	// flag.StringVar(&name, "name", "john", "Name of the user!")
	// flag.IntVar(&age, "age", 0, "enter your age")
	// flag.BoolVar(&male, "male", true, "Gender of the user")

	// flag.Parse()

	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Male:", male)

	subcommand1 := flag.NewFlagSet("firsSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "enter is the status is processing:")
	secondFlag := subcommand1.Int("bytes", 1024, "result byte length")

	flagSc2 := subcommand2.String("language", "go", "enter your language")

	flag.Parse()
	

	if flag.NArg() < 2{
		fmt.Println("Program requires additional commands.")
		//  = customUsage // call the func to print the default usage instead of its pointer
		os.Exit(1)
	}

	switch flag.Args()[1]{
	case "firstSub":
		subcommand1.Parse(flag.Args()[2:])
		fmt.Println("subCommand 1")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(flag.Args()[2:])
		fmt.Println("subCommand 2")
		fmt.Println("language:", *flagSc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}


}
