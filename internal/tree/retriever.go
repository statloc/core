package tree

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func List(path string) (ListResponse, error) {
	response, err := os.ReadDir(path)

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		return ListResponse{}, &PathError{Path: path}
	}

	entries := []Node{}
	for _, entry := range response {
		entries = append(
			entries,
			Node{
				Name:  filepath.Join(path, entry.Name()),
				IsDir: entry.IsDir(),
			},
		)
	}

	return ListResponse{Nodes: entries}, nil
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
