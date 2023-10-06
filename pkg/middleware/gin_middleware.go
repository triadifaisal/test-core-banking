package middleware

import (
	"bytes"
	"encoding/json"

	"corebanking/pkg/helper/logger"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// GinBodyErrorLogger ...
func GinBodyErrorLogger(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	url := "-"
	if u := c.Request.URL; u != nil {
		url = u.String()
	}
	c.Writer = blw
	// before process
	c.Next()
	// after process
	statusCode := c.Writer.Status()

	headers := ""
	if c.Request.Header != nil {
		marshalled, _ := json.Marshal(c.Request.Header)
		headers = string(marshalled)
	}

	if statusCode >= 400 && statusCode < 500 {
		logger.Log.Warnf("Responded with 4xx. Code: %v, URL: %v, Request Headers: %v, Response: %v", statusCode, url, headers, blw.body.String())
	}

	if statusCode >= 500 {
		logger.Log.Warnf("Responded with 5xx. Code: %v, URL: %v, Request Headers: %v, Response: %v", statusCode, url, headers, blw.body.String())
	}
}
