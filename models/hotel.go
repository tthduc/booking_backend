package models

type CreateHotelParams struct {
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type GetHotelParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListHotelRequest struct {
	PageID   int32 `json:"page_id" binding:"required,min=1"`
	PageSize int32 `json:"page_size" binding:"required,min=5,max=10"`
}
