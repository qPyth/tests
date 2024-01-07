package transport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"tests/internal/models"
)

// EntTest godoc
// @Summary      getting ENT test
// @Description  get test by language and profiles subjects`s ids
// @Tags         ENT test
// @Accept       json
// @Produce      json
// @Param        lang	path	string  true  "Language of test: ru or kz"
// @Param        profile1	path	int  true  "id of first profile subject"
// @Param        profile2	path	int  true  "id of second profile subject"
// @Success      200  {object}  services.EntTestOutput
// @Failure      400  {object}  transport.ErrorResponse
// @Router       /ent-test [get]
func (h *Handler) EntTest(c *gin.Context) {
	{
		var mathLit int = 218
		var kazHistory int = 220
		var readingLit int = 187

		lang := c.Query("lang")
		profileSub1, err := strconv.Atoi(c.Query("profile1"))
		profileSub2, err := strconv.Atoi(c.Query("profile2"))
		if err != nil {
			c.JSON(400, ErrorResponse{ErrorCode: 400, Message: http.StatusText(400)})
			return
		}
		if lang == "kz" {
			mathLit = 219
			kazHistory = 221
			readingLit = 186
		}

		test, err := h.services.GetTest(mathLit, kazHistory, readingLit, profileSub1, profileSub2)
		if err != nil {
			if errors.Is(err, models.ErrQuestionNotFound) {
				log.Println(err)
				c.JSON(400, ErrorResponse{ErrorCode: 400, Message: http.StatusText(400)})
				return
			}
			log.Println(err)
			c.JSON(500, ErrorResponse{ErrorCode: 500, Message: http.StatusText(500)})
			return
		}
		c.JSON(200, test)
	}
}
