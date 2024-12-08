package router

import (
	"controlEntity/pkg/handler"
	"net/http"
)

func InitRoutes(rout *http.ServeMux, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) *http.ServeMux {
	mux := http.NewServeMux()
	// Маршруты для пользователей
	mux.HandleFunc("/api/v1/{users}", userHandler.CreateUser)      // POST
	mux.HandleFunc("/api/v1/users/list", userHandler.GetUserID)    // GET
	mux.HandleFunc("/api/v1/users/", userHandler.GetListUser)      // GET
	mux.HandleFunc("/api/v1/users/update", userHandler.UpdateUser) // PUT
	mux.HandleFunc("/api/v1/users/delete", userHandler.DeleteUser) // DELETE

	// Маршруты для продуктов
	mux.HandleFunc("/api/v1/products", productHandler.CreateProduct)          // POST
	mux.HandleFunc("/api/v1/products/list", productHandler.GetProductID)      // GET
	mux.HandleFunc("/api/v1/products/", productHandler.GetListProduct)        // GET
	mux.HandleFunc("/api/v1/products/update", productHandler.UpdateProduct)   // PUT
	mux.HandleFunc("/api/v1/products/delete", productHandler.DeleteProduct)   // DELETE
	mux.HandleFunc("/api/v1/users/products", productHandler.GetProductUserID) // GET

	return mux

}
