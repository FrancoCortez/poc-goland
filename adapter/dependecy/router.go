package dependecy

import (
	"cenco-pim/app/router/middleware"
	"cenco-pim/app/service/health"
	"cenco-pim/util/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func GetRouter(c *InitDependencyImpl, l *logger.Logger, v *validator.Validate, db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/healthz", health.HandleHealth)
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)
		r.Mount("/families", c.GetFamilyRoute().Routers(l))
		//	// Routes for books
		//	//srvBook := book.NewApp(l, v, db)
		//	//r.Method("GET", "/books", requestlog.NewHandler(srvBook.HandleListBooks, l))
		//	//r.Method("POST", "/books", requestlog.NewHandler(srvBook.HandleCreateBook, l))
		//	//r.Method("GET", "/books/{id}", requestlog.NewHandler(srvBook.HandleReadBook, l))
		//	//r.Method("PUT", "/books/{id}", requestlog.NewHandler(srvBook.HandleUpdateBook, l))
		//	//r.Method("DELETE", "/books/{id}", requestlog.NewHandler(srvBook.HandleDeleteBook, l))
	})

	return r
}
