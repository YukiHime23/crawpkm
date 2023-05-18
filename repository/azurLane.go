package repository

import (
	"context"
	"goCraw/model"
	"gorm.io/gorm"
)

func SaveFile(ctx context.Context, db *gorm.DB, list []model.AzurLane) {
	db.Save(&list)
}
