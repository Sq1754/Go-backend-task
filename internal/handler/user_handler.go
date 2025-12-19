package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/sq1754/user-age-api/internal/repository"
	"github.com/sq1754/user-age-api/internal/service"
)

type UserHandler struct {
	repo     *repository.UserRepository
	validate *validator.Validate
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo:     repo,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid dob format")
	}

	user, err := h.repo.Create(c.Context(), req.Name, dob)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create user")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	user, err := h.repo.GetByID(c.Context(), int32(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	response := service.ToUserModel(user)
	return c.JSON(response)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid dob format")
	}

	user, err := h.repo.Update(c.Context(), int32(id), req.Name, dob)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	if err := h.repo.Delete(c.Context(), int32(id)); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.repo.List(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch users")
	}

	response := make([]interface{}, 0, len(users))
	for _, u := range users {
		response = append(response, service.ToUserModel(u))
	}

	return c.JSON(response)
}
