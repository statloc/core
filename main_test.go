package statloc_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	core "github.com/statloc/core"
	"github.com/statloc/core/internal/mapping"
)

type MainSuite struct {
    suite.Suite
    results core.Statistics
}

func (s *MainSuite) SetupSuite() {
    rawResults, _ := os.ReadFile(filepath.Join("testdata", "results.json"))
    s.results = mapping.LoadJSON[core.Statistics](string(rawResults))
}

func (s *MainSuite) TestGetStatistics() {
    response, err := core.GetStatistics("testdata")

    assert.Nil(s.T(), err)
    assert.NotPanics(s.T(), func() {core.GetStatistics("testdata")}) //nolint:errcheck

    for title, item := range s.results.Components {
        assert.Equal(s.T(), item.LOC, response.Components[title].LOC)
        assert.Equal(s.T(), item.Files, response.Components[title].Files)
    }
    for title, item := range s.results.Languages{
        assert.Equal(s.T(), item.LOC, response.Languages[title].LOC)
        assert.Equal(s.T(), item.Files, response.Languages[title].Files)
    }
    assert.Equal(s.T(), s.results.Total.LOC, response.Total.LOC)
    assert.Equal(s.T(), s.results.Total.Files, response.Total.Files)
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}
