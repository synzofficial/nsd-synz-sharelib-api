package ginutil

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validatorInst *validator.Validate

func InitValidatorInst(validate *validator.Validate) {
	validatorInst = validate
}

func NewDefaultValidatorInst() *validator.Validate {
	return validator.New()
}

func BindingRequestLogging[Request any](c *gin.Context, request Request) (context.Context, Request, error) {
	if err := c.ShouldBind(&request); err != nil {
		return nil, request, err
	}

	if err := c.ShouldBindUri(&request); err != nil {
		return nil, request, err
	}

	if err := c.ShouldBindHeader(&request); err != nil {
		return nil, request, err
	}

	// set request logging {log_message, log_masking etc..}

	// extract request header, method, uri to context
	ctx := ContextWithRequest(c.Request.Context(), c.Request)

	// validate struct using validator instead.
	if validatorInst == nil {
		InitValidatorInst(NewDefaultValidatorInst())
	}

	if err := validatorInst.Struct(request); err != nil {
		return nil, request, err
	}

	return ctx, request, nil
}
