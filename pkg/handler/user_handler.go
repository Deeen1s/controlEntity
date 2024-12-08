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

type UserHandler struct {
	serviceUser service.UserService //бизнес-логика для работы с User
}

func NewUserHandler(serviceUser service.UserService) *UserHandler {
	return &UserHandler{serviceUser: serviceUser} //возвращаем объект хендлера с подключённым сервисом
}

// POST /api/v1/users - Создание пользователя
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User                                          //входные данные для создания пользователя
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { //парсим JSON из тела запроса
		http.Error(w, "Некорректный ввод", http.StatusBadRequest) //если ошибка, возвращаем 400 Bad Request
		return
	}

	id, err := h.serviceUser.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError) //ошибка 500
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()       //создание временных меток
	w.WriteHeader(http.StatusCreated) //устанавливаем код ответа 201 Created
	//json.NewEncoder(w).Encode(map[string]int{"id": id}) //возвращаем ID созданного пользователя в JSON
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})

}

// GET /api/v1/users/{id} - Получение пользователя по ID
func (h *UserHandler) GetUserID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //получаем ID пользователя из параметра запроса
	id, err := strconv.Atoi(idStr)   //преобразуем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest) //если ID некорректен, возвращаем 400
		return
	}
	user, err := h.serviceUser.GetUserID(r.Context(), id) //вызываем бизнес-логику для получения пользователя
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //если пользователь не найден, возвращаем 404
		return
	}
	json.NewEncoder(w).Encode(user) //возвращаем данные пользователя в JSON
}

// PUT /api/v1/users/{id} - Обновление пользователя
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/users/") //извлекаем ID из пути запроса
	id, err := strconv.Atoi(idStr)                            //конвертируем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID польщователя", http.StatusBadRequest) //ошибка 400 Bad Request
		return
	}

	var user models.User                                          //создаём объект user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { //декодируем JSON из тела запроса в user
		http.Error(w, "Некорретный ввод", http.StatusBadRequest) //если ошибка, отправляем 400 Bad Request
		return
	}

	user.ID = id                //устанавливаем ID пользователя
	user.UpdatedAt = time.Now() //устанавливаем время обновления

	if err := h.serviceUser.UpdateUser(r.Context(), id, user); err != nil {
		http.Error(w, "Ошибка в изменении пользователя", http.StatusInternalServerError) //ошибка 500 Internal Server Error
		return
	}

	w.WriteHeader(http.StatusOK)                    //устанавливаем статус 200 OK
	w.Write([]byte("Пользователь успешно изменен")) //отправляем подтверждение об успешном обновлении
}

// DELETE /api/v1/users/{id} - Удаление пользователя
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //получаем ID пользователя из  запроса
	_, err := strconv.Atoi(idStr)    //преобразуем строку в число
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest) //если ID некорректен, возвращаем 400
		return
	}

}

// GET /api/v1/users - Получение списка пользователей
func (h *UserHandler) GetListUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.serviceUser.GetListUser(r.Context()) //вызов бизнес-логики для получения списка пользователей
	if err != nil {
		http.Error(w, "Ошибка получения списка пользователей", http.StatusInternalServerError) //если ошибка, возвращаем 500
		return
	}
	json.NewEncoder(w).Encode(users) //возвращаем список пользователей в JSON

}
