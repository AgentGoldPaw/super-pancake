package slices_test

import (
	"github.com/golden-protocol/super-pancake/slices"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type slicesSearchSuite struct {
	suite.Suite
}

func (s *slicesSearchSuite) TestFindValueSuccess() {
	values := []string{"search", "for", "a", "string"}
	valueToFind := "string"
	index := slices.IndexOfString(values, valueToFind)
	assert.Equal(s.T(), len(values)-1, index)
}

func (s *slicesSearchSuite) TestFindValueFailure() {
	values := []string{"search", "for", "a", "string"}
	valueToFind := "notfound"
	index := slices.IndexOfString(values, valueToFind)
	assert.Equal(s.T(), -1, index)
}

func TestSensorSuite(t *testing.T) {
	suite.Run(t, new(slicesSearchSuite))
}
