package customercontroller

type CreateCustomCausesRequest struct {
	Name string `json:"name" binding:"required"`
}
