package tree

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func List(path string) (Nodes, error) {
	response, err := os.ReadDir(path)

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		return nil, &PathError{Message: fmt.Sprintf("\"%s\" is not a directory", path)}
	}

	entries := Nodes{}
	for _, entry := range response {
		entries = append(
			entries,
			Node{
				Name:  entry.Name(),
				IsDir: entry.IsDir(),
			},
		)
	}

	return entries, nil
}

func Chdir(path string) error {
    err := os.Chdir(path)

    var pathError *os.PathError
	if errors.As(err, &pathError) {
	    return &PathError{Message: fmt.Sprintf("\"%s\" is not a directory", path)}
	}

	return nil
}

func ReadNodeLineByLine(path string, hook LineHook, counter *uint64) {
	file, _ := os.Open(path)
	defer file.Close() // nolint:errcheck

	scanner := bufio.NewScanner(file)

	// go line by line
	for scanner.Scan() {
		hook(scanner.Text(), counter)
	}
}
