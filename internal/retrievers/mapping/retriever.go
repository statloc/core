package mapping

import (
	"encoding/json"
	"fmt"
)

var (
    Extensions ExtensionsMapping
    Components ComponentsMapping
)

func LoadJSON[mapType any](content string) mapType {
    var mapping mapType
    err := json.Unmarshal([]byte(string(content)), &mapping)
    if err != nil {
        panic(&InvalidJSON{Message: fmt.Sprintf("Error parsing JSON mapping %s", content)})
    }
    return mapping
}

func Load(rawComponents string, rawExtensions string) (ComponentsMapping, ExtensionsMapping) {
    Components = LoadJSON[ComponentsMapping](rawComponents)
    Extensions = LoadJSON[ExtensionsMapping](rawExtensions)
    return Components, Extensions
}
