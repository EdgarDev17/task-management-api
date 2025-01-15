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

type mongoTaskQueryRepository struct {
	collection *mongo.Collection
	logger     repositories.Logger
}

func NewMongoTaskQueryRepository(db *mongo.Database, logger repositories.Logger) repositories.TaskQueryRepositoryI {
	return &mongoTaskQueryRepository{
		collection: db.Collection("tareas"),
		logger:     logger,
	}
}

func (db *mongoTaskQueryRepository) GetAll(ctx context.Context) ([]*models.TaskQuery, error) {
	var tasks []*models.TaskQuery

	opts := options.Find().SetSort(bson.D{{Key: "fecha_creacion", Value: -1}})

	cursor, err := db.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		db.logger.Error("error al buscar las tareas", zap.Error(err))
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.TaskQuery
		if err := cursor.Decode(&task); err != nil {
			db.logger.Error("error al decodificar la tarea", zap.Error(err))
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (db *mongoTaskQueryRepository) GetById(ctx context.Context, id string) (*models.TaskQuery, error) {
	var task models.TaskQuery

	err := db.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		db.logger.Error("error al buscar la tarea por ID", zap.Error(err))
		return nil, err
	}

	return &task, nil
}

func (db *mongoTaskQueryRepository) Upsert(ctx context.Context, task *models.Task) error {
	uuidString := task.ID.String()

	filter := bson.M{"_id": uuidString}
	update := bson.M{
		"$set": bson.M{
			"_id":            uuidString,
			"titulo":         task.Title,
			"descripcion":    task.Description,
			"estado":         task.State,
			"fecha_creacion": task.CreatedAt,
			"id_tablero":     task.BoardID.String(),
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := db.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		db.logger.Error("error al hacer upsert de la tarea", zap.Error(err))
		return fmt.Errorf("error al upsert task: %w", err)
	}

	db.logger.Info("tarea actualizada/insertada",
		zap.Int64("matched", result.MatchedCount),
		zap.Int64("modified", result.ModifiedCount),
		zap.Bool("upserted", result.UpsertedID != nil))

	return nil
}

func (db *mongoTaskQueryRepository) Delete(ctx context.Context, task *models.Task) error {
	filter := bson.M{"_id": task.ID.String()}

	result, err := db.collection.DeleteOne(ctx, filter)
	if err != nil {
		db.logger.Error("error al eliminar la tarea", zap.Error(err))
		return fmt.Errorf("error deleting task from query db: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("task with id %s not found in query db", task.ID)
	}

	return nil
}

func (db *mongoTaskQueryRepository) GetTasksByBoardId(ctx context.Context, boardId string) ([]*models.TaskQuery, error) {
	var tasks []*models.TaskQuery

	opts := options.Find().SetSort(bson.D{{Key: "fecha_creacion", Value: -1}})
	filter := bson.M{"id_tablero": boardId}

	cursor, err := db.collection.Find(ctx, filter, opts)

	if err != nil {
		db.logger.Error("error al buscar las tareas por id de tablero", zap.Error(err))
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.TaskQuery
		if err := cursor.Decode(&task); err != nil {
			db.logger.Error("error al decodificar la tarea", zap.Error(err))
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}
