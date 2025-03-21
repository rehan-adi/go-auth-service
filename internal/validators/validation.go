package validators

import "github.com/go-playground/validator/v10"

type SignupValidator struct {
	Username string `json:"username" validate:"required,min=2,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SigninValidator struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
}

var validate = validator.New()

func ValidateSignupData(input SignupValidator) map[string]string {
	return validateStruct(input)
}

func ValidateSigninData(input SigninValidator) map[string]string {
	return validateStruct(input)
}

func ValidateUpdateUserData(input UpdateUserRequest) map[string]string {
	return validateStruct(input)
}

func validateStruct(input interface{}) map[string]string {
	errs := make(map[string]string)
	if err := validate.Struct(input); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = "Invalid " + e.Field()
		}
	}
	return errs
}
