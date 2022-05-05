package service

import (
	"cenco-pim/app/entity"
	"cenco-pim/app/repository"
)

type FamilyServiceImpl struct {
	familyRepository repository.FamilyRepository
}

func NewFamilyServiceImpl(familyRepository repository.FamilyRepository) FamilyService {
	return &FamilyServiceImpl{familyRepository: familyRepository}
}

func (f FamilyServiceImpl) GetAllFamilies() (entity.FamilyDtos, error) {
	families, error := f.familyRepository.GetAllFamilies()
	if families == nil {
		return nil, nil
	}
	return families.ToDto(), error
}
