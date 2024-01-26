package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type countPostRequest struct {
	Count int `json:"count"`
}

// type SetCountService interface {
// 	SetCount(ctx context.Context, count int) error
// }

// type RegisterApi struct{
// 	Service SetCountService
// }

type ApiRegister struct {
	DB  *sql.DB
	Ctx context.Context
}

type Service struct {
	tx ApiRegister
}

func (a *ApiRegister) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "yoyo~!"})
}

func (a *ApiRegister) HelloPost(c *gin.Context) {
	reqData := &countPostRequest{}
	if err := c.ShouldBindJSON(reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	fmt.Printf("count: %d\n", reqData.Count)
	// 登録
	tx, err := a.DB.BeginTx(a.Ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var count int
	if err = tx.QueryRowContext(a.Ctx, "SELECT count FROM count WHERE id = $1", 1).Scan(&count); err != nil && err == sql.ErrNoRows {
		_, err := tx.ExecContext(a.Ctx, "INSERT INTO count(count) VALUES ($1)", reqData.Count)
		if err != nil {
			return
		}
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	// update
	_, err = tx.ExecContext(a.Ctx, "UPDATE count SET count = $1 WHERE id = $2", reqData.Count, 1)
	if err != nil {
		log.Println(err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "hai"})
}
