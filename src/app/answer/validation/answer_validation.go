package validation

type AddAnswerRequest struct {
	ID         int    `json:"id" form:"id"`
	Option     string `json:"option" valid:"required"`
	Answer     string `json:"answer" valid:"required"`
	Score      int    `json:"score" valid:"required"`
	QuestionId string `form:"question_id" json:"question_id" validate:"required"`
}
