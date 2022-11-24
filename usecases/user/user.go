package user

import "github.com/elvis-chuks/go-clean-arch/domain"

type userUseCase struct {
	repo domain.UserRepository
}

func (u userUseCase) Create(user domain.User) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) GetUserById(id string) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) GetUsers() []domain.User {
	//TODO implement me
	panic("implement me")
}

/*
ideally your usecase implementation should be repository agnostic
you want to combine some of the repository functions you've written into a single usecase
for example, a login usecase would combine repositories such as getEmail, checkPassword and
domain functions like hash password etc, combining this into a single usecase keeps your
delivery package neat and simple to implement the same functionalities on grpc and http
*/

func (u userUseCase) Echo(user domain.User) domain.User {
	return u.repo.Echo(user)
}

func New(repo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
