package tests

import (
    "postwoman/handlers"
    "postwoman/routes"

    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestGoodRequest(t *testing.T) {
    body := `
        {
            "user_email": "null"
            "url": "http://localhost:1323/healthcheck"
            "method": "GET"
            "origin": ""
            "headers": ""
            "body": ""
            "status": "200"
            "date": "1723697427081"
            "hidden": "false"
        }
    `
    mockRes := `
        $  status: 200
        <br /><br />
        <textarea id="response-textarea" readonly>postwoman:~&#013;</textarea>
    `

    e := routes.ConfigGlobalRoutes()
    e = routes.TemplateRoutes()
    e = routes.UserRoutes()
    e = routes.RequestRoutes()

    req := httptest.NewRequest(http.MethodPost, "/curl/request", strings.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    context := e.NewContext(req, rec)

    if assert.NoError(t, handlers.ExecuteCurlRequest(context)) {
        assert.Equal(t, mockRes, rec.Body.String())
    }
}

func TestBadRequest(t *testing.T) {
    body := `
        {
            "user_email": "null"
            "url": "http://localhost:1323/healthcheck"
            "method": "GET"
            "origin": ""
            "headers": ""
            "body": ""
            "status": "200"
            "date": "1723697427081"
            "hidden": "false"
        }
    `
    mockRes := `
        $  status: 200
        <br /><br />
        <textarea id="response-textarea" readonly>incorrect res&#013;</textarea>
    `

    e := routes.ConfigGlobalRoutes()
    e = routes.TemplateRoutes()
    e = routes.UserRoutes()
    e = routes.RequestRoutes()

    req := httptest.NewRequest(http.MethodPost, "/curl/request", strings.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    context := e.NewContext(req, rec)

    if assert.Error(t, handlers.ExecuteCurlRequest(context)) {
        assert.Equal(t, mockRes, rec.Body.String())
    }
}
