package main

import (
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.JSON(200, Response{
		Message: "hello world",
	})
	fmt.Print(recorder.Body.String())
}