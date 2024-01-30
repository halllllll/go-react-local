package handlerrrr

// import (
// 	"fmt"
// 	"net/http"
// 	"sample/go-react-local-app/repository"

// 	"github.com/gin-gonic/gin"
// )

// type countPostRequest struct {
// 	Count int `json:"count"`
// }

// type ApiRegister struct {
// 	Repo repository.Repository
// }

// func (a *ApiRegister) GetCount(c *gin.Context) {
// 	if val, err := a.Repo.Get(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"success": true, "count": val})
// 	}
// }

// func (a *ApiRegister) SetCount(c *gin.Context) {
// 	reqData := &countPostRequest{}
// 	if err := c.ShouldBindJSON(reqData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	fmt.Printf("count: %d\n", reqData.Count)
// 	// 登録
// 	if err := a.Repo.Set(reqData.Count); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "newCount": reqData.Count})
// }
