package routes

import (
	"net/http"

	apiContext "github.com/globocom/huskyCI/api/context"
	"github.com/labstack/echo"
)

//GetAPIVersion returns the API version
func GetAPIVersion(c echo.Context) error {
	configAPI := apiContext.APIConfiguration
	requestResult := map[string]string{
		"version": configAPI.Version,
		"date":    configAPI.ReleaseDate,
	}
	return c.JSON(http.StatusOK, requestResult)
}
