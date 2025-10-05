package tree

type (
	Node struct {
		Name  string
		IsDir bool
	}

	LineHook func(
		text string,
		counter *uint64,
	)

	ListResponse struct {
		Nodes []Node
	}
)
