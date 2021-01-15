package ginuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// GetUser a user
func GetUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("usr-id")
		id, err := strconv.Atoi(idString)

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)

		bizUser := userbusiness.NewGetUserBiz(store)
		user, err := bizUser.GetUser(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
