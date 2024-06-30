package main

import (
	"bufio"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, _ := os.ReadDir(dir)
	env := make(Environment)
	var err error

	for _, e := range dirEntries {
		if e.IsDir() {
			continue
		}

		func() {
			f, innerErr := os.Open(dir + "/" + e.Name())
			defer func() {
				if err = f.Close(); err != nil {
					return
				}
			}()

			scanner := bufio.NewScanner(f)
			scanner.Scan()
			value := sanitize(scanner.Text())
			env[e.Name()] = EnvValue{
				Value:      value,
				NeedRemove: false,
			}

			if innerErr != nil {
				err = innerErr
			}
		}()

		if err != nil {
			return nil, err
		}
	}

	return env, nil
}

func sanitize(s string) string {
	s = strings.ReplaceAll(s, "\x00", "\n")
	s = strings.TrimRight(s, " \t")
	return s
}
