package ginrestaurantfood

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodstorage"
)

// DeleteRestaurantFood a restaurantfood
func DeleteRestaurantFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		
		ridString := c.Param("rid")
		rid, err := strconv.Atoi(ridString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fidString := c.Param("fid")
		fid, err := strconv.Atoi(fidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantfoodstorage.NewSQLStore(db)
		bizRestaurantFood := restaurantfoodbusiness.NewDeleteRestaurantFoodBiz(store)

		if err := bizRestaurantFood.DeleteRestaurantFood(c.Request.Context(), rid, fid); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
