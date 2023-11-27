package src

import (
	"database/sql"
	"log"
	"testfastprint/src/delivery"
	"testfastprint/src/repository"
	"testfastprint/src/usecase"
	"testfastprint/util"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start()
}

type server struct {
	httpServer *gin.Engine
	db         *sql.DB
	port       string
	linkAPI    string
}

func InitServer(db *sql.DB, env *util.Config) Server {

	g := gin.Default()

	return &server{
		httpServer: g,
		db:         db,
		port:       env.HTTPServerAddress,
		linkAPI:    env.LinkAPI,
	}

}

func (s *server) Start() {

	repo := repository.NewStore(s.db)
	usecase := usecase.NewUsecase(repo)
	delivery := delivery.NewDelivery(usecase, s.linkAPI)
	s.httpServer.POST("/products", delivery.InsertProduct)
	s.httpServer.GET("/products", delivery.GetAllProduct)
	s.httpServer.GET("/product/:id", delivery.GetProduk)
	s.httpServer.PUT("/product/:id", delivery.UpdateProduct)
	s.httpServer.DELETE("/product/:id", delivery.DeleteProduct)

	err := s.httpServer.Run(s.port)
	if err != nil {
		log.Fatal(err)
	}

}
