package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	config2 "backend/config"
	"backend/data/repo"
	"backend/http/handlers"
	"backend/services"
)

// InitRouter initializes the Gin router
// It also sets up the custom CORS configuration and the RecoverStack middleware
func InitRouter() *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	// If we panic, throw a 500 with the error stack
	router.Use(RecoverStack(os.Stderr))

	// Define custom CORS configuration
	config := cors.DefaultConfig()

	env, err := config2.LoadConfig()
	if err != nil {
		panic(err)
	}

	var (
		localAddress       = env.WebsiteAddress
		localRemoteAddress = env.WebsiteAddressRemote
		storyBookAddress   = env.StoryBookAddress
	)

	if port := env.WebsitePort; port != "" {
		localAddress += ":" + port
		localRemoteAddress += ":" + port
	}

	if port := env.StoryBookPort; port != "" {
		storyBookAddress += ":" + port
	}

	fmt.Print(localAddress)

	config.AllowOrigins = []string{localAddress, storyBookAddress, localRemoteAddress}
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// Apply custom CORS configuration to all routes
	router.Use(cors.New(config))

	// Apply rate limiting to all routes
	if env.RateLimiterRequestsPerSecond != "" {
		// convert string to int
		rps, err := strconv.Atoi(env.RateLimiterRequestsPerSecond)
		if err != nil {
			panic(err)
		}

		router.Use(RateLimitMiddleware(rps))
	}

	return router
}

// InitRoutes initializes the Gin routes
// It also sets up the AuthMiddleware that checks that each request has a valid Firebase token
func InitRoutes() *gin.Engine {
	router := InitRouter()

	router.GET("/healthz", HealthZ)

	router.Use(LoggerMiddleware())

	var (
		userRepo     = repo.NewUserRepo()
		skillRepo    = repo.NewSkillRepo()
		industryRepo = repo.NewIndustryRepo()
		// reviewRepo   = repo.NewReviewRepo()
	)

	var (
		userService     = services.NewUserService(userRepo)
		skillService    = services.NewSkillService(skillRepo)
		industryService = services.NewIndustryService(industryRepo)
		// reviewService   = services.New
	)

	h := handlers.Handler{
		UserService:     userService,
		SkillsService:   skillService,
		IndustryService: industryService,
	}

	v1 := router.Group("/v1")
	setV1Routes(v1, h)

	userRoutes := v1.Group("/user")
	userRoutes.Use(AuthMiddleware())
	setV1AuthenticatedUserRoutes(userRoutes, h)

	workerRoutes := v1.Group("/worker")
	workerRoutes.Use(AuthMiddleware())
	setV1AuthenticatedWorkerRoutes(workerRoutes, h)

	businessRoutes := v1.Group("/business")
	businessRoutes.Use(AuthMiddleware())
	setV1AuthenticatedBusinessRoutes(businessRoutes, h)

	jobListingRoutes := v1.Group("/joblistings")
	jobListingRoutes.Use(AuthMiddleware())
	setV1AuthenticatedJobListingRoutes(jobListingRoutes, h)

	return router
}

func setV1Routes(v1 *gin.RouterGroup, h handlers.Handler) {
	v1.GET("/skills", h.HandleGetSkills)
	v1.GET("/industries", h.HandleGetIndustries)
}

/* /v1/user/ */
func setV1AuthenticatedUserRoutes(v1User *gin.RouterGroup, h handlers.Handler) {
	v1User.POST("", h.HandleCreateUser)
	v1User.GET("", h.HandleGetLoggedInUser)
	v1User.POST("/firebaseSignUp", h.HandleFirebaseSignUp)
	v1User.GET("/firebaseSignIn", h.HandleFirebaseSignIn)
}

/* /v1/worker/ */
func setV1AuthenticatedWorkerRoutes(v1Worker *gin.RouterGroup, h handlers.Handler) {

}

/* /v1/business/ */
func setV1AuthenticatedBusinessRoutes(v1Business *gin.RouterGroup, h handlers.Handler) {

}

/* /v1/joblistings/ */
func setV1AuthenticatedJobListingRoutes(v1JobListings *gin.RouterGroup, h handlers.Handler) {

}

func HealthZ(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Server running")
}
