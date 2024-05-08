package delivery

import (
	"testfastprint/src/usecase"

	"github.com/gin-gonic/gin"
)

type Delivery interface {
	InsertProductMany(c *gin.Context)
	InsertOneProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
	GetProdukForUpdate(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	LoadInsertView(c *gin.Context)
}

type delivery struct {
	usecase usecase.Usecase
	linkAPI string
}

func NewDelivery(usecase usecase.Usecase, linkAPI string) Delivery {
	return &delivery{
		usecase: usecase,
		linkAPI: linkAPI,
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
