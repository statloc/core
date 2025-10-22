package tree_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	t "github.com/statloc/core/internal/tree"
)

type (
    TreeSuite struct {
        suite.Suite
        tree            t.Tree
        dir             string
        nonExistingPath string
        hook            func (
            text    string,
            counter *uint64,
        )
    }
)

func (s *TreeSuite) SetupSuite() {
	s.dir = filepath.Join("..", "..", "testdata")
	s.nonExistingPath = "non_existing_path"
	s.hook = func(text string, counter *uint64) {
		*counter++
	}
	workdir, _ := os.Getwd()
	s.tree = t.Tree{WorkDir: workdir}
}

func (s *TreeSuite) TestList() {
    response, err := s.tree.List(s.dir)
    assert.Nil(s.T(), err)
    assert.IsType(s.T(), response, t.Nodes{})
    assert.Len(s.T(), response, 5)

    _, err = s.tree.List(s.nonExistingPath)
    assert.NotNil(s.T(), err)
}

func (s *TreeSuite) TestChdir() {
    err := s.tree.Chdir("non_existing_dir")
    assert.NotNil(s.T(), err)

    err = s.tree.Chdir(filepath.Join("..", "..", "testdata"))
    assert.Nil(s.T(), err)

    err = s.tree.Chdir(filepath.Join("..", "internal", "tree"))
    assert.Nil(s.T(), err)
}

func (s *TreeSuite) TestReadNodeLineByLine() {
    counter := new(uint64)

    s.tree.ReadNodeLineByLine(filepath.Join(s.dir, "main.go"), s.hook, counter)

    assert.Equal(s.T(), uint64(7), *counter)
}

func TestTreeSuite(t *testing.T) {
	suite.Run(t, new(TreeSuite))
}
