package adapters

import (
	"fmt"

	"github.com/236sec/hexagonal/core"
	"github.com/gofiber/fiber/v2"
)

type HttpCustomerHandler struct {
	service core.CustomerService
}

func NewHttpCustomerHandler(service core.CustomerService) *HttpCustomerHandler {
	return &HttpCustomerHandler{service: service}
}

func (h *HttpCustomerHandler) CreateUser(c *fiber.Ctx) error {
	var customer core.CreateCustomerDto
	if err := c.BodyParser(&customer); err != nil {
    fmt.Println(err)
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
  }
	err := h.service.CreateUser(customer)
  if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
  }

  return c.Status(fiber.StatusCreated).JSON(customer)
}