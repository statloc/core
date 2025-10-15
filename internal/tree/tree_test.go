package tree_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/statloc/core/internal/tree"
)

type (
    TreeSuite struct {
        suite.Suite
        dir             string
        nonExistingPath string
        file            string
        fileText        string
        hook            func (
            text    string,
            counter *uint64,
        )
    }
)

func (s *TreeSuite) SetupSuite() {
	s.dir = filepath.Join("..", "..", "testdata")
	s.nonExistingPath = "non_existing_path"
	s.file = ""
	s.hook = func(text string, counter *uint64) {
		*counter++
	}

	file, _ := os.ReadFile(s.file)
	s.fileText = string(file)
}

func (s *TreeSuite) TestList() {
    response, err := tree.List(s.dir)
    assert.Nil(s.T(), err)
    assert.IsType(s.T(), response, tree.Nodes{})
    assert.Len(s.T(), response, 4)

    _, err = tree.List(s.nonExistingPath)
    assert.NotNil(s.T(), err)
}

func (s *TreeSuite) TestChdir() {
    err := tree.Chdir("non_existing_dir")
    assert.NotNil(s.T(), err)

    err = tree.Chdir(filepath.Join("..", "..", "testdata"))
    assert.Nil(s.T(), err)

    err = tree.Chdir(filepath.Join("..", "internal", "tree"))
    assert.Nil(s.T(), err)
}

func (s *TreeSuite) TestReadNodeLineByLine() {
    counter := new(uint64)

    tree.ReadNodeLineByLine(filepath.Join(s.dir, "main.go"), s.hook, counter)

    assert.Equal(s.T(), uint64(7), *counter)
}

func TestTreeSuite(t *testing.T) {
	suite.Run(t, new(TreeSuite))
}
