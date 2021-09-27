package main

import (
    "fmt"
    "github.com/go-playground/validator"
    "github.com/jiramot/go-oauth2-adapter/pkg/restclient"
    "github.com/jiramot/go-oauth2-adapter/pkg/util"
    "github.com/labstack/echo/v4"
    "net/http"
    "os"
)

func main() {
    adminEndpoint := os.Getenv("OAUTH2_ADMIN_ENDPOINT")
    if adminEndpoint == "" {
        adminEndpoint = "http://localhost:8081"
    }
    e := echo.New()
    e.Validator = &Validator{validator: validator.New()}
    e.POST("/token", func(c echo.Context) error {
        tokenRequest := new(TokenRequest)
        if err := util.BindAndValidateRequest(c, tokenRequest); err != nil {
            return c.String(http.StatusBadRequest, "Bad request")
        }
        tokenResponse := new(TokenResponse)
        err := restclient.PostJson(fmt.Sprintf("%s/oauth2/auth/token", adminEndpoint), tokenRequest, tokenResponse)
        if err != nil {
            return c.String(http.StatusBadRequest, "Bad request")
        }

        return c.JSON(http.StatusOK, tokenResponse)
    })
    e.Logger.Fatal(e.Start(":9000"))
}

type TokenRequest struct {
    Cif      string `json:"cif"`
    Amr      string `json:"amr"`
    ClientId string `json:"client_id"`
}

type TokenResponse struct {
    AccessToken string `json:"access_token"`
    ExpireAt    int64  `json:"expires_at"`
    TokenType   string `json:"token_type"`
}

type Validator struct {
    validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
