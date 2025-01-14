package mongorepo

import (
	"context"
	"fmt"
	"task-management-api/internal/domain/models"
	"task-management-api/internal/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// Esta es la implementacion concreta del contrato
// para el repositorio de la db command
type MongoBoardQueryRepository struct {
	collection *mongo.Collection
	logger     repositories.Logger
}

// Funcion que sirve como constructor
// Esta funcion debe retornar una implementacion que cumpla con el contrato query
func NewMongoBoardQueryRepository(db *mongo.Database, logger repositories.Logger) repositories.BoardQueryRepositoryI {
	return &MongoBoardQueryRepository{
		collection: db.Collection("tableros"),
		logger:     logger,
	}
}

func (db *MongoBoardQueryRepository) GetAll(ctx context.Context) ([]*models.BoardQuery, error) {
	var boards []*models.BoardQuery

	cursor, err := db.collection.Find(ctx, bson.M{})

	if err != nil {
		db.logger.Error("error al buscar los documentos", zap.Error(err))
		return nil, err
	}

	// este me asegura que el cursor se una vez haya concluido su uso
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var board models.BoardQuery
		if err := cursor.Decode(&board); err != nil {
			db.logger.Error("error al  DECODE el documento", zap.Error(err))
			return nil, err
		}

		boards = append(boards, &board)
	}

	return boards, nil
}

func (db *MongoBoardQueryRepository) GetById(ctx context.Context, id string) (*models.BoardQuery, error) {
	var board models.BoardQuery

	err := db.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&board)

	if err != nil {
		fmt.Println("Error en el repositorio mongodb:", err)
		return nil, err
	}

	return &board, nil
}

func (db MongoBoardQueryRepository) Upsert(ctx context.Context, board *models.Board) error {
	// convierto el UUID a su representación en string
	uuidString := board.ID.String()

	filter := bson.M{"_id": uuidString}

	// En el update incluyo el _id explícitamente
	update := bson.M{
		"$set": bson.M{
			"_id":         uuidString, // Incluimos el ID en el set
			"nombre":      board.Name,
			"descripcion": board.Description,
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := db.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("error al upsert board: %w", err)
	}

	fmt.Printf("Matched: %d, Modified: %d, Upserted: %v\n",
		result.MatchedCount,
		result.ModifiedCount,
		result.UpsertedID != nil)

	return nil
}

func (db *MongoBoardQueryRepository) Delete(ctx context.Context, board *models.Board) error {
	filter := bson.M{"_id": board.ID.String()}
	result, err := db.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting board from query db: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("board with id %s not found in query db", board.ID)
	}

	return nil
}
