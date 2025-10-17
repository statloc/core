package matching_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/statloc/core/internal/matching"
)

type MatchingSuite struct {
    suite.Suite
    mapping      map[string]string
    wrongMapping map[string]string
}

func (s *MatchingSuite) SetupSuite() {
    s.mapping = map[string]string{
        "^wrong$": "1",
        "^some_strin.?$": "2",
        "^some_[Ss].*$": "3",
    }
    s.wrongMapping = map[string]string{
        "^wrong$": "1",
        "^some_strin.?$": "2",
    }
}

func (s *MatchingSuite) TestFindMatch() {
    response, exists := matching.FindMatch[string]("some_String", s.mapping)

    assert.True(s.T(), exists)
    assert.Equal(s.T(), "3", response)

    _, exists = matching.FindMatch[string]("some_String", s.wrongMapping)

    assert.False(s.T(), exists)
}

func TestMatchingSuite(t *testing.T) {
	suite.Run(t, new(MatchingSuite))
}
