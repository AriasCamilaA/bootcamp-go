package products

import (
	productsService "ejercicioScaffolding/internal/application/service"
	productsModel "ejercicioScaffolding/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service productsService.ProductService
}

func NewProductHandler(service productsService.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid product ID", Error: err.Error()})
			return
		}

		product, err := h.service.GetProductByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Product not found", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Product found", Data: product})
	}
}

func (h *ProductHandler) GetAllProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.service.GetAllProducts()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Products retrieved successfully", Data: products})
	}
}

func (h *ProductHandler) SearchProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt := r.URL.Query().Get("priceGt")
		priceGtInt, err := strconv.Atoi(priceGt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid priceGt parameter", Error: err.Error()})
			return
		}

		products, err := h.service.SearchProducts(priceGtInt)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		if len(products) < 1 {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "No products found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Products found", Data: products})
	}
}

func (h *ProductHandler) AddProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct productsModel.Product

		w.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid request payload", Error: err.Error()})
			return
		}

		err = h.service.AddProduct(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to add product", Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Product added successfully", Data: newProduct})
	}
}

func (h *ProductHandler) UpdateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid product ID", Error: err.Error()})
			return
		}

		var updatedProduct productsModel.Product
		err = json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid request payload", Error: err.Error()})
			return
		}

		err = h.service.UpdateProduct(id, updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to update product", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Product updated successfully", Data: updatedProduct})
	}
}

func (h *ProductHandler) ReplaceProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid product ID", Error: err.Error()})
			return
		}

		var newProduct productsModel.Product
		err = json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid request payload", Error: err.Error()})
			return
		}

		err = h.service.ReplaceProduct(id, newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to replace product", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Product replaced successfully", Data: newProduct})
	}
}

func (h *ProductHandler) DeleteProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Invalid product ID", Error: err.Error()})
			return
		}

		err = h.service.DeleteProduct(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(productsModel.Response{Message: "Failed to delete product", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productsModel.Response{Message: "Product deleted successfully"})
	}
}
