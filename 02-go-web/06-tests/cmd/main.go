package main

import (
	productsService "ejercicioTest/internal/application/service"
	productsRepo "ejercicioTest/internal/infraestructure/repository"
	productsHandler "ejercicioTest/internal/transport/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Controlador para manejar la ruta /ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	// Crear un nuevo enrutador
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Inicializar el repositorio, servicio y handlers
	repo := productsRepo.NewMapProductRepository("docs/products.json")
	service := productsService.NewProductService(repo)
	handler := productsHandler.NewProductHandler(service)

	// Definir rutas para ping
	r.Get("/ping", pingHandler)

	// Definir rutas para productos
	r.Route("/products", func(r chi.Router) {
		r.Get("/", handler.GetAllProductsHandler())
		r.Get("/{id}", handler.GetProductByIDHandler())
		r.Get("/search", handler.SearchProductsHandler())
		r.Post("/", handler.AddProductHandler())
		r.Put("/{id}", handler.UpdateProductHandler())
		r.Patch("/{id}", handler.ReplaceProductHandler())
		r.Delete("/{id}", handler.DeleteProductHandler())
	})

	// Iniciar el servidor
	fmt.Println("Servidor encendido en el puerto http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
