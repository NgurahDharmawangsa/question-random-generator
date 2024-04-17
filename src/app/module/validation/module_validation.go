package validation

type AddModuleRequest struct {
	ID          int    `json:"id" form:"id"`
	Name        string `json:"name" valid:"required"`
}
