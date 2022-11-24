package mongodb

import (
	"github.com/Kamva/mgm/v2"
	"github.com/elvis-chuks/go-clean-arch/domain"
	"go.uber.org/zap"
)

type mongoUserRepository struct {
	Coll   *mgm.Collection
	Logger *zap.Logger
}

func (m mongoUserRepository) Create(user domain.User) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (m mongoUserRepository) GetUserById(id string) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (m mongoUserRepository) GetUsers() []domain.User {
	//TODO implement me
	panic("implement me")
}

func (m mongoUserRepository) Echo(user domain.User) domain.User {
	return user
}

func NewMongoUserRepository(l *zap.Logger) domain.UserRepository {
	return &mongoUserRepository{
		Logger: l,
		Coll:   mgm.Coll(&domain.User{}),
	}
}
