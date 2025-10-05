package mapping

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
    Extensions ExtensionsMapping
    Components ComponentsMapping
)

func LoadJSON[mapType any](fileName string) mapType {
    content, err := os.ReadFile(fileName)
    if err != nil {
        panic(InvalidJSON{Message: fmt.Sprintf("Error opening JSON at %s", fileName)})
    }

    var mapping mapType
    err = json.Unmarshal([]byte(string(content)), &mapping)
    if err != nil {
        panic(&InvalidJSON{Message: fmt.Sprintf("Error parsing JSON mapping at %s", fileName)})
    }

    return mapping
}

func Load(componentsPath string, extensionsPath string) (ComponentsMapping, ExtensionsMapping) {
    Components = LoadJSON[ComponentsMapping](componentsPath)
    Extensions = LoadJSON[ExtensionsMapping](extensionsPath)
    return Components, Extensions
}
