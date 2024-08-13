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
            "name": "Jon Snow",
            "email": "jon@labstack.com"
        }
    `
    mockRes := `
        $  status: 200
        <br /><br />
        <textarea id="response-textarea" readonly>&#013;</textarea>
    `

    e := routes.ConfigGlobalRoutes()
    e = routes.TemplateRoutes()
    e = routes.UserRoutes()
    e = routes.RequestRoutes()

    req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
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
            "user_email": ""
            "url"
            "method"
            "origin"
            "headers"
            "body"
            "status"
            "date"
            "hidden"
        }
    `
    mockRes := `
        $  status: 200
        <br /><br />
        <textarea id="response-textarea" readonly>&#013;</textarea>
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
