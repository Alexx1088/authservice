package main

import (
	"database/sql"
	"fmt"
	"github.com/Alexx1088/authservice/internal/db"
	"github.com/Alexx1088/authservice/internal/migrate"
	"github.com/Alexx1088/authservice/internal/repository"
	"github.com/Alexx1088/authservice/internal/service"
	pb "github.com/Alexx1088/authservice/proto"
	userpb "github.com/Alexx1088/userservice/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	migrationsPath := "file://migrations"

	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	if err := migrate.RunMigrations(migrationsPath, dbURL); err != nil {
		log.Fatalf("Migration failed: %v", err)
		log.Println("Connected to DB and migration applied successfully.")
	}

	// Connect to UserService
	conn, err := grpc.Dial("userservice:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to UserService: %v", err)
	}
	defer conn.Close()
	userClient := userpb.NewUserServiceClient(conn)

	// Initialize database/repository here (example)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()
	repo := repository.NewAuthRepository(db) // implement this if not yet done

	// Create AuthService server
	authServer := &service.AuthServiceServer{
		Repo:        repo,
		UserService: userClient,
	}

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, authServer)

	log.Println("AuthService is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
