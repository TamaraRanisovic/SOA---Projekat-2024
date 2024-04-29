package handlers

import (
	"Rest/model"
	"Rest/repo"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type TourHandler struct {
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *repo.TourRepository
}

// Injecting the logger makes this code much more testable.
func NewTourHandler(l *log.Logger, r *repo.TourRepository) *TourHandler {
	return &TourHandler{l, r}
}

func (p *TourHandler) GetAllTours(rw http.ResponseWriter, h *http.Request) {
	tours, err := p.repo.GetAll()
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if tours == nil {
		return
	}

	err = tours.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *TourHandler) GetAllToursByGuide(w http.ResponseWriter, r *http.Request) {
	var tokenBody struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&tokenBody)
	if err != nil {
		log.Println("Failed to decode tokenBody:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode tokenBody\n"))
		return
	}

	log.Println("tokenBodyJSONTRIMSPACE: ", `{"token": "`+strings.TrimSpace(tokenBody.Token)+`"}`)

	authenticateGuideURL := "http://user_management_service:8085/authenticate-guide/"
	resp, err := http.Post(authenticateGuideURL, "application/json", bytes.NewBuffer([]byte(`{"token": "`+strings.TrimSpace(tokenBody.Token)+`"}`)))
	if err != nil {
		log.Println("Failed to make POST request to User Management microservice:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to authenticate user\n"))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		getUserIDURL := "http://user_management_service:8085/get/user/token"
		resp2, err := http.Post(getUserIDURL, "application/json", bytes.NewBuffer([]byte(`{"token": "`+strings.TrimSpace(tokenBody.Token)+`"}`)))
		if err != nil {
			log.Println("Failed to make POST request to User Management microservice:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to authenticate user\n"))
			return
		}

		body, err := io.ReadAll(resp2.Body)
		if err != nil {
			log.Println("Failed to read response body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to read response body\n"))
			return
		}

		log.Println("ID GUIDEResponse body:", string(body))

		var bodyJSON struct {
			Guide_ID string `json:"id"`
		}

		err = json.Unmarshal(body, &bodyJSON)
		if err != nil {
			log.Println("Failed to decode response body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to decode response body\n"))
			return
		}

		tours, err := p.repo.GetByGuideId(bodyJSON.Guide_ID)
		if err != nil {
			p.logger.Print("Database exception: ", err)
		}

		jsonBytes, err := json.Marshal(tours)
		if err != nil {
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonBytes)

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

	} else {
		log.Println("Unauthorized: only guides can perform this action")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unauthorized: only guides can perform this action\n"))
	}

}

func (p *TourHandler) GetTourById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	tour, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if tour == nil {
		http.Error(rw, "Tour with given id not found", http.StatusNotFound)
		p.logger.Printf("Tour with id: '%s' not found", id)
		return
	}

	err = tour.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

/*
	func (p *PatientsHandler) GetPatientsByName(rw http.ResponseWriter, h *http.Request) {
		name := h.URL.Query().Get("name")

		patients, err := p.repo.GetByName(name)
		if err != nil {
			p.logger.Print("Database exception: ", err)
		}

		if patients == nil {
			return
		}

		err = patients.ToJSON(rw)
		if err != nil {
			http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
			p.logger.Fatal("Unable to convert to json :", err)
			return
		}
	}
*/

type TourFormData struct {
	Name        string
	Description string
	Length      float64
	Tags        []string
	Difficulty  int
	Price       float64
	Token       string
}

func parseTourFormData(req *http.Request) (TourFormData, error) {
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		return TourFormData{}, errors.New("failed to parse form data")
	}

	name := req.Form.Get("name")
	if name == "" {
		return TourFormData{}, errors.New("name is a required field")
	}

	length, _ := strconv.ParseFloat(req.Form.Get("length"), 64)

	price, _ := strconv.ParseFloat(req.Form.Get("price"), 64)

	difficulty, _ := strconv.Atoi(req.Form.Get("difficulty"))

	tags := strings.Split(req.Form.Get("tags"), ",")

	token := req.Form.Get("token")

	tourFormData := TourFormData{
		Name:        req.Form.Get("name"),
		Description: req.Form.Get("description"),
		Length:      length,
		Tags:        tags,
		Difficulty:  difficulty,
		Price:       price,
		Token:       token,
	}

	return tourFormData, nil
}

