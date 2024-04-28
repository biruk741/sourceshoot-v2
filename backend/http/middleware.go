package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"

	config2 "backend/config"
	"backend/firebase"

	"net/http"
	"strings"
	"time"
)

var config *config2.Config

// AuthMiddleware
/* This is a Gin middleware that checks for a valid Firebase token in the Authorization header.
 * If the token is valid, the Firebase UID is stored in the request context for later use.
 */
func AuthMiddleware() gin.HandlerFunc {
	var err error
	if config == nil {
		config, err = config2.LoadConfig()
		if err != nil {
			panic(err)
		}
	}

	return func(c *gin.Context) {
		if config.DisableAuthMiddleware == "true" && config.Environment == config2.ENVIRONMENT_LOCAL {
			c.Set("firebaseID", config.MockFirebaseID)
			c.Next()
			return
		}

		token := c.GetHeader("Authorization")

		if firebase.FirebaseAuth == nil {
			err := firebase.InitFirebase()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				c.Abort()
				return
			}
		}

		if splitAuth := strings.Split(token, "Bearer "); len(splitAuth) > 1 {
			token = splitAuth[1]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized or invalid token"})
			c.Abort()
			return
		}

		// Example of Firebase token verification (use Firebase Admin SDK):
		decodedToken, err := firebase.FirebaseAuth.VerifyIDToken(c, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		userID := decodedToken.UID

		// Store the userID in the request context for later use
		c.Set("firebaseID", userID)

		c.Next()
	}
}

// RequestResponseLogger is a custom middleware that logs incoming requests and responses.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process the request
		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		config, err := config2.LoadConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error loading config"})
			c.Abort()
			return
		}

		var loggerInstance *zap.Logger
		switch config.Environment {
		case "LOCAL":
			zapConfig := zap.NewDevelopmentConfig()
			zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			loggerInstance, err = zapConfig.Build()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating logger"})
				c.Abort()
				return
			}
		}

		// Log the request and response details
		loggerInstance.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status_code", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}

// RateLimitMiddleware creates a middleware that limits the number
// of requests allowed in a specific time window.
func RateLimitMiddleware(rps int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(rps), rps)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}

		c.Next()
	}
}
