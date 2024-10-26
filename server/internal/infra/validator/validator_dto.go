package infra_validator

import "github.com/go-playground/validator/v10"

type Error struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func CustomValidator(data interface{}) []*Error {
	v := validator.New()

	err := v.Struct(data)

	errors := make([]*Error, 0)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el Error
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}

		return errors
	}

	return nil
}
