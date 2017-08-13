package user

import (
	//"github.com/bboortz/go-utils/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
}

func TestGetCurrentUser(t *testing.T) {
	a := assert.New(t)
	//log := logger.NewLogger().Build()
	user, err := GetCurrentUser()
	a.Nil(err)
	a.NotNil(user)
	a.NotEmpty(user.Username)
	a.NotEmpty(user.Groupname)
	a.NotEmpty(user.HomeDir)
}
