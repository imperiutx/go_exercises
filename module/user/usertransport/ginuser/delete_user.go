package ginuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// DeleteUser a user
func DeleteUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("usr-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)
		bizUser := userbusiness.NewDeleteUserBiz(store)

		if err := bizUser.DeleteUser(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "deleted"})
	}
}
