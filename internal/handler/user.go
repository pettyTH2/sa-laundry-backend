package handler

import (
	"laundry-backend/internal/entity"
	"laundry-backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := h.userUsecase.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserByRole(c *fiber.Ctx) error {
    role := c.Params("role")
    users, err := h.userUsecase.GetByRole(role)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(users)
}

func (h *UserHandler) GetUserByPhone(c *fiber.Ctx) error {
    phoneNumber := c.Params("phone_number")
    user, err := h.userUsecase.GetByPhoneNumber(phoneNumber)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    if user == nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
    }
    return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	user.ID = uint(id)
	if err := h.userUsecase.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}
