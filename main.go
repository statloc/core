package statloc

import (
	_ "embed"
	"errors"
	"path/filepath"

	"github.com/statloc/core/internal/mapping"
	"github.com/statloc/core/internal/tree"
)

var (
    //go:embed "assets/extensions.json"
    rawExtensions string

    //go:embed "assets/components.json"
    rawComponents string
)

func GetStatistics(path string) (*StatisticsResponse, error) {
    mapping.Load(rawComponents, rawExtensions)

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

	items["Total"] = &TableItem{Files: 0, LOC: 0}

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

func goAroundCalculating(
	list               tree.ListResponse,
	existingStatistics *StatisticsResponse,
	component          *string,
) {
	for _, node := range list.Nodes {
		if node.IsDir {
            componentType, exists := mapping.Components[filepath.Base(node.Name)]

            if exists { component = &componentType }

            list, _ = tree.List(node.Name)
			goAroundCalculating(list, existingStatistics, component)

		} else {
            LOC := uint64(1)
            tree.ReadNodeLineByLine(node.Name, proceedLine, &LOC)

            existingStatistics.Items["Total"].Append(LOC, 1)

            fileType, exists := mapping.Extensions[filepath.Ext(node.Name)]
            if exists {
                existingStatistics.Items[fileType].Append(LOC, 1)
            }

            if component != nil {
                existingStatistics.Items[*component].Append(LOC, 1)
            }
		}
	}
}

func proceedLine(text string, counter *uint64) {
	*counter++
}
