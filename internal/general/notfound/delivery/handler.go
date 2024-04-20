package delivery

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NotFoundHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewNotFoundHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) NotFoundHandler {
	return NotFoundHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (h NotFoundHandler) GetNotFound(c *fiber.Ctx) error {

	errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

	return c.Status(404).Render("error", fiber.Map{
		"Title":   h.Conf.App.Name,
		"Message": errorMessage,
	})
}
