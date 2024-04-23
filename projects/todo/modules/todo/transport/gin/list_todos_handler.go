package transport

import (
	"fmt"
	"net/http"
	"todo/common"
	"todo/modules/todo/business"
	"todo/modules/todo/entity"
	"todo/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTodoItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fmt.Println("--------------- HANDLER")

		// Parse and bind paging params
		var paging common.Paging
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		paging.Process()

		// Parse and bind paging filter
		var filter entity.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		biz := business.NewListTodosBiz(store)
		result, err := biz.ListTodos(ctx.Request.Context(), &filter, &paging)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, common.SuccessResponse(result, paging, nil))
	}
}
