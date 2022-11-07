package handlers

import (
	productdto "backend/dto/product"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

// Create `path_file` Global variable here ...
// var path_file = "http://localhost:5000/uploads/"

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range products {
		products[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	var product []models.Product
	product, err := h.ProductRepository.GetProductUser(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		fmt.Println(err)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range product {
		product[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataUpload := r.Context().Value("dataFile") // add this code
	filename := dataUpload.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	request := productdto.ProductRequest{
		Name:  r.FormValue("name"),
		Desc:  r.FormValue("desc"),
		Price: price,
		Qty:   qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product := models.Product{
		Name:   request.Name,
		Desc:   request.Desc,
		Price:  request.Price,
		Image:  filename,
		Qty:    request.Qty,
		UserID: userId,
	}

	// err := mysql.DB.Create(&product).Error
	product, err = h.ProductRepository.CreateProduct(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

// func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
// 	userId := int(userInfo["id"].(float64))

// 	//get image filename
// 	dataContext := r.Context().Value("dataFile")
// 	filename := dataContext.(string)

// 	price, _ := strconv.Atoi(r.FormValue("price"))
// 	qty, _ := strconv.Atoi(r.FormValue("qty"))

// 	request := productdto.ProductRequest{
// 		Name:  r.FormValue("name"),
// 		Desc:  r.FormValue("desc"),
// 		Price: price,
// 		Qty:   qty,
// 		UserID:  userId,
// 	}

// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Get all category data by id []
// 	var product []models.Product
// 	if len(productId) != 0 {
// 		product, _ = h.ProductRepository.FindProducts(productId)
// 	}

// 	product, _ := h.ProductRepository.GetProduct(id)

// 	product.Name = request.Name
// 	product.Desc = request.Desc
// 	product.Price = request.Price
// 	product.Qty = request.Qty

// 	if filename != "false" {
// 		product.Image = filename
// 	}

// 	product, err = h.ProductRepository.UpdateProduct(product)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleteProduct, err := h.ProductRepository.DeleteProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: deleteProduct}
	json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(u models.Product) models.ProductResponse {
	return models.ProductResponse{
		Name:  u.Name,
		Desc:  u.Desc,
		Price: u.Price,
		Image: u.Image,
		Qty:   u.Qty,
		User:  u.User,
	}
}
