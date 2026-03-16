package models

import (
	"time"
)

type Weather struct {
	// Название города
	City string `json:"city"`
	// Температура в градусах
	Temperature float64 `json:"temperature"`
	// Время последнего обновления данных
	UpdatedAt time.Time `json:"updated_at"`
}
