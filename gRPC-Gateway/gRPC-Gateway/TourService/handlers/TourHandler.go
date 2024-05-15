package handler

import (
	"context"
	"log"
	"tourservice/proto/tours"
	"tourservice/repo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TourHandler struct {
	logger *log.Logger
	repo   *repo.TourRepository
	tours.UnimplementedTourServiceServer
}

// Injecting the logger makes this code much more testable.
func NewTourHandler(l *log.Logger, r *repo.TourRepository) *TourHandler {
	return &TourHandler{logger: l, repo: r}
}

func (h TourHandler) GetAllTours(ctx context.Context, s *emptypb.Empty) (*tours.GetAllToursResponse, error) {
	toursList, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response tours.GetAllToursResponse
	for _, tour := range toursList {
		response.Tours = append(response.Tours, &tours.Tour{
			Name:        tour.Name,
			Description: tour.Description,
			Length:      tour.Length,
			Tags:        tour.Tags,
			Difficulty:  int32(tour.Difficulty),
			Price:       tour.Price,
			GuideId:     tour.Guide_ID,
		})
	}
	return &response, nil
}

func (h TourHandler) GetTourById(ctx context.Context, req *tours.GetTourByIdRequest) (*tours.GetTourByIdResponse, error) {
	id := req.Id

	tour, err := h.repo.GetById(id)
	if err != nil {
		h.logger.Printf("Database exception: %v", err)
		return nil, status.Errorf(codes.Internal, "Database exception: %v", err)
	}

	if tour == nil {
		h.logger.Printf("Tour with id: '%s' not found", id)
		return nil, status.Errorf(codes.NotFound, "Tour with given id not found")
	}

	// Convert the tour to GetTourByIdResponse
	response := &tours.GetTourByIdResponse{
		Tour: &tours.Tour{
			Name:        tour.Name,
			Description: tour.Description,
			Length:      tour.Length,
			Tags:        tour.Tags,
			Difficulty:  int32(tour.Difficulty),
			Price:       tour.Price,
			GuideId:     tour.Guide_ID,
		},
	}
	return response, nil
}
