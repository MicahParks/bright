package main

import (
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"
)

var (
	max = big.NewFloat(100)
	min = big.NewFloat(20)
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	var brightness *big.Float
	var err error
	const usage = "Argument must be a number between 20 and 100."

	if len(os.Args) == 1 {
		brightness = big.NewFloat(1)
	} else if len(os.Args) != 2 {
		logger.Fatalln(usage)
	} else {
		brightness, _, err = big.ParseFloat(os.Args[1], 10, 0, big.ToPositiveInf)
		if err != nil {
			logger.Println(usage)
			logger.Fatalf("Error: %s", err.Error())
		}

		if brightness.Cmp(min) == -1 || brightness.Cmp(max) == 1 {
			logger.Fatalf("Must be between %s and %s.", min.String(), max.String())
		}

		brightness.Mul(brightness, big.NewFloat(.01))
	}

	out, err := exec.Command("xrandr", "--listactivemonitors").Output()
	if err != nil {
		logger.Fatalf("Failed to list active monitors.\nError: %s", err.Error())
	}

	errMsg := "Can't find monitor names."

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		logger.Fatalln(errMsg)
	}

	displayNames := make([]string, 0)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		split := strings.Split(line, "  ")
		if len(split) != 2 {
			logger.Fatalln(errMsg)
		}
		displayNames = append(displayNames, split[1])
	}

	for _, display := range displayNames {
		err = exec.Command("xrandr", "--output", display, "--brightness", brightness.String()).Run()
		if err != nil {
			logger.Fatalln("Failed to set brightness: %s", err.Error())
		}
	}

	logger.Println("Brightness set.")
}
