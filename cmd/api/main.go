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

	// postgresql connection
	// Obtén las variables de entorno
	user := os.Getenv("POSTGRE_USER")
	dbname := os.Getenv("POSTGRE_DB")
	password := os.Getenv("POSTGRE_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	// Crea la cadena de conexión
	postgreUri := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)

	postgresdb, err := sql.Open("postgres", postgreUri)

	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongdbUri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	defer postgresdb.Close()

	mongodb := client.Database("boardmanagement")

	// Retorna una instacia del router que vamos a utilizar
	router := gin.Default()

	boardQueryRepo := mongorepo.NewMongoBoardQueryRepository(mongodb)
	boardQueryService := services.NewBoardQueryService(boardQueryRepo)
	boardQueryHandler := handlers.NewBoardHandler(boardQueryService)

	// command db
	boardCommandRepo := postgrerepo.NewPostgreSQLBoardCommandRepository(postgresdb)
	boardCommandService := services.NewBoardCommandService(boardCommandRepo)
	boardCommandHandler := handlers.NewBoardCommandHandler(boardCommandService)

	router.GET("/boards", boardQueryHandler.GetAll)
	router.GET("/boards/:id", boardQueryHandler.GetById)

	router.POST("/boards", boardCommandHandler.Create)

	router.Run() // escuchando en: 0.0.0.0:8080
}
