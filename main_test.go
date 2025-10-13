package statloc_test

import (
	"os"
	"path/filepath"
	"testing"

	core "github.com/statloc/core"
	"github.com/statloc/core/internal/retrievers/mapping"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

    assert.Equal(s.T(), s.results["Go"]["LOC"], response.Items["Go"].LOC)
    assert.Equal(s.T(), s.results["Go"]["Files"], response.Items["Go"].Files)
    assert.Equal(s.T(), s.results["Rust"]["LOC"], response.Items["Rust"].LOC)
    assert.Equal(s.T(), s.results["Rust"]["Files"], response.Items["Rust"].Files)
    assert.Equal(s.T(), s.results["Python"]["LOC"], response.Items["Python"].LOC)
    assert.Equal(s.T(), s.results["Python"]["Files"], response.Items["Python"].Files)
    assert.Equal(s.T(), s.results["Tests"]["LOC"], response.Items["Tests"].LOC)
    assert.Equal(s.T(), s.results["Tests"]["Files"], response.Items["Tests"].Files)
    assert.Equal(s.T(), s.results["Total"]["LOC"], response.Items["Total"].LOC)
    assert.Equal(s.T(), s.results["Total"]["Files"], response.Items["Total"].Files)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
