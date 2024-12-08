package service

import (
	"context"
	models "controlEntity/pkg/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error) //создание юзера
	GetUserID(ctx context.Context, id int) (models.User, error)            //чтение юзера по айди
	UpdateUser(ctx context.Context, id int, user models.User) error        //обновление данных
	DeleteUser(ctx context.Context, id int) error                          //удаление юзера
	GetListUser(ctx context.Context) ([]models.User, error)                //получение списка
}

type ProductService interface {
	CreateProduct(ctx context.Context, product models.Product) (models.Product, error) //создание продукта
	GetProductID(ctx context.Context, id int) (models.Product, error)                  //чтение продукта по id
	UpdateProduct(ctx context.Context, id int, product models.Product) error           //обновление данных
	DeleteProduct(ctx context.Context, id int) error                                   //удаление данных
	GetListProduct(ctx context.Context) ([]models.Product, error)                      //получение списка
	GetProductUserID(ctx context.Context, userID int) ([]models.Product, error)        //вывод продуктов у юзера
}

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error) //создание юзера
	GetUserID(ctx context.Context, id int) (models.User, error)            //чтение юзера по айди
	UpdateUser(ctx context.Context, id int, user models.User) error        //обновление данных
	DeleteUser(ctx context.Context, id int) error                          //удаление юзера
	GetListUser(ctx context.Context) ([]models.User, error)                //получение списка
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product) (models.Product, error) //создание продукта
	GetProductID(ctx context.Context, id int) (models.Product, error)                  //чтение продукта по id
	UpdateProduct(ctx context.Context, id int, product models.Product) error           //обновление данных
	DeleteProduct(ctx context.Context, id int) error                                   //удаление данных
	GetListProduct(ctx context.Context) ([]models.Product, error)                      //получение списка
	GetProductUserID(ctx context.Context, userID int) ([]models.Product, error)        //вывод продуктов у юзера
}

type Repository struct {
	UserRepository
	ProductRepository
}
