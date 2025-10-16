package mapping_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/statloc/core/internal/mapping"
)

type MappingSuite struct {
    suite.Suite
    extensions     string
    components     string
    broken         string
}

func (s *MappingSuite) SetupSuite() {
    rawExtensions, _ := os.ReadFile(filepath.Join("..", "..", "assets", "languages.json"))
    s.extensions = string(rawExtensions)
    rawComponents, _ := os.ReadFile(filepath.Join("..", "..", "assets", "components.json"))
    s.components = string(rawComponents)
    rawBroken, _ := os.ReadFile(filepath.Join("..", "..", "testdata", "broken.json.txt"))
    s.broken = string(rawBroken)
}

func (s *MappingSuite) TestLoadJSON() {
    response := mapping.LoadJSON[map[string]string](s.extensions)

    assert.IsType(s.T(), map[string]string{}, response)
    assert.Panics(s.T(), func() {mapping.LoadJSON[map[string]string](s.broken)})
}

func (s *MappingSuite) TestLoadMapping() {
    mapping.Load(
        s.components,
        s.extensions,
    )

    assert.NotNil(s.T(), mapping.Components)
    assert.NotNil(s.T(), mapping.Languages)
}

func TestMappingSuite(t *testing.T) {
	suite.Run(t, new(MappingSuite))
}
