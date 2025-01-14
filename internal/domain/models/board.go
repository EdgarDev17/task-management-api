package models

import (
	"time"

	"github.com/google/uuid"
)

// Board representa la estructura para ambas bases de datos
// Para escritura
type Board struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"nombre" gorm:"column:nombre"`
	Description string    `json:"descripcion" gorm:"column:descripcion"`
	CreatedAt   time.Time `json:"fecha_creacion" gorm:"column:fecha_creacion;default:CURRENT_TIMESTAMP"`
}

// Para lectura
type BoardQuery struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"nombre" bson:"nombre"`
	Description string    `json:"descripcion" bson:"descripcion"`
	CreatedAt   time.Time `json:"fecha_creacion" bson:"fecha_creacion"`
}
