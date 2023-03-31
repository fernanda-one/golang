package Dto

type BookUpdateDto struct {
	ID          uint64 `json:"ID" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}

type BookCreateDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" from:"user_id,omitempty"`
}
