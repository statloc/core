package statloc

import (
	"errors"

	"github.com/statloc/core/internal/retrievers/mapping"
	"github.com/statloc/core/internal/retrievers/tree"
)

func GetStatistics(path string) (*StatisticsResponse, error) {
    mapping.Load(
        "assets/components.json",
        "assets/extensions.json",
    )

    list, err := tree.List(path)

	var treePathError *tree.PathError
	if errors.As(err, &treePathError) {
		return nil, &PathError{Path: path}
	}

	items := make(map[string]*TableItem)

	for _, value := range mapping.Components {
        items[value] = &TableItem{Files: 0, LOC: 0}
	}
	for _, value := range mapping.Extensions {
        items[value] = &TableItem{Files: 0, LOC: 0}
	}

	statistics := &StatisticsResponse{ Items: items }

	goAroundCalculating(list, statistics, nil)

	total := TableItem{Files: 0, LOC: 0}
	for title, item := range statistics.Items {
	    if item.Files == 0 {
			delete(statistics.Items, title)
		} else {
            total.Append(item.LOC, item.Files)
		}
	}

	return statistics, nil
}
