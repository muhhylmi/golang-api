package web

type (
	RequestCreateUser struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Gender   string `json:"gender" validate:"required,oneof=M F"`
	}
)
