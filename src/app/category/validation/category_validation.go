package validation

type AddCategoryRequest struct {
	Name  string `json:"name" valid:"required"`
	Order uint `json:"order" valid:"required"`
}
