package handlers

import (
	"fmt"
	"os"
	"serve-ressources/requests"
	"serve-ressources/responses"
	s "serve-ressources/server"
	"serve-ressources/services"
	"time"

	"net/http"

	"github.com/h2non/bimg"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type RessourceHandler struct {
	server *s.Server
}

func NewRessourceHandler(server *s.Server) *RessourceHandler {
	return &RessourceHandler{
		server: server,
	}
}

func cacher(redis *services.RedisService, key string, buffer []byte) {
	duration := time.Duration(10) * time.Second
	err := redis.Client.Set(redis.Ctx, key, buffer, duration).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Cached", key, "for", duration)
}

// GetRessource godoc
// @Summary Get ressource
// @Description Get ressource
// @Tags CDN
// @Accept json
// @Produce json
// @Param id path string true "Ressource ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /cdn/{id} [get]
func (h *RessourceHandler) GetRessource(c echo.Context) error {
	id := c.Param("id")

	ressourceRequest := new(requests.RessourceBimg)
	key := fmt.Sprintf("%s-%d-%d-%d-%t-%t", id, ressourceRequest.W, ressourceRequest.H, ressourceRequest.Q, ressourceRequest.C, ressourceRequest.A)

	val, err := h.server.REDIS.Client.Get(h.server.REDIS.Ctx, key).Result()

	fmt.Println(err)

	if err != redis.Nil && err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	if val != "" {
		res := c.Response()

		res.Header().Set("Content-Disposition", "attachment; filename="+id)
		res.Header().Set("Content-Type", http.DetectContentType([]byte(val)))
		res.Header().Set("Content-Length", fmt.Sprintf("%d", len(val)))
		res.Header().Set("X-Cache", "HIT")

		res.Write([]byte(val))

		return nil
	}

	if err := c.Bind(ressourceRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if err := ressourceRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	file, err := os.Open(fmt.Sprintf("%s/%s", h.server.Config.CDN.UploadPath, id))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	fileSize := fileInfo.Size()

	buffer := make([]byte, fileSize)
	file.Read(buffer)

	contentType := http.DetectContentType(buffer)

	var newImage []byte
	if contentType == "image/gif" && ressourceRequest.A {
		// TODO: Gifs are not working
		newImage = buffer
	} else {
		newImage, err = bimg.NewImage(buffer).Process(bimg.Options{
			Width:  	ressourceRequest.W,
			Height: 	ressourceRequest.H,
			Quality:	ressourceRequest.Q,
			Crop:   	ressourceRequest.C,
		})
		if err != nil {
			return responses.ErrorResponse(c, http.StatusInternalServerError, err)
		}
	}

	res := c.Response()
	res.Header().Set("Content-Disposition", "attachment; filename="+id)
	res.Header().Set("Content-Type", contentType)
	res.Header().Set("Content-Length", fmt.Sprintf("%d", len(newImage)))
	res.Header().Set("X-Cache", "MISS")

	cacher(h.server.REDIS, key, newImage)

	res.Write(newImage)

	return nil
}
