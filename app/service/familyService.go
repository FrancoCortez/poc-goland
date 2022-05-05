package service

import "cenco-pim/app/entity"

type FamilyService interface {
	GetAllFamilies() (entity.FamilyDtos, error)
}