func (p *TourHandler) AddTourHandler(w http.ResponseWriter, r *http.Request) {

	tourFormData, err := parseTourFormData(r)

	var tokenBody struct {
		Token string `json:"token"`
	}

	if tourFormData.Token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No token provided\n"))
		return
	}

	tokenBody.Token = tourFormData.Token
	log.Println("tokenBody.Token: ", tokenBody.Token)

	/*tokenBodyJSON, err := json.Marshal(tokenBody)
	if err != nil {
		log.Println("Failed to marshal tokenBody:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal tokenBody\n"))
		return
	}*/
	log.Println("tokenBodyJSONTRIMSPACE: ", `{"token": "`+strings.TrimSpace(tourFormData.Token)+`"}`)

	authenticateGuideURL := "http://user_management_service:8085/authenticate-guide/"
	resp, err := http.Post(authenticateGuideURL, "application/json", bytes.NewBuffer([]byte(`{"token": "`+strings.TrimSpace(tourFormData.Token)+`"}`)))
	if err != nil {
		log.Println("Failed to make POST request to User Management microservice:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to authenticate user\n"))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		getUserIDURL := "http://user_management_service:8085/get/user/token" // Change this to the actual decode endpoint
		resp2, err := http.Post(getUserIDURL, "application/json", bytes.NewBuffer([]byte(`{"token": "`+strings.TrimSpace(tourFormData.Token)+`"}`)))
		if err != nil {
			log.Println("Failed to make POST request to User Management microservice:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to authenticate user\n"))
			return
		}

		body, err := io.ReadAll(resp2.Body)
		if err != nil {
			log.Println("Failed to read response body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to read response body\n"))
			return
		}

		// Log the response body
		log.Println("ID GUIDEResponse body:", string(body))

		// Decode the response body into bodyJSON struct
		var bodyJSON struct {
			Guide_ID string `json:"id"`
		}

		err = json.Unmarshal(body, &bodyJSON)
		if err != nil {
			log.Println("Failed to decode response body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to decode response body\n"))
			return
		}

		tour := model.Tour{
			Name:        tourFormData.Name,
			Description: tourFormData.Description,
			Length:      tourFormData.Length,
			Tags:        tourFormData.Tags,
			Difficulty:  tourFormData.Difficulty,
			Price:       tourFormData.Price,
			Guide_ID:    bodyJSON.Guide_ID,
		}

		err = p.repo.Insert(&tour)
		if err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		// Respond with success HTML
		w.Header().Set("Content-Type", "text/html") // Set content type before writing response
		w.WriteHeader(http.StatusOK)

		htmlContent, err := os.ReadFile("/app/static/html/success.html")
		if err != nil {
			handleError(w, fmt.Errorf("failed to read HTML file: %v", err), http.StatusInternalServerError)
			return
		}
		w.Write(htmlContent)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unauthorized: only guides can perform this action\n"))
	}
}

/*func (p *TourHandler) AddTourHandler(w http.ResponseWriter, r *http.Request) {
	tourFormData, err := parseTourFormData(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tour := model.Tour{
		Name:        tourFormData.Name,
		Description: tourFormData.Description,
		Length:      tourFormData.Length,
		Tags:        tourFormData.Tags,
		Difficulty:  tourFormData.Difficulty,
		Price:       tourFormData.Price,
	}

	//p.repo.Insert(&tour)

	err = p.repo.Insert(&tour)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"message": "Tour created successfully"})

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	htmlContent, err := os.ReadFile("html/success.html")
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(htmlContent))
}
*/

