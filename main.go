package statloc

import (
	_ "embed"
	"errors"
	"path/filepath"

	"github.com/statloc/core/internal/mapping"
	"github.com/statloc/core/internal/tree"
)

var (
    //go:embed "assets/languages.json"
    rawLanguagesMapping  string

    //go:embed "assets/components.json"
    rawComponentsMapping string
)

func GetStatistics(path string) (*Statistics, error) {
    mapping.Load(rawComponentsMapping, rawLanguagesMapping)

    list, err := tree.List(path)

	var treePathError *tree.PathError
	if errors.As(err, &treePathError) {
		return nil, &PathError{Path: path}
	}

	statistics := &Statistics{
		Languages:  initItems(mapping.Languages),
		Components: initItems(mapping.Components),
		Total:      TableItem{Files: 0, LOC: 0},
	}

	componentsSet := &componentSet{
	    Elements: make(map[string]struct{}),
		Tail:     nil,
	}

	tree.Chdir(path) //nolint:errcheck
	goAroundCalculating(list, statistics, nil, componentsSet)
	tree.Chdir("..") //nolint:errcheck

	cleanStatistics(statistics.Languages)
	cleanStatistics(statistics.Components)

	return statistics, nil
}

func goAroundCalculating(
	list          tree.Nodes,
	statistics    *Statistics,
	tailComponent *component,
	componentsSet *componentSet,
) {
	for _, node := range list {
		if node.IsDir {
            newComponentTitle, exists := mapping.Components[filepath.Base(node.Name)]

            exists = exists && !componentsSet.In(newComponentTitle)
            if exists {
                tailComponent = componentsSet.Add(newComponentTitle)
            }

            list, _ = tree.List(node.Name)

            tree.Chdir(node.Name) //nolint:errcheck
			goAroundCalculating(list, statistics, tailComponent, componentsSet)
			tree.Chdir("..") //nolint:errcheck

			if exists {
				componentsSet.Pop()
			}

		} else {
		    language, exists := mapping.Languages[filepath.Ext(node.Name)]
            if exists {
                LOC := uint64(1)
                tree.ReadNodeLineByLine(node.Name, proceedLine, &LOC)

                statistics.Total.Append(LOC, 1)
                statistics.Languages[language].Append(LOC, 1)

                newComponentTitle, exists := mapping.Components[filepath.Base(node.Name)]

                if exists && !componentsSet.In(newComponentTitle) {
                    statistics.Components[newComponentTitle].Append(LOC, 1)
                }

                for componentTitle := range componentsSet.Elements {
                    statistics.Components[componentTitle].Append(LOC, 1)
                }
            }
		}
	}
}

func initItems(mapping_ map[string]string) Items {
   	items := make(Items)
    for _, value := range mapping_ {
        items[value] = &TableItem{Files: 0, LOC: 0}
	}
	return items
}

func cleanStatistics(items Items) {
   	for title, item := range items {
	    if item.Files == 0 {
			delete(items, title)
		}
	}
}

func proceedLine(line string, counter *uint64) {
	*counter++
}
