package tree_test

import (
	"fmt"
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

func (s *TreeSuite) TestCopy() {
    response := s.tree.Copy()

    assert.Equal(s.T(), t.Tree{WorkDir: s.tree.WorkDir}, response)
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
    wd := s.tree.WorkDir

    assert.NotNil(s.T(), err)
    assert.EqualError(s.T(), err, "\"non_existing_dir\": no such file or directory")
    assert.Equal(s.T(), wd, s.tree.WorkDir)

    err = s.tree.Chdir(filepath.Join("..", "..", "testdata", "results.json"))
    assert.NotNil(s.T(), err)
    assert.EqualError(
        s.T(),
        err,
        fmt.Sprintf("\"%s\" is not a directory", filepath.Join("..", "..", "testdata", "results.json")),
    )

    err = s.tree.Chdir(filepath.Join("..", "..", "testdata"))
    assert.Nil(s.T(), err)
    assert.Equal(s.T(), filepath.Join(wd, "..", "..", "testdata"), s.tree.WorkDir)

    err = s.tree.Chdir(filepath.Join("..", "internal", "tree"))
    assert.Equal(s.T(), wd, s.tree.WorkDir)
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
