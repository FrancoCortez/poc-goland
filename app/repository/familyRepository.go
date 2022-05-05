package repository

import (
	"cenco-pim/app/entity"
)

type FamilyRepository interface {
	GetAllFamilies() (entity.Families, error)
}
