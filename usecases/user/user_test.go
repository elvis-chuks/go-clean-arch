package user

import (
	"fmt"
	"github.com/elvis-chuks/go-clean-arch/domain"
	"github.com/elvis-chuks/go-clean-arch/pkg/env"
	"github.com/elvis-chuks/go-clean-arch/pkg/logger"
	"github.com/elvis-chuks/go-clean-arch/repository/mongodb"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initEnv() {
	env.LoadConfig("../../")
}

func TestUserUseCase(t *testing.T) {
	l, err := logger.Init()

	if err != nil {
		t.Error(err)
	}

	repo := mongodb.New(l)

	userUseCase_ := New(repo.UserRepo)

	assert.NotNil(t, userUseCase_)

}

func TestUserUseCase_Echo(t *testing.T) {
	l, err := logger.Init()

	if err != nil {
		t.Error(err)
	}

	initEnv()

	fmt.Println(viper.Get("DB_NAME"))

	repo := mongodb.New(l)

	userUseCase_ := New(repo.UserRepo)

	data := userUseCase_.Echo(domain.User{
		Firstname: "Elvis",
		Lastname:  "Chuks",
	})

	fmt.Println(data)
}
