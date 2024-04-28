package http

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

// RecoverStack
/* This is a Gin middleware that recovers from panics and writes the stack trace to the given io.Writer.*/
func RecoverStack(out io.Writer) gin.HandlerFunc {
	var logger *log.Logger
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if cantWrite(err) {
					if logger != nil {
						logger.Printf("Can't Write\n%s", err)
					}
					c.Error(err.(error))
					c.Abort()
				} else {
					if logger != nil {
						logger.Printf("Panic Recovered\n%s\n%s", err, string(debug.Stack()))
					}
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

func cantWrite(err interface{}) bool {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				return true
			}
		}
	}

	// if we recive abort handler from a reverse proxy.
	if err != nil && err == http.ErrAbortHandler {
		return true
	}

	return false
}
