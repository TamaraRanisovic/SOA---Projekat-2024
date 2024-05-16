package handler

import (
	"context"

	"blogservice/model"
	"blogservice/proto/blogs"
	"blogservice/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogHandler struct {
	BlogService *service.BlogService
	blogs.UnimplementedBlogServiceServer
}

func NewBlogHandler(s *service.BlogService) *BlogHandler {
	return &BlogHandler{BlogService: s}
}

func (h BlogHandler) GetBlogById(ctx context.Context, req *blogs.GetBlogByIdRequest) (*blogs.GetBlogByIdResponse, error) {
	id := req.Id

	blog, err := h.BlogService.FindBlog(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Blog with given id not found")
	}

	response := &blogs.GetBlogByIdResponse{
		Blog: &blogs.Blog{
			Id:          blog.ID.String(),
			Title:       blog.Title,
			Description: blog.Description,
			DateCreated: blog.DateCreated.Format("2006-01-02"),
			Status:      model.StatusToString(blog.Status),
			Pictures:    convertPicturesToProto(blog.Pictures),
		},
	}

	return response, nil
}

func (h BlogHandler) GetAllBlogs(ctx context.Context, s *emptypb.Empty) (*blogs.GetAllBlogsResponse, error) {
	blogsList, err := h.BlogService.FindAllBlogs()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to fetch blogs")
	}
	var response blogs.GetAllBlogsResponse
	for _, blog := range *blogsList {
		response.Blogs = append(response.Blogs, &blogs.Blog{
			Id:          blog.ID.String(),
			Title:       blog.Title,
			Description: blog.Description,
			DateCreated: blog.DateCreated.Format("2006-01-02"),
			Status:      model.StatusToString(blog.Status),
			Pictures:    convertPicturesToProto(blog.Pictures),
		})
	}
	return &response, nil

}

/*
func (h *BlogHandler) CreateBlog(ctx context.Context, req *blogs.CreateBlogRequest) (*blogs.CreateBlogResponse, error) {
	formData := FormData{
		Title:       req.Title,
		Description: req.Description,
		DateCreated: req.DateCreated,
		Status:      req.Status,
		Pictures:    nil, // Note: handle pictures properly if needed
	}

	// Validate form data
	err := validateFormData(formData)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid form data: %v", err)
	}

	// Create blog object
	blog, err := createBlog(formData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to create blog: %v", err)
	}

	// Create blog using service
	err = h.BlogService.Create(blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to save blog: %v", err)
	}

	response := &blogs.CreateBlogResponse{
		Blog: &blogs.Blog{
			Id:          blog.ID,
			Title:       blog.Title,
			Description: blog.Description,
			DateCreated: blog.DateCreated.Format("2006-01-02"),
			Status:      blog.Status.String(),
			Pictures:    convertPicturesToProto(blog.Pictures),
		},
	}

	return response, nil
}*/

func convertPicturesToProto(pictures []model.Picture) []*blogs.Picture {
	var pbPictures []*blogs.Picture
	for _, pic := range pictures {
		pbPictures = append(pbPictures, &blogs.Picture{Url: pic.URL})
	}
	return pbPictures
}

/*
func convertBlogsToProto(blogs []model.Blog) []*blogs.Blog {
	var pbBlogs []*blogs.Blog
	for _, blog := range blogs {
		pbBlogs = append(pbBlogs, &blogs.Blog{
			Id:          blog.ID,
			Title:       blog.Title,
			Description: blog.Description,
			DateCreated: blog.DateCreated.Format("2006-01-02"),
			Status:      blog.Status.String(),
			Pictures:    convertPicturesToProto(blog.Pictures),
		})
	}
	return pbBlogs
}
*/
