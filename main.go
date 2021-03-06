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
	var args []string
	if os.Args[1] == "test" {
		// run with "delver test" removed these
		args = os.Args[2:]
	} else {
		// run with "delver" removed this
		args = os.Args[1:]
	}

	cmd, err := getCmd(args)
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
