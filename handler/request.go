package handler

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type SaveImageRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"-"`
}

func (r *SaveImageRequest) Validate() error {
	if r.Title == "" && r.Description == "" && r.Image == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Image != nil && !strings.Contains(r.Image.Header.Get("Content-Type"), "image/") {
		return fmt.Errorf("mimetype invalid")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Image == nil {
		return errParamIsRequired("image", "image file")
	}
	return nil
}
