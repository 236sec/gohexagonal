package adapters

import (
	"fmt"

	"github.com/236sec/hexagonal/service"
	"github.com/gofiber/fiber/v2"
)

type HttpCustomerHandler struct {
	service service.CustomerService
}

func NewHttpCustomerHandler(service service.CustomerService) *HttpCustomerHandler {
	return &HttpCustomerHandler{service: service}
}

func (h *HttpCustomerHandler) CreateUser(c *fiber.Ctx) error {
	var customer service.CustomerResponse
	if err := c.BodyParser(&customer); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateUser(customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(customer)
}
