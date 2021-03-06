package middleware

import (
	"net/http"

	"github.com/smartystreets/httpx"
)

type FileHandler struct {
	filename    string
	contentType string
}

func NewFileHandler(fullPathToFile, contentType string) *NestableHandler {
	return NewNestableHandler(&FileHandler{filename: fullPathToFile, contentType: contentType})
}

func (this *FileHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if len(this.contentType) > 0 {
		response.Header().Set(httpx.HeaderContentType, this.contentType)
	}

	http.ServeFile(response, request, this.filename)
}
