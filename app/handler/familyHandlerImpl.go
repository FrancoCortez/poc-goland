package handler

import (
	"cenco-pim/app/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type FamilyHandlerImpl struct {
	familyService service.FamilyService
}

func NewFamilyHandlerImpl(familyService service.FamilyService) FamilyHandler {
	return &FamilyHandlerImpl{familyService: familyService}
}

func (f FamilyHandlerImpl) GetAllFamilies(w http.ResponseWriter, r *http.Request) {
	//if err := json.NewEncoder(w).Encode("olas"); err != nil {
	//	//app.logger.Error().Err(err).Msg("")
	//	//e.AppError(w, e.ErrJsonCreationFailure)
	//	return
	//}
	families, _ := f.familyService.GetAllFamilies()
	fmt.Println(families)
	if err := json.NewEncoder(w).Encode(families); err != nil {
		return
	}
}
