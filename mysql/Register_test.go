package mysql

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Azunyan1111/multilogin/structs"
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

func TestSelectUserByUuid(t *testing.T) {
	user, err := SelectUserByUuid("uuid2")
	if err != nil{
		panic(err)
	}
	assert.Equal(t,"hoge",user.UserName)
}

func TestInsertUser(t *testing.T) {
	var user structs.User
	user.UserName = "涼風青葉"
	user.Email = "aoba@eaglejump.co.jp"
	uid,err := InsertUser(user)
	if err != nil{
		panic(err)
	}
	selectUser,err := SelectUserByUuid(uid)
	if err != nil{
		panic(err)
	}
	assert.Equal(t,uid,selectUser.Uid)
	if err := DeleteUserByUid(uid); err != nil{
		panic(err)
	}
}