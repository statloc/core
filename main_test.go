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

type ServiceSuite struct {
    suite.Suite
    results map[string]map[string]uint64
}

func (s *ServiceSuite) SetupSuite() {
    rawResults, _ := os.ReadFile(filepath.Join("testdata", "results.json"))
    s.results = mapping.LoadJSON[map[string]map[string]uint64](string(rawResults))
}

func (s *ServiceSuite) TestGetStatistics() {
    response, err := core.GetStatistics("testdata")

    assert.Nil(s.T(), err)

    for title, item := range response.Languages {
        assert.Equal(s.T(), s.results[title]["LOC"], item.LOC)
        assert.Equal(s.T(), s.results[title]["Files"], item.Files)
    }

    for title, item := range response.Components {
        assert.Equal(s.T(), s.results[title]["LOC"], item.LOC)
        assert.Equal(s.T(), s.results[title]["Files"], item.Files)
    }

    assert.Equal(s.T(), s.results["Total"]["LOC"], response.Total.LOC)
    assert.Equal(s.T(), s.results["Total"]["Files"], response.Total.Files)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
