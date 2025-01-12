package mongorepo

import (
	"context"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Esta es la implementacion concreta del contrato
// para el repositorio de la db command
type MongoBoardQueryRepository struct {
	collection *mongo.Collection
}

// Funcion que sirve como constructor
// Esta funcion debe retornar una implementacion que cumpla con el contrato query
func NewMongoBoardQueryRepository(db *mongo.Database) repositories.BoardQueryRepositoryI {
	return &MongoBoardQueryRepository{
		collection: db.Collection("tableros"),
	}
}

func (db *MongoBoardQueryRepository) GetAll(ctx context.Context) ([]*models.Board, error) {
	var boards []*models.Board

	cursor, err := db.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	// este me asegura que el cursor se una vez haya concluido su uso
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var board models.Board
		if err := cursor.Decode(&board); err != nil {
			return nil, err
		}

		boards = append(boards, &board)
	}

	return boards, nil
}

func (db *MongoBoardQueryRepository) GetById(ctx context.Context, id string) (*models.Board, error) {
	var board models.Board

	// Convierte el uuid a un ObjectId
	// Ya que ObjectId es el standard para IDs en mongo
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error al convertir UUID a ObjectId:", err)
		return nil, err
	}

	// Realiza la consulta utilizando el ObjectId
	err = db.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&board)
	if err != nil {
		fmt.Println("Error en el repositorio mongodb:", err)
		return nil, err
	}

	fmt.Println("Dato de mongodb: ", &board)
	return &board, nil

}
