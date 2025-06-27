package handlers

import (
    "github.com/labstack/echo/v4"

    "keycurl/utils"
)

var db = utils.DB()
var env = utils.GetEnv()

type jsonMessage struct {
    Key      string    `json:"key"`
    Value    string    `json:"value"`
}

func errorJSON(key string, value string) jsonMessage {
    return jsonMessage{Key: key, Value: value}
}

func HealthCheck(c echo.Context) error {
    return c.String(220, "keycurl:~")
}
