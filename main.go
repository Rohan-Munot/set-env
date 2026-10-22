package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No arguments provided")
		fmt.Fprintln(os.Stderr, "Usage: se <environment> or se current to show the current environment")
		fmt.Fprintln(os.Stderr, "For more information: se help")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "Too many arguments provided")
		os.Exit(1)
	}

	if args[0] == "current" {
		currentEnv := os.Getenv("SWITCH_ENV")
		if currentEnv == "" {
			fmt.Fprintln(os.Stderr, "No environment set yet")
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Current environment: %s\n", currentEnv)
		os.Exit(0)
	}

	if args[0] == "help" {
		fmt.Fprintln(os.Stderr, "se — load .env.<name> into the current shell")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintln(os.Stderr, "  se <name>     apply .env.<name> (e.g. se local → .env.local)")
		fmt.Fprintln(os.Stderr, "  se current    show which env is active in this shell")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Vars last until the shell exits. A new terminal starts unset.")
		os.Exit(0)
	}

	envArg := strings.TrimSpace(args[0])
	fileName := fmt.Sprintf(".env.%s", envArg)

	content, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintf(os.Stderr, "\033[31mFile %s not found\n\033[0m", fileName)
		} else {
			fmt.Fprintf(os.Stderr, "\033[31mError reading file %s: %v\n\033[0m", fileName, err)
		}
		os.Exit(1)
	}

	// unset all the keys in the current environment
	for k := range strings.FieldsSeq(os.Getenv("SWITCH_ENV_KEYS")) {
		fmt.Printf("unset %s\n", k)
	}

	var keys []string

	for line := range strings.SplitSeq(string(content), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			fmt.Fprintf(os.Stderr, "\033[31mError parsing line %s\n\033[0m", line)
			fmt.Fprintln(os.Stderr, "Each line must be in the format KEY=VALUE")
			os.Exit(1)
		}
		key = strings.TrimSpace(key)
		keys = append(keys, key)
		fmt.Printf("export %s=%s\n", key, strconv.Quote(strings.TrimSpace(value)))

	}
	fmt.Printf("export SWITCH_ENV=%s\n", strconv.Quote(envArg))
	fmt.Printf("export SWITCH_ENV_KEYS=%s\n", strconv.Quote(strings.Join(keys, " ")))
	fmt.Fprintf(os.Stderr, "\033[32menv %s set successfully!\n\033[0m", envArg)

}
