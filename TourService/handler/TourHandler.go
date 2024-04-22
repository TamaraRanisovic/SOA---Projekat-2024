package handlers

import (
	"Rest/model"
	"Rest/repo"

	"context"
	"log"
	"net/http"

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
*/
func (p *TourHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		p.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
