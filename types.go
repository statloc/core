package statloc

type (
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

func (t *TableItem) Append(LOC uint64, files uint64) {
	t.LOC += LOC
	t.Files += files
}
