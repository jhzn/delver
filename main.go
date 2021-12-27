package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func printHelp() {
	text := `delver is a program which takes the same arguments as 'go test' and starts dlv(delve) with those arguments.

Example:
	delver test -v -count=1 -run '^TestMyFunc$' ./pkg/api/tests

Note that you first have to make sure that the command works with 'go test'

If you encounter any issue, raise an issue on github: https://github.com/jhzn/delver
`
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
	}
	switch os.Args[1] {
	case "-h", "--help":
		printHelp()
	}
	//remove program name("delver") and 1st arg("test")
	cmd, err := getCmd(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delver is running cmd: \n%v\n\n", cmd)
	proc, err := runCmd(cmd...)
	if err != nil {
		log.Fatal(err)
	}
	proc.Wait()
}

func runCmd(args ...string) (*os.Process, error) {
	arg, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}
	args[0] = arg

	procAttr := os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}
	p, err := os.StartProcess(args[0], args, &procAttr)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func mergeSlices(a, b []string) []string {
	merged := []string{}
	for _, a := range a {
		merged = append(merged, a)
	}
	for _, a := range b {
		merged = append(merged, a)
	}
	return merged
}

func getCmd(flags []string) ([]string, error) {
	pkgPath := flags[len(flags)-1]

	goTestArgs := []string{}
	for _, f := range flags[0 : len(flags)-1] {
		goTestArgs = append(goTestArgs, strings.Replace(f, "-", "-test.", 1))
	}
	goTestArgs = append(goTestArgs, pkgPath)

	delveArgs := []string{
		"dlv",
		"test",
		fmt.Sprintf("--build-flags=%s", pkgPath),
		"--",
	}

	return mergeSlices(delveArgs, goTestArgs), nil
}
