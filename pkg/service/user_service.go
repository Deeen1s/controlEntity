package service

import (
	"context"
	"controlEntity/pkg/model"
	"errors"
	"time"
)

type UserServiceDB struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserServiceDB {
	return &UserServiceDB{
		repo: repo,
	}
}

func (s *UserServiceDB) CreateUser(ctx context.Context, user models.User) (models.User, error) {

	if user.Name == "" || user.Email == "" {
		return models.User{}, errors.New("Нужно указать имя и email")
	}
	user.CreatedAt = time.Now() //время создания
	user.UpdatedAt = time.Now() //время обновления
	return s.repo.CreateUser(ctx, user)
}

func (s *UserServiceDB) GetUserID(ctx context.Context, id int) (models.User, error) {
	return s.repo.GetUserID(ctx, id)
}

func (s *UserServiceDB) UpdateUser(ctx context.Context, id int, user models.User) error {

	if user.Name == "" || user.Email == "" {
		return errors.New("Некорректно введены данные")
	}
	user.UpdatedAt = time.Now() //время обновления
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserServiceDB) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("Некорретный ID юзера")
	}
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserServiceDB) GetListUser(ctx context.Context) ([]models.User, error) {
	return s.repo.GetListUser(ctx)
}
