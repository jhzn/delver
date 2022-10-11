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
	case "", "-h", "--help":
		printHelp()
	}

	dlvCmv, err := getCmd(cleanArgs(os.Args))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delver is running cmd: \n%v\n\n", dlvCmv)
	proc, err := runCmd(dlvCmv...)
	if err != nil {
		log.Fatal(err)
	}
	proc.Wait()
}

// cleanArgs cleans up the argument list to delver
func cleanArgs(args []string) []string {
	// program was run with "delver test ..."
	if args[1] == "test" {
		return args[2:]
	}
	// program was run with "delver ..."
	return args[1:]
}

func runCmd(args ...string) (*os.Process, error) {
	dlvBinary, err := exec.LookPath(args[0])
	if err != nil {
		return nil, err
	}
	// Override binary to be absolute path
	args[0] = dlvBinary

	procAttr := os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}
	p, err := os.StartProcess(dlvBinary, args, &procAttr)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func getCmd(flags []string) ([]string, error) {
	lastIndex := len(flags) - 1
	pkgPath := flags[lastIndex]

	goTestArgs := []string{}
	for _, f := range flags[0:lastIndex] {
		goTestArgs = append(goTestArgs, strings.Replace(f, "-", "-test.", 1))
	}
	goTestArgs = append(goTestArgs, pkgPath)

	buildArgs := `-gcflags="all=-N -l"`
	delveArgs := []string{
		"dlv",
		"test",
		fmt.Sprintf(`--build-flags='%s' %s`, buildArgs, pkgPath),
		"--",
	}
	// Merge slices
	return append(delveArgs, goTestArgs...), nil
}
