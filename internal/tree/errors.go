package tree

type PathError struct {
	Message string
}

func (e *PathError) Error() string {
	return e.Message
}
