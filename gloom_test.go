package gloom

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddAndQuery(t *testing.T) {
	assert := assert.New(t)

	bf := New(3, 256)

	testDataToAdd := [][]byte{
		[]byte("Test"),
		[]byte("Testing"),
		[]byte("asdf"),
		[]byte("asdfasdf"),
	}
	testDataToNotAdd := [][]byte{
		[]byte("Test1"),
		[]byte("Testinf"),
		[]byte("asdf123123"),
		[]byte("asdfasdf33"),
	}
	for _, data := range testDataToAdd {
		bf.Add(data)
	}
	for i, data := range testDataToAdd {
		assert.Truef(bf.Query(data), "Querying test data #%d expected 'true' but got 'false'", i)
	}
	for i, data := range testDataToNotAdd {
		assert.Falsef(bf.Query(data), "Querying test data #%d expected 'false' but got 'true'", i)
	}
}