/*
	func (p *TourHandler) PostTour(rw http.ResponseWriter, h *http.Request) {
		tour := h.Context().Value(KeyProduct{}).(*model.Tour)
		p.repo.Insert(tour)
		rw.WriteHeader(http.StatusCreated)
	}

/*

	func (p *PatientsHandler) PatchPatient(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]
		patient := h.Context().Value(KeyProduct{}).(*data.Patient)

		p.repo.Update(id, patient)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) AddPhoneNumber(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]

		var phoneNumber string
		d := json.NewDecoder(h.Body)
		d.Decode(&phoneNumber)

		p.repo.AddPhoneNumber(id, phoneNumber)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) DeletePatient(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]

		p.repo.Delete(id)
		rw.WriteHeader(http.StatusNoContent)
	}

	func (p *PatientsHandler) AddAnamnesis(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]
		anamnesis := h.Context().Value(KeyProduct{}).(*data.Anamnesis)

		p.repo.AddAnamnesis(id, anamnesis)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) AddTherapy(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]
		therapy := h.Context().Value(KeyProduct{}).(*data.Therapy)

		p.repo.AddTherapy(id, therapy)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) ChangeAddress(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]
		address := h.Context().Value(KeyProduct{}).(*data.Address)

		p.repo.UpdateAddress(id, address)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) ChangePhone(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]
		index, err := strconv.Atoi(vars["index"])
		if err != nil {
			http.Error(rw, "Unable to decode index", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		var phoneNumber string
		d := json.NewDecoder(h.Body)
		d.Decode(&phoneNumber)

		p.repo.ChangePhone(id, index, phoneNumber)
		rw.WriteHeader(http.StatusOK)
	}

	func (p *PatientsHandler) Receipt(rw http.ResponseWriter, h *http.Request) {
		vars := mux.Vars(h)
		id := vars["id"]

		total, err := p.repo.Receipt(id)
		if err != nil {
			p.logger.Print("Database exception: ", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		totalJson := map[string]float64{"total": total}

		e := json.NewEncoder(rw)
		e.Encode(totalJson)
	}

	func (p *PatientsHandler) Report(rw http.ResponseWriter, h *http.Request) {
		report, err := p.repo.Report()
		if err != nil {
			p.logger.Print("Database exception: ", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		e := json.NewEncoder(rw)
		e.Encode(report)
	}
*/
func (p *TourHandler) MiddlewareTourDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tour := &model.Tour{}
		err := tour.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, tour)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

/*
// Solution: we added middlewares for Anamnesis, Therapy and Address objects

	func (p *PatientsHandler) MiddlewareAnamnesisDeserialization(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
			anamnesis := &data.Anamnesis{}
			err := anamnesis.FromJSON(h.Body)
			if err != nil {
				http.Error(rw, "Unable to decode json", http.StatusBadRequest)
				p.logger.Fatal(err)
				return
			}

			ctx := context.WithValue(h.Context(), KeyProduct{}, anamnesis)
			h = h.WithContext(ctx)

			next.ServeHTTP(rw, h)
		})
	}

	func (p *PatientsHandler) MiddlewareTherapyDeserialization(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
			therapy := &data.Therapy{}
			err := therapy.FromJSON(h.Body)
			if err != nil {
				http.Error(rw, "Unable to decode json", http.StatusBadRequest)
				p.logger.Fatal(err)
				return
			}

			ctx := context.WithValue(h.Context(), KeyProduct{}, therapy)
			h = h.WithContext(ctx)

			next.ServeHTTP(rw, h)
		})
	}

	func (p *PatientsHandler) MiddlewareAddressDeserialization(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
			address := &data.Address{}
			err := address.FromJSON(h.Body)
			if err != nil {
				http.Error(rw, "Unable to decode json", http.StatusBadRequest)
				p.logger.Fatal(err)
				return
			}

			ctx := context.WithValue(h.Context(), KeyProduct{}, address)
			h = h.WithContext(ctx)

			next.ServeHTTP(rw, h)
		})
	}

func (p *TourHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
*/
// Handle HTTP errors
func handleError(writer http.ResponseWriter, err error, status int) {
	http.Error(writer, err.Error(), status)
}
