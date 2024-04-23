package transport

import (
	"net/http"
	"strconv"
	"todo/common"
	"todo/modules/todo/business"
	"todo/modules/todo/entity"
	"todo/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateTodoItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		// Get param id from request and validate
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		// When the param id is ok -> parse data from request
		var data entity.TodoUpdateBody
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// update data to db
		store := storage.NewSQLStore(db)
		biz := business.NewUpdateTodoBiz(store)
		if err := biz.UpdateTodoById(ctx.Request.Context(), id, &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SingleSuccessResponse(true))
	}
}
