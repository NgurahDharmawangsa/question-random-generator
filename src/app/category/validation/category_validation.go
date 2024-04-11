package validation

type AddCategoryRequest struct {
	Name  string `json:"name" valid:"required"`
	Order string `json:"order" valid:"required"`
}
