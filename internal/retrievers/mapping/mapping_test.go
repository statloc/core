package mapping_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/statloc/core/internal/retrievers/mapping"
)

type MappingSuite struct {
    suite.Suite
    dir            string
    extensionsFile string
    componentsFile string
}


func (s *MappingSuite) SetupSuite() {
    s.dir = "testdata"
    s.extensionsFile = filepath.Join("..", "..", "..", "assets", "extensions.json")
    s.componentsFile = filepath.Join("..", "..", "..", "assets", "components.json")
}

func (s *MappingSuite) TestLoadJSON() {
    response := mapping.LoadJSON[map[string]string](s.extensionsFile)

    assert.IsType(s.T(), map[string]string{}, response)

    assert.Panics(s.T(), func() {mapping.LoadJSON[map[string]string](filepath.Join(s.dir, "broken.json.txt"))})
}

func (s *MappingSuite) TestLoadMapping() {
    mapping.Load(
        s.componentsFile,
        s.extensionsFile,
    )

    assert.NotNil(s.T(), mapping.Components)
    assert.NotNil(s.T(), mapping.Extensions)
}

func TestMappingSuite(t *testing.T) {
	suite.Run(t, new(MappingSuite))
}
