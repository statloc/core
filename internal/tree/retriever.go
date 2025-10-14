package tree

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func List(path string) (ListResponse, error) {
	response, err := os.ReadDir(path)

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		return nil, &PathError{Message: fmt.Sprintf("%s is not a directory", path)}
	}

	entries := ListResponse{}
	for _, entry := range response {
		entries = append(
			entries,
			Node{
				Name:  filepath.Join(path, entry.Name()),
				IsDir: entry.IsDir(),
			},
		)
	}

	return entries, nil
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
