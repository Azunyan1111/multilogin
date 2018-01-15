package mysql

import (
	"github.com/Azunyan1111/multilogin/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataBaseInit(t *testing.T) {
	if assert.NoError(t, DataBaseInit()) {
		ra, err := ConnectionTest()
		assert.NoError(t, err)
		assert.Equal(t, int64(0), ra)
	}
}

func TestSelectUserByUuid(t *testing.T) {
	DataBaseInit()
	user, err := SelectUserByUuid("uuid2")
	assert.NoError(t, err)
	assert.Equal(t, "hoge", user.UserName)
}

func TestInsertUser(t *testing.T) {
	DataBaseInit()
	var user structs.Usered
	user.UserName = "涼風青葉"
	user.Email = "aoba@eaglejump.co.jp"
	uid, err := InsertUser(user)
	assert.NoError(t, err)
	selectUser, err := SelectUserByUuid(uid)
	assert.NoError(t, err)
	assert.Equal(t, uid, selectUser.Uid)
	err = DeleteUserByUid(uid)
	assert.NoError(t, err)
}
