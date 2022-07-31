package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "transactions/db/sqlc"
	"transactions/util"
	"github.com/Straycats/redismq"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	Consname := util.RandomOwner()
	testQueue := redismq.CreateQueue("localhost", "6379", "", 0, "clicks")
	testQueue.Put("1")
	consumer, err := testQueue.AddConsumer(Consname)
	if err != nil {
		panic(err)
	}

	
	
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.FromAccountID, req.Currency) {
		return
	}

	if !server.validAccount(ctx, req.ToAccountID, req.Currency) {
		return
	}


	if !server.validMoneyAccount(ctx, req.ToAccountID) {
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

		p, err := consumer.Get()
		if err != nil {
			fmt.Println(err)
		}
		if (p.Payload == "1"){
			ctx.JSON(http.StatusOK, result)
		}
		err = p.Ack()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Aded Queue")
		consumer.Quit()
		consumer, err = testQueue.AddConsumer(Consname)
		if err != nil {
			panic(err)
		}
	ctx.JSON(http.StatusOK, result)
}

func(server *Server) validMoneyAccount(ctx *gin.Context, accountID int64) bool{
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil{
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
	}
}
if account.Balance <= 0 {
	err := fmt.Errorf("account [%d] balance is not sufficient to transfer : %d ", account.ID, account.Balance)
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return false
}
return true
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
