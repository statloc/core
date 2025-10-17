package statloc

type (
    component struct {
        Title string
        Prev  *component
    }

    componentSet struct {
        Tail     *component
        Elements map[string]struct{}
    }

    TableItem struct {
        LOC   uint64
        Files uint64
    }

    Items map[string]*TableItem

    Statistics struct {
        Languages  Items
        Components Items
        Total      TableItem
    }
)

func (s *componentSet) Add(title string) *component {
    s.Elements[title] = struct{}{}
    s.Tail = &component{Title: title, Prev: s.Tail}
    return s.Tail
}

func (s *componentSet) Pop() {
    delete(s.Elements, s.Tail.Title)
    s.Tail = s.Tail.Prev
}

func (s *componentSet) In(title string) bool {
    _, exists := s.Elements[title]
    return exists
}

func (t *TableItem) Append(LOC uint64, files uint64) {
	t.LOC += LOC
	t.Files += files
}
