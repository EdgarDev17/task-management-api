package models

import (
	"time"

	"github.com/google/uuid"
)

// Board representa la estructura para ambas bases de datos
type Board struct {
	ID          uuid.UUID `json:"id" bson:"_id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"nombre" bson:"nombre" gorm:"column:nombre"`
	Description string    `json:"descripcion" bson:"descripcion" gorm:"column:descripcion"`
	CreatedAt   time.Time `json:"fecha_creacion" bson:"fecha_creacion" gorm:"column:fecha_creacion;default:CURRENT_TIMESTAMP"`
}

// ToMongo convierte el UUID de PostgreSQL a ObjectID para MongoDB
func (b *Board) ToMongo() *Board {
	mongoBoard := *b
	mongoBoard.ID = uuid.UUID{} // Limpiar UUID
	return &mongoBoard
}

// TableName especifica el nombre de la tabla en PostgreSQL
func (Board) TableName() string {
	return "boards"
}
