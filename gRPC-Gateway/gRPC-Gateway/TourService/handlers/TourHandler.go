package handler

import (
	"context"
	"errors"
	"log"
	"tourservice/model"
	"tourservice/proto/tours"
	"tourservice/proto/users"

	"tourservice/repo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TourHandler struct {
	logger *log.Logger
	repo   *repo.TourRepository
	tours.UnimplementedTourServiceServer
	UserServiceClient users.UserServiceClient
}

// Injecting the logger makes this code much more testable.
func NewTourHandler(l *log.Logger, r *repo.TourRepository, userServiceClient users.UserServiceClient) *TourHandler {
	return &TourHandler{logger: l, repo: r, UserServiceClient: userServiceClient}
}

func (h TourHandler) GetAllTours(ctx context.Context, s *emptypb.Empty) (*tours.GetAllToursResponse, error) {
	toursList, err := h.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response tours.GetAllToursResponse
	for _, tour := range toursList {
		response.Tours = append(response.Tours, &tours.Tour{
			Id:          tour.ID.String(),
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
			Id:          tour.ID.String(),
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

func (h TourHandler) AddTour(ctx context.Context, req *tours.AddTourRequest) (*emptypb.Empty, error) {

	authReq := &users.TokenRequest{Token: req.Token}
	_, err := h.UserServiceClient.AuthenticateGuide(ctx, authReq)
	if err != nil {
		log.Println("Failed to authenticate guide:", err)
		return nil, errors.New("failed to authenticate guide")
	}

	getUserReq := &users.TokenRequest{Token: req.Token}
	userResp, err := h.UserServiceClient.GetUserByToken(ctx, getUserReq)
	if err != nil {
		log.Println("Failed to get guide ID:", err)
		return nil, errors.New("failed to get guide ID")
	}

	newTour := model.Tour{
		Name:        req.Name,
		Description: req.Description,
		Length:      req.Length,
		Tags:        req.Tags,
		Difficulty:  int(req.Difficulty),
		Price:       req.Price,
		Guide_ID:    userResp.Id,
	}

	err = h.repo.Insert(&newTour)
	if err != nil {
		log.Println("Failed to insert tour:", err)
		return nil, errors.New("failed to insert tour")
	}

	return &emptypb.Empty{}, nil
}