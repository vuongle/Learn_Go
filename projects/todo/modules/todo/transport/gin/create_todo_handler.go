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

func CreateTodoItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fmt.Println("--------------- HANDLER")
		var data entity.TodoCreationBody

		// the above UnmarshalJSON() method is AUTO called here(inside flow of ShouldBind)
		err := ctx.ShouldBind(&data) // pass pointer of data (not pass data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		biz := business.NewCreateTodoBiz(store) // create an instance of business struct
		if err := biz.CreateTodo(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.SingleSuccessResponse(data.Id))
	}
}
