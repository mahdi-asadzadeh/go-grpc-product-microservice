package services

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/db"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/models"
	"github.com/mahdi-asadzadeh/go-grpc-product-microservice/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	H db.Handler
}

func (s *Server) ImagesProduct(req *pb.ImagesProductRequest, stream pb.ProductService_ImagesProductServer) error {
	return nil
}

func (s *Server) UploadImagesProduct(stream pb.ProductService_UploadImagesProductServer) error {
	return nil
}

func (s *Server) DetailProduct(ctx context.Context, req *pb.DetailProductRequest) (*pb.DetailProductResponse, error) {
	var product models.Product
	err := s.H.DB.Where("id = ?", req.GetId()).First(&product).Error
	if err != nil {
		return nil, status.Error(codes.NotFound, "Not found product.")
	}
	return &pb.DetailProductResponse{
		Product: &pb.Product{
			Id:       int64(product.ID),
			Slug:     product.Slug,
			Title:    product.Title,
			Body:     product.Body,
			Price:    float32(product.Price),
			CreateAt: product.CreatedAt.String(),
			UpdateAt: product.UpdatedAt.String(),
			DeleteAt: "",
		},
	}, nil
}

func (s *Server) ListProduct(req *pb.ListProductRequest, stream pb.ProductService_ListProductServer) error {
	var products []models.Product
	var count int64
	offSet := (req.GetPage() - 1) * req.GetPageSize()

	s.H.DB.Model(&products).Count(&count)
	err := s.H.DB.Offset(int(offSet)).Limit(int(req.GetPageSize())).Find(&products).Error
	if err != nil {
		return status.Error(codes.Internal, "Products are empty.")
	}

	for _, pro := range products {
		res := pb.ListProductResponse{Product: &pb.Product{
			Id:       int64(pro.ID),
			Slug:     pro.Slug,
			Title:    pro.Title,
			Body:     pro.Body,
			Price:    float32(pro.Price),
			CreateAt: pro.CreatedAt.String(),
			UpdateAt: pro.UpdatedAt.String(),
			DeleteAt: "",
		}}
		stream.Send(&res)
	}
	return nil
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product
	product.Slug = slug.Make(req.GetSlug())
	product.Title = req.GetTitle()
	product.Body = req.GetBody()
	product.Price = float64(req.GetPrice())

	err := s.H.DB.Create(&product).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Invalid product data.")
	}
	return &pb.CreateProductResponse{Product: &pb.Product{
		Id:       int64(product.ID),
		Slug:     product.Slug,
		Title:    product.Title,
		Body:     product.Body,
		Price:    float32(product.Price),
		CreateAt: product.CreatedAt.String(),
		UpdateAt: product.UpdatedAt.String(),
		DeleteAt: "",
	}}, nil
}
