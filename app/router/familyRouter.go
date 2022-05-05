package router

import (
	"cenco-pim/app/handler"
	"cenco-pim/util/logger"
	"cenco-pim/util/requestlog"
	"github.com/go-chi/chi/v5"
)

type Family struct {
	familyHandler handler.FamilyHandler
	logger        *logger.Logger
}

func NewFamily(familyHandler handler.FamilyHandler, logger *logger.Logger) Family {
	return Family{
		familyHandler: familyHandler,
		logger:        logger,
	}
}

func (f *Family) Routers(l *logger.Logger) chi.Router {
	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(f.familyHandler.GetAllFamilies, l))
	return r
}
