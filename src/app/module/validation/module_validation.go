package validation

import "github.com/lib/pq"

type AddModuleRequest struct {
	ID          int           `json:"id" form:"id"`
	Identifier  string        `json:"identifier"`
	Name        string        `json:"name" valid:"required"`
	QuestionIds pq.Int64Array `json:"question_ids"`
}
