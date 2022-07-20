package v1

import (
	"GRPC-project/pkg/api/proto/v1"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "movie"
)

type transactionServiceServer struct {
	v1.UnimplementedTransactionServiceServer
	db *sqlx.DB
}

func NewMovieServiceServer(db *sqlx.DB) v1.TransactionServiceServer {
	return &transactionServiceServer{db: db}
}

func (s *transactionServiceServer) CreateCustomer(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("INSERT INTO Customer (CustomerName, Phone, Email, Product) VALUES (?, ?, ?, ?)", request.Customer.CustomerName, request.Customer.Phone, request.Customer.Email, request.Customer.Product)
	tx.Commit()
	return &v1.CreateResponse{
		Api:      apiVersion,
		Customer: request.Customer,
	}, nil
}

func (s *transactionServiceServer) GetAllCustomer(ctx context.Context, request *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	rows, err := s.db.Queryx("SELECT CustomerName, Phone, Email, Product FROM Customer")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to query-> "+err.Error())
	}

	var customers []*v1.Customer
	for rows.Next() {
		var m v1.Customer
		if err := rows.Scan(&m.CustomerName, &m.Phone, &m.Email, &m.Product); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from row-> "+err.Error())
		}
		customers = append(customers, &m)
	}
	return &v1.ReadAllResponse{
		Api:      apiVersion,
		Customer: customers,
	}, nil
}

type responseForGetCustomerByProduct struct {
	TransactionID int32  `db:"TransactionID"`
	CustomerName  string `db:"CustomerName"`
	Phone         string `db:"Phone"`
	Email         string `db:"Email"`
	Product       string `db:"Product"`
}

func (s *transactionServiceServer) GetCustomerByProduct(ctx context.Context, request *v1.ReadRequest) (*v1.ReadResponse, error) {
	start := time.Now()
	tx := s.db.MustBegin()
	listMovie := "SELECT TransactionID, CustomerName, Phone, Email, Product FROM Customer WHERE Product = ?"
	var queryAns []*responseForGetCustomerByProduct
	err := tx.SelectContext(ctx, &queryAns, listMovie, request.Product)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to query-> "+err.Error())
	}

	var customers []*v1.Customer
	for _, v := range queryAns {
		var m v1.Customer
		m.TransactionID = v.TransactionID
		m.CustomerName = v.CustomerName
		m.Phone = v.Phone
		m.Email = v.Email
		m.Product = v.Product
		customers = append(customers, &m)
	}
	tx.Commit()
	fmt.Println("Get movie by genre executed in: ", time.Since(start))
	return &v1.ReadResponse{
		Api:      apiVersion,
		Customer: customers,
	}, nil
}

func (s *transactionServiceServer) UpdateCustomer(ctx context.Context, request *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("UPDATE Customer SET Phone = ?, Email = ?, Product = ? WHERE CustomerName = ?", "567-912-3513", "1980@fb.com", "Nước hoa", request.CustomerName)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateResponse{
		Api: apiVersion,
	}, nil
}

func (s *transactionServiceServer) DeleteCustomer(ctx context.Context, request *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(request.Api); err != nil {
		return nil, err
	}
	tx := s.db.MustBegin()
	tx.MustExec("DELETE FROM Customer WHERE CustomerName = ?", request.Product)
	err := tx.Commit()
	if err != nil {
		return nil, err
	}
	return &v1.DeleteResponse{
		Api: apiVersion,
	}, nil
}

// checkAPI checks if the API version requested by client is supported by server
func (s *transactionServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *transactionServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}
