package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/charmbracelet/huh"
)

func main() {
	config, err := getConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	entries, err := parseSshConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var host string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions(entries...)...).
				Title("Select host").
				Value(&host),
		),
	).WithTheme(huh.ThemeBase())

	err = form.Run()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("connecting to host " + host)

	sshBinary, err := exec.LookPath("ssh")
	if err != nil {
		fmt.Println(errors.Join(
			err,
			errors.New("ssh binary not found"),
		))
		os.Exit(1)
	}

	err = syscall.Exec(sshBinary, []string{"ssh", host}, os.Environ())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getConfig() ([]byte, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("cannot get home dir"),
		)
	}

	config, err := os.ReadFile(dirname + "/.ssh/config")
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("cannot read config"),
		)
	}

	return config, nil
}
