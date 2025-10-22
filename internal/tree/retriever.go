package tree

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Tree struct {
    WorkDir string
}

func (t *Tree) Copy() Tree {
    return Tree{WorkDir: t.WorkDir}
}

func (t *Tree) List(path string) (Nodes, error) {
	response, err := os.ReadDir(filepath.Join(t.WorkDir, path))

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

func (t *Tree) Chdir(path string) error {
    fullPath := filepath.Join(t.WorkDir, path)
    dir, err := os.Open(fullPath)

    var pathError *os.PathError
	if errors.As(err, &pathError) {
	    return &PathError{Message: fmt.Sprintf("\"%s\": no such file or directory", path)}
	}

	stat, _ := dir.Stat()

	if !stat.IsDir() {
	    return &PathError{Message: fmt.Sprintf("\"%s\" is not a directory", path)}
	}

	t.WorkDir = fullPath
	return nil
}

func (t *Tree) ReadNodeLineByLine(path string, hook LineHook, counter *uint64) {
	file, _ := os.Open(filepath.Join(t.WorkDir, path))
	defer file.Close() // nolint:errcheck

	scanner := bufio.NewScanner(file)

	// go line by line
	for scanner.Scan() {
		hook(scanner.Text(), counter)
	}
}
