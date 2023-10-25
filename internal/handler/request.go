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

type UpdateImageRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *UpdateImageRequest) Validate() error {
	if r.Title != "" || r.Description != "" {
		return nil
	}
	// If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignUpRequest) Validate() error {
	if r.Username == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *SignInRequest) Validate() error {
	if r.Email == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}
