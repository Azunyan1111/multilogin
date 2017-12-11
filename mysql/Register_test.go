package mysql

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDataBaseInit(t *testing.T) {
	if assert.NoError(t, DataBaseInit()) {
		ra, err := ConnectionTest()
		if err != nil{
			panic(err)
		}
		assert.Equal(t,int64(0),ra)
	}
}