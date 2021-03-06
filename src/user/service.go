package user

import (
	"../utils"
	"context"
)

type Service interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id uint) (*User, error)

	Store(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, id uint, user User) (*User, error)
	ChangePassword(ctx context.Context, id uint, email, password string) error
}
type service struct {
	repo URepository
}

func NewUserService(r URepository) *service {
	return &service{repo: r}
}

func (s service) FindAll(ctx context.Context) (u []*User, err error) {

	return s.repo.FindAll(ctx)
}

func (s service) FindByID(ctx context.Context, id uint) (u *User, err error) {

	return s.repo.FindByID(ctx, id)
}

func (s service) Store(ctx context.Context, user User) (u *User, err error) {

	return s.repo.Store(ctx, user)
}

func (s service) Update(ctx context.Context, id uint, user User) (u *User, err error) {

	return s.repo.Update(ctx, id, user)
}

func (s service) ChangePassword(ctx context.Context, id uint, email, password string) (err error) {

	pw, err := utils.EncryptPassword(password);
	if err != nil {
		return err
	}
	return s.repo.ChangePassword(ctx, id, email, pw)
}
