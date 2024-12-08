package service

import (
	"context"
	"controlEntity/pkg/model"
	"errors"
	"time"
)

type ProductServiceDB struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductServiceDB {
	return &ProductServiceDB{repo: repo}
}

func (s *ProductServiceDB) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {

	if product.Name == "" {
		return models.Product{}, errors.New("Имя не заполнено")
	}
	if product.Price <= 0 {
		return models.Product{}, errors.New("Цена не может быть <= 0")
	}
	product.CreatedAt = time.Now() //время создания
	product.UpdatedAt = time.Now() //время обновления
	return s.repo.CreateProduct(ctx, product)
}

func (s *ProductServiceDB) GetProductID(ctx context.Context, id int) (models.Product, error) {
	if id <= 0 {
		return models.Product{}, errors.New("Некорректный ID")
	}
	return s.repo.GetProductID(ctx, id)
}

func (s *ProductServiceDB) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	if product.ID < 0 {
		return errors.New("Неверный ID")
	}

	if product.Name == "" {
		return errors.New("Требуется имя продукта")
	}

	if product.Price <= 0 {
		return errors.New("Цена меньше 0")
	}

	if product.Description == "" {
		return errors.New("Пустое описание")
	}

	product.UpdatedAt = time.Now() //время обновления

	return s.repo.UpdateProduct(ctx, id, product)
}

func (s *ProductServiceDB) DeleteProduct(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("Неверный ID")
	}
	return s.repo.DeleteProduct(ctx, id)
}

func (s *ProductServiceDB) GetListProduct(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetListProduct(ctx)
}

func (s *ProductServiceDB) GetProductUserID(ctx context.Context, userID int) ([]models.Product, error) {
	if userID <= 0 {
		return nil, errors.New("Юзер не найден")
	}

	return s.repo.GetProductUserID(ctx, userID)
}
