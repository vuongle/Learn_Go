package req

type ReqUserUpdate struct {
	Fullname string `json:"full_name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
}
