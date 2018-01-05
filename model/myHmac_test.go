package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var uid string
var message string
var hash string

func TestGetHmac(t *testing.T) {
	uid = "26d2983e-3d5a-421c-bf6f-d4608025e555"
	message = "this is hash message"
	hash = GetHmac(uid, message)
	assert.Equal(t, "6d5fdb602338bdb0e45d48098abb3c702fa74a206cec121e1da2c6fbfcea546b", hash)
}

func TestCheckHmac(t *testing.T) {
	assert.Equal(t, true, CheckHmac(uid, message, hash))
	assert.Equal(t, false, CheckHmac(uid, message, hash[1:]))
	assert.Equal(t, false, CheckHmac(uid, message[1:], hash))
	assert.Equal(t, false, CheckHmac(uid[1:], message, hash))
}
