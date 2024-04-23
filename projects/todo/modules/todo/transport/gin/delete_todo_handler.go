package transport

import (
	"net/http"
	"strconv"
	"todo/common"
	"todo/modules/todo/business"
	"todo/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTodoItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Get param id from request and validate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		biz := business.NewDeleteTodoBiz(store)
		if err := biz.DeleteTodoById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SingleSuccessResponse(true))
	}
}
