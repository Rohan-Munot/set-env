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
		setEnv := os.Getenv("SWITCH_ENV")
		if setEnv == "" {
			setEnv = "ENV not set yet!"
		}
		fmt.Println(setEnv)
		return
	}

	if len(args) > 1 {
		fmt.Println("Too many arguments provided")
		os.Exit(1)
	}

	envArg := strings.TrimSpace(args[0])
	fileName := fmt.Sprintf(".env.%s", envArg)

	content, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File %s not found\n", fileName)
		} else {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
		}
		os.Exit(1)
	}

	for line := range strings.SplitSeq(string(content), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		fmt.Printf("export %s=%s\n", strings.TrimSpace(key), strconv.Quote(strings.TrimSpace(value)))
	}
	fmt.Printf("export SWITCH_ENV=%s\n", strconv.Quote(envArg))

}
