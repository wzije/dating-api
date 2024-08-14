package src

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Home(c *fiber.Ctx) error
}

type handler struct {
	database Database
}

func (c handler) Register(ctx *fiber.Ctx) error {
	request := new(RegisterRequest)

	if err := ctx.BodyParser(request); err != nil {
		return JsonError(ctx, fiber.StatusBadRequest, err)
	}

	if err := ValidateStruct(request); err != nil {
		return JsonErrorWithReason(ctx, fiber.StatusBadRequest, "invalid param", err)
	}

	if request.Password != request.PasswordConfirmation {
		return JsonError(ctx, fiber.StatusBadRequest, errors.New("password does not match"))
	}

	result, _ := c.database.GetUserByEmail(request.Email)

	if result != nil && result.Email == request.Email {
		return JsonError(ctx, fiber.StatusBadRequest, errors.New("email is already registered"))
	}

	hashedPassword, err := HashPassword(request.Password)

	if err != nil {
		return JsonError(ctx, fiber.StatusBadRequest, errors.New("unable to hash password"))
	}

	user := User{
		Email:    request.Email,
		Password: hashedPassword,
		Name:     request.Name,
		Gender:   request.Gender,
		Address:  request.Address,
	}

	_, err = c.database.StoreUser(user)

	if err != nil {
		return JsonError(ctx, fiber.StatusInternalServerError, err)
	}

	//if there is error
	return Json(ctx, Response{
		Code:    fiber.StatusCreated,
		Message: "register successful",
		Data:    nil,
	})
}

func (c handler) Login(ctx *fiber.Ctx) error {
	request := new(LoginRequest)

	//parsing body param to request
	if err := ctx.BodyParser(&request); err != nil {
		return JsonError(ctx, fiber.StatusBadRequest, err)
	}

	//validate request
	if err := ValidateStruct(request); err != nil {
		return JsonErrorWithReason(ctx, fiber.StatusBadRequest, "invalid param", err)
	}

	//get by email and check correct password
	user, err := c.database.GetUserByEmail(request.Email)

	if err != nil {
		return JsonError(ctx, fiber.StatusInternalServerError, err)
	}

	if !VerifyPassword(request.Password, user.Password) {
		return JsonError(ctx, fiber.StatusBadRequest, errors.New("password does not match"))
	}

	token, err := GenerateToken(user)
	if err != nil {
		return JsonError(ctx, fiber.StatusInternalServerError, err)
	}

	return Json(ctx, Response{
		Code:    fiber.StatusOK,
		Data:    token,
		Message: "login successful",
	})
}

func (c handler) Home(ctx *fiber.Ctx) error {
	user := PayloadData.User

	return Json(ctx, Response{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    user,
	})
}

func NewHandler(database Database) Handler {
	return &handler{
		database: database,
	}
}
