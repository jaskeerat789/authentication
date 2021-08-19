package controlers

import (
	"auth/product"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

type ProductController struct {
	l hclog.Logger
}

func NewProductController() *ProductController {
	log := hclog.New(&hclog.LoggerOptions{
		Name: "Product controller",
	})
	return &ProductController{l: log}
}

func (pc *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	pc.l.Info("GET /products")
	pc.l.Debug("Get all products handler")
	payload := product.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	err := payload.ToJSON(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		pc.l.Error("Failed to encode products into json")
	}
	respondWithJSON(w, http.StatusOK)

}

func (pc *ProductController) GetProductBySlug(w http.ResponseWriter, r *http.Request) {
	pc.l.Info("GET /products/{slug}/feedback")
	pc.l.Debug("Get a product by slug handler")
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	p, err := product.FindBySlug(vars["slug"])
	if err != nil {
		pc.respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	err = p.ToJson(w)
	if err != nil {
		pc.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK)
}

func (pc *ProductController) respondWithError(w http.ResponseWriter, code int, message string) {
	pc.l.Error(message)
	respondWithJSON(w, code)
}

func respondWithJSON(w http.ResponseWriter, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}
