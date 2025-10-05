package tree

import "fmt"

type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("%s is not a directory", e.Path)
}
