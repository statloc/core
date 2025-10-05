package statloc

type (
    TableItem struct {
        LOC   uint64
        Files uint64
    }

    StatisticsResponse struct {
        Items map[string]*TableItem // item's title as a key
    }
)

func (t *TableItem) Append(LOC uint64, files uint64) {
	t.LOC += LOC
	t.Files += files
}
