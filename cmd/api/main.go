package main

import (
	_ "github.com/lib/pq"

	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"task-management-api/internal/infrastructure/database/mongorepo"
	"task-management-api/internal/infrastructure/database/postgrerepo"
	"task-management-api/internal/interfaces/handlers"
	"task-management-api/internal/usecases/services"
	"task-management-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongdbUri := os.Getenv("MONGODB_URI")
	user := os.Getenv("POSTGRE_USER")
	dbname := os.Getenv("POSTGRE_DB")
	password := os.Getenv("POSTGRE_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	postgreUri := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)
	postgresdb, err := sql.Open("postgres", postgreUri)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongdbUri))
	if err != nil {
		panic(err)
	}

	// Initialize cleanup handlers
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	defer postgresdb.Close()

	mongodb := client.Database("boardmanagement")
	zapLogger := logger.NewZapLogger()

	// Initialize repositories
	boardQueryRepo := mongorepo.NewMongoBoardQueryRepository(mongodb, zapLogger)

	// Stores
	eventStore, err := postgrerepo.NewPostgresEventStore(postgresdb)

	if err != nil {
		log.Fatal(err)
	}

	// Inicializo el manejador de eventos (de stores)
	eventHandler := handlers.NewEventHandler(eventStore, boardQueryRepo, zapLogger)
	eventHandler.Start()

	// Initialize services and handlers
	boardQueryService := services.NewBoardQueryService(boardQueryRepo)
	boardQueryHandler := handlers.NewBoardHandler(boardQueryService)

	boardCommandRepo := postgrerepo.NewPostgreSQLBoardCommandRepository(postgresdb, eventStore, zapLogger)
	boardCommandService := services.NewBoardCommandService(boardCommandRepo, zapLogger)
	boardCommandHandler := handlers.NewBoardCommandHandler(boardCommandService)

	// Set up router
	router := gin.Default()
	router.GET("/api/v1/boards", boardQueryHandler.GetAll)
	router.GET("/api/v1/boards/:id", boardQueryHandler.GetById)
	router.POST("/api/v1/boards", boardCommandHandler.Create)
	router.PUT("/api/v1/boards", boardCommandHandler.Update)
	router.DELETE("/api/v1/boards/:id", boardCommandHandler.Delete)

	router.Run()
}
