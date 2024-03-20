package handler

import (
	"database-example/dto"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountService *service.AccountService
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

func (handler *AccountHandler) GetByUsernameAndPassword(writer http.ResponseWriter, req *http.Request) {
	// Decode the request body to get credentials
	var creds dto.Credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service to find the account by username and password
	account, err := handler.AccountService.FindAccountByUsernameAndPassword(creds.Username, creds.Password)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Set response content type and status code
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Encode the account into JSON and send the response
	json.NewEncoder(writer).Encode(account)
}

// Function for getting all Accounts
/*func (handler *AccountHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	accounts, err := handler.AccountService.FindAllAccounts()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	for i := range accounts {
		user, err := handler.AccountService.FindUserByID(accounts[i].UserID)
		if err != nil {
			// Handle error if necessary
			continue
		}
		accounts[i].User = user
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(accounts)
}*/

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
