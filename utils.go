package statloc

import (
	"path/filepath"

	"github.com/statloc/core/internal/retrievers/mapping"
	"github.com/statloc/core/internal/retrievers/tree"
)

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
