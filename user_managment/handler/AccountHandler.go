package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountService *service.AccountService
	UserService *service.UserService
}

// Function for getting Account by given id
// Printing into terminal
// Returning json object
func (handler *AccountHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Account with id %s", id)
	account, err := handler.AccountService.FindAccount(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(account)
}

// Function for getting all Accounts
func (handler *AccountHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	accounts, err := handler.AccountService.FindAllAccounts()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	/*allAccounts := *accounts
	for i := range allAccounts {
        user, err := handler.UserService.FindUser(allAccounts[i].ID) 
        if err != nil {
            // Handle error if necessary
            continue
        }
        allAccounts[i].User = user
    }

	allAccounts, err := json.MarshalIndent(accounts, "", "    ")
    if err != nil {
        // Handle error if necessary
        writer.Header().Set("Content-Type", "application/json")
        writer.WriteHeader(http.StatusInternalServerError)
        return
    }*/

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(accounts)
}

// Function for creating a new account
func (handler *AccountHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.AccountService.Create(&account)
	if err != nil {
		println("Error while creating a new account")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

//TODO: Function for blocking an account
