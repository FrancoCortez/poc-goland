package repository

import (
	"cenco-pim/app/entity"
	"cenco-pim/util/logger"
	"gorm.io/gorm"
)

type FamilyRepositoryImpl struct {
	logger *logger.Logger
	db     *gorm.DB
}

func NewFamilyRepository(db *gorm.DB, logger *logger.Logger) FamilyRepository {
	return &FamilyRepositoryImpl{
		logger: logger,
		db:     db,
	}
}

func (r *FamilyRepositoryImpl) GetAllFamilies() (entity.Families, error) {
	families := make([]*entity.Family, 0)
	r.db = r.db.Table("families")
	r.db = r.db.Joins("INNER JOIN ")
	if err := r.db.Find(&families).Error; err != nil {
		return nil, err
	}
	if len(families) == 0 {
		return nil, nil
	}
	return families, nil
}
