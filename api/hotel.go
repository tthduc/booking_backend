package api

import (
	db "booking-backed/db/sqlc"
	"booking-backed/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createHotel(ctx *gin.Context) {
	var req models.CreateHotelParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateHotelParams{
		Name:     req.Name,
		Address:  req.Address,
		Location: req.Location,
	}

	hotel, err := server.store.CreateHotel(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

func (server *Server) getHotel(ctx *gin.Context) {
	var req models.GetHotelParams
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hotel, err := server.store.GetHotel(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, hotel)
}

func (server *Server) listHotel(ctx *gin.Context) {
	var req models.ListHotelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListHotelsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListHotels(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
