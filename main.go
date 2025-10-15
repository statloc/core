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

	languages := make(Items)
	components := make(Items)
	total := TableItem{Files: 0, LOC: 0}

	for _, value := range mapping.Extensions {
        languages[value] = &TableItem{Files: 0, LOC: 0}
	}
	for _, value := range mapping.Components {
        components[value] = &TableItem{Files: 0, LOC: 0}
	}

	statistics := &StatisticsResponse{
		Languages:  languages,
		Components: components,
		Total:      total,
	}

	tree.Chdir(path) //nolint:errcheck
	goAroundCalculating(list, statistics, nil)
	tree.Chdir("..") //nolint:errcheck

	cleanStatistics(statistics.Languages)
	cleanStatistics(statistics.Components)

	return statistics, nil
}

func cleanStatistics(items Items) {
   	for title, item := range items {
	    if item.Files == 0 {
			delete(items, title)
		}
	}
}

func goAroundCalculating(
	list               tree.Nodes,
	existingStatistics *StatisticsResponse,
	component          *string,
) {
	for _, node := range list {
		if node.IsDir {
            newComponent, exists := mapping.Components[filepath.Base(node.Name)]

            if exists {
                component = &newComponent
            }

            list, _ = tree.List(node.Name)

            tree.Chdir(node.Name) //nolint:errcheck
			goAroundCalculating(list, existingStatistics, component)
			tree.Chdir("..") //nolint:errcheck

		} else {
		    language, exists := mapping.Extensions[filepath.Ext(node.Name)]
            if exists {
                LOC := uint64(1)
                tree.ReadNodeLineByLine(node.Name, proceedLine, &LOC)

                existingStatistics.Total.Append(LOC, 1)
                existingStatistics.Languages[language].Append(LOC, 1)

                newComponent, exists := mapping.Components[filepath.Base(node.Name)]

                if exists {
                    component = &newComponent
                }
                if component != nil {
                    existingStatistics.Components[*component].Append(LOC, 1)
                }
            }
		}
	}
}

func proceedLine(text string, counter *uint64) {
	*counter++
}
