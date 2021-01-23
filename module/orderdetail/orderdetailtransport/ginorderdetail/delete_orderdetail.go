package ginorderdetail

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailbusiness"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailstorage"
)

// DeleteOrderDetail a orderdetail
func DeleteOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("ord-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewDeleteOrderDetailBiz(store)

		if err := bizOrderDetail.DeleteOrderDetail(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
