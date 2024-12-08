package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string // Хост базы данных (например, "localhost")
	Port     string // Порт базы данных (например, "5432")
	Username string // Имя пользователя базы данных
	Password string // Пароль от базы данных
	DBName   string // Имя базы данных
	SSLMode  string // Режим SSL (например, "disable" для локальной разработки)
}

// NewPostgresDB создает подключение к базе данных PostgreSQL
func NewPostgresDB(cfg Config) (*sql.DB, error) {
	// Формируем строку подключения на основе данных из Config
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

/*
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error) //создание юзера
	GetUserID(ctx context.Context, id int) (*models.User, error)   //чтение юзера по айди
	UpdateUser(ctx context.Context, user models.User) error        //обновление данных
	DeleteUser(ctx context.Context, id int) error                  //удаление юзера
	GetListUser(ctx context.Context) ([]models.User, error)        //получение списка
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product models.Product) (int, error)     //создание продукта
	GetProductID(ctx context.Context, id int) (models.Product, error)           //чтение продукта по id
	UpdateProduct(ctx context.Context, product models.Product) error            //обновление данных
	DeleteProduct(ctx context.Context, id int) error                            //удаление данных
	GetListProduct(ctx context.Context) ([]models.Product, error)               //получение списка
	GetProductUserID(ctx context.Context, userID int) ([]models.Product, error) //вывод продуктов у юзера
}

type Repository struct {
	UserRepository
	ProductRepository
}*/
