package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/jhonipereira/go-micro-product/data"
)

func (app *Config) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.Repo.GetAll()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProductByName(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	products, err := app.Repo.GetByName(requestPayload.Name)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(products)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetProductByID(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Id int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	product, err := app.Repo.GetOne(requestPayload.Id)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(product)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "",
		Data:    data,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Photos      string `json:"photos"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if requestPayload.Id == 0 || requestPayload.Name == "" || requestPayload.Description == "" || requestPayload.Photos == "" {
		app.errorJSON(w, errors.New("invalid request"), http.StatusBadRequest)
		return
	}
	productToUpdate := data.Product{
		ID:          requestPayload.Id,
		Name:        requestPayload.Name,
		Photos:      requestPayload.Photos,
		Description: requestPayload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = app.Repo.Update(productToUpdate)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "updated",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.Repo.DeleteByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			app.errorJSON(w, errors.New("product not found"), http.StatusNotFound)
			return
		}
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "deleted",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Photos      string `json:"photos"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if requestPayload.Name == "" || requestPayload.Description == "" || requestPayload.Photos == "" {
		app.errorJSON(w, errors.New("invalid request"), http.StatusBadRequest)
		return
	}

	productToInsert := data.Product{
		Name:        requestPayload.Name,
		Photos:      requestPayload.Photos,
		Description: requestPayload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	product, err := app.Repo.Insert(productToInsert)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "created",
		Data:    product,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}
