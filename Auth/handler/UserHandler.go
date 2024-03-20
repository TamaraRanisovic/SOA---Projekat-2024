package handler

/*
type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("User sa id-em %s", id)
	user, err := handler.UserService.FindUser()(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func (handler *UserHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.Create(&user)
	if err != nil {
		println("Error while creating a new user")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}*/
