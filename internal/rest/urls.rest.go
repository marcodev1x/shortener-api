package rest

import (
	"fmt"
	"shortner-url/internal"
	"shortner-url/internal/structs"
	"shortner-url/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UrlRest struct {
	usecase *usecases.UrlUsecase
}

func NewUrlRest(usecases *usecases.UrlUsecase) *UrlRest {
	return &UrlRest{usecases}
}

func (r *UrlRest) FindUrlByHashedId(c *gin.Context) {
	var request structs.ById

	if err := internal.BindJSON(c, &request); err != nil {
		c.Error(err)
		return
	}

	url, err := r.usecase.FindUrlByHashedId(request.Id)

	if err != nil || url == nil {
		c.Error(err)
		return
	}

	c.JSON(200, gin.H{
		"url": url.ShortenedUrl,
	})

}

func (r *UrlRest) CreateUrl(c *gin.Context) {
	var request structs.CreateUrl

	if err := internal.BindJSON(c, &request); err != nil {
		c.Error(err)
		return
	}

	created, err := r.usecase.CreateUrl(request.Url, request.ExpiresAt)

	fmt.Println(err, created)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, gin.H{
		"created": created,
	})
}
