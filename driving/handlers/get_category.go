package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rifame/driven"
	"rifame/driven/repository"
	"rifame/driving/response"
)

func FindAll(driven driven.Driven) func(ctx *gin.Context) {
	log.Println("get categories")
	repo := repository.CategoryRepository{DB: driven.DB}

	return func(ctx *gin.Context) {
		tags, err := repo.FindAll()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}

		ctx.JSON(http.StatusOK, response.Collection{List: tags, Length: len(tags)})
	}
}
