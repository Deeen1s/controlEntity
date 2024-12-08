package handler

import (
	models "controlEntity/pkg/model"
	"controlEntity/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProductHandler struct {
	serviceProd service.ProductService //бизнес-логика для работы с Product
}

func NewProductHandler(serviceProd service.ProductService) *ProductHandler {
	return &ProductHandler{serviceProd: serviceProd} //возвращаем объект хендлера с подключённым сервисом
}

// POST /api/v1/product - Создание пользователя

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product                                       //входные данные для создания продукта
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil { //парсим JSON из тела запроса
		http.Error(w, "Некорректный ввод", http.StatusBadRequest) //если ошибка, возвращаем 400 Bad Request
		return
	}

	id, err := h.serviceProd.CreateProduct(r.Context(), product)
	if err != nil {
		http.Error(w, "Ошибка создания продукта", http.StatusInternalServerError) //ошибка 500
		return
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now() //создание временных меток

	w.WriteHeader(http.StatusCreated) //устанавливаем код ответа 201 Created
	//json.NewEncoder(w).Encode(map[string]int{"id": id}) //возвращаем ID созданного пользователя в JSON
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// GET /api/v1/product/{id} - Получение продукта по

func (h *ProductHandler) GetProductID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //получаем ID продукта из параметра запроса
	id, err := strconv.Atoi(idStr)   //преобразуем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest) //если ID некорректен, возвращаем 400
		return
	}

	product, err := h.serviceProd.GetProductID(r.Context(), id) //вызываем бизнес-логику для получения пользователя
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //если пользователь не найден, возвращаем 404
		return
	}
	json.NewEncoder(w).Encode(product) //возвращаем данные пользователя в JSON
}

// PUT /api/v1/product/{id} - Обновление продукта
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/product/") //извлекаем ID из пути запроса
	id, err := strconv.Atoi(idStr)                              //конвертируем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID продукта", http.StatusBadRequest) //ошибка 400 Bad Request
		return
	}

	var product models.Product                                       //создаём объект user
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil { //декодируем JSON из тела запроса в user
		http.Error(w, "Некорретный ввод", http.StatusBadRequest) //если ошибка, отправляем 400 Bad Request
		return
	}

	product.ID = id                //устанавливаем ID пользователя
	product.UpdatedAt = time.Now() //устанавливаем время обновления

	if err := h.serviceProd.UpdateProduct(r.Context(), id, product); err != nil {
		http.Error(w, "Ошибка в изменении продукта", http.StatusInternalServerError) //ошибка 500 Internal Server Error
		return
	}

	w.WriteHeader(http.StatusOK)               //устанавливаем статус 200 OK
	w.Write([]byte("Продукт успешно изменен")) //отправляем подтверждение об успешном обновлении
}

// DELETE /api/v1/product/{id} - Удаление продукта
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //получаем ID пользователя из  запроса
	_, err := strconv.Atoi(idStr)    //преобразуем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID продукта", http.StatusBadRequest) //если ID некорректен, возвращаем 400
		return
	}

}

// GET /api/v1/users - Получение списка пользователей
func (h *ProductHandler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	product, err := h.serviceProd.GetListProduct(r.Context()) //вызов бизнес-логики для получения списка продуктов
	if err != nil {
		http.Error(w, "Ошибка получения списка продуктов", http.StatusInternalServerError) //если ошибка, возвращаем 500
		return
	}
	json.NewEncoder(w).Encode(product) //возвращаем список продуктов в JSON

}

func (h *ProductHandler) GetProductUserID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
	parts := strings.Split(path, "/")

	if len(parts) < 2 || parts[1] != "products" { //проверяем правильность пути
		http.Error(w, "Некорретный пользователь", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(parts[0]) //извлекаем user_id из пути
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	//вызываем бизнес-логику
	products, err := h.serviceProd.GetProductUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
