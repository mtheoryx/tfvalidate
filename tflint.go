package main

import (
	"flag"
	"fmt"
	"github.com/justinm/tflint/tflint"
	"log"
	"os"
)

func getCwd() string {
	return os.Getenv("PWD")
}

func main() {
	pathToPlan := flag.String("plan", "", "Path to the plan to lint")
	pathToLint := flag.String("lint", "", "Path to lint rules, defaults to $CWD/.tflint.json")

	flag.Parse()

	if len(*pathToPlan) == 0 {
		log.Fatal("--plan must be supplied")
	}

	if len(*pathToLint) == 0 {
		*pathToLint = getCwd() + "/.tflint.json"
	}

	parser := tflint.Parser{}

	ruleset, err := parser.Parse(*pathToLint)
	if err != nil {
		log.Fatal(err)
	}

	plan, err := tflint.OpenPlan(*pathToPlan)
	if err != nil {
		log.Fatal("Failed to open plan at " + *pathToPlan)
	}

	linter := tflint.Linter{}

	linter.Lint(plan, ruleset)

	if len(linter.Violations) == 0 {
		fmt.Print("No errors found")
	} else {
		for _, violation := range linter.Violations {
			fmt.Printf("Violation: %s", violation.Reason)
		}
	}
}
