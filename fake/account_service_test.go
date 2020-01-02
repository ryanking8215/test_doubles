package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AccountService_PasswordHash(t *testing.T) {
	repo := NewFakeAccountRepository()
	svc := NewAccountService(repo)
	user := &User{Email: "john@bmail.com"}
	v := svc.GetPasswordHash(user)
	assert.Equal(t, "111", v)
}
