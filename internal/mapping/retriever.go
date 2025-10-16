package mapping

import (
	"encoding/json"
	"fmt"
)

var (
    Languages  LanguagesMapping
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

func Load(rawComponents string, rawExtensions string) (ComponentsMapping, LanguagesMapping) {
    Components = LoadJSON[ComponentsMapping](rawComponents)
    Languages = LoadJSON[LanguagesMapping](rawExtensions)
    return Components, Languages
}
