package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yago-123/stackit/pkg/storage"
	"github.com/yago-123/stackit/pkg/types"
)

type handler struct {
	store storage.Store
}

func newHandler(store storage.Store) *handler {
	return &handler{
		store: store,
	}
}

func (h *handler) PostUser(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON payload",
		})
		return
	}

	// todo(): handle validation of data etc

	if err := h.store.PersistUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to persist user",
		})
		return
	}

	c.Status(http.StatusOK)
}

// Disclaimer: if the number of users is large, this will not be efficient. Will need pagination
func (h *handler) GetUsers(c *gin.Context) {
	allUsers, err := h.store.RetrieveUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, allUsers)
}
