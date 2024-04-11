package validation

type AddQuestionRequest struct {
	ID         int    `json:"id" form:"id"`
	Question   string `json:"question" valid:"required"`
	CategoryId string `form:"category_id" json:"category_id" validate:"required"`
}
