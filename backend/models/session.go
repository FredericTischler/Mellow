package models

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	SessionID    uuid.UUID `json:"session_id"`
	UserID       uuid.UUID `json:"user_id"`
	CreationDate time.Time `json:"creation_date"`
	LastActivity time.Time `json:"last_activity"`
}
