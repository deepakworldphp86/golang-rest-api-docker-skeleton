package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/deepakworldphp86/golang-api/cache"
	"github.com/deepakworldphp86/golang-api/config"
	"github.com/deepakworldphp86/golang-api/controller"
	"github.com/deepakworldphp86/golang-api/middleware"
	"github.com/deepakworldphp86/golang-api/repository"
	"github.com/deepakworldphp86/golang-api/service"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	dbpgsql           *gorm.DB                     = config.SetupDatabaseConnectionPgsql()
	
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	bookRepository    repository.BookRepository    = repository.NewBookRepository(db)
	receiptRepository repository.ReceiptRepository = repository.NewReceiptRepository(db)
	categoryRepository repository.CategoryRepository = repository.NewCategoryRepository(dbpgsql)

	rd           *redis.Client      = config.SetupRedisConnection()
	receiptCache cache.ReceiptCache = cache.NewReceiptCache(rd, 1000) // random expiration lmao

	jwtService     service.JWTService     = service.NewJWTService()
	authService    service.AuthService    = service.NewAuthService(userRepository)
	userService    service.UserService    = service.NewUserService(userRepository)
	bookService    service.BookService    = service.NewBookService(bookRepository)
	receiptService service.ReceiptService = service.NewReceiptService(receiptRepository)
	//categoryService    service.CategoryService = service.NewCategoryService(categoryRepository)


	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	bookController    controller.BookController    = controller.NewBookController(bookService, jwtService)
	receiptController controller.ReceiptController = controller.NewReceiptController(receiptService, receiptCache)
	graphqlController controller.GraphqlController = controller.NewGraphqlController(jwtService)


)

func Init() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "PUT", "PATCH", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("login", authController.Login)
		authRoutes.POST("register", authController.Register)
	}

	userRoutes := r.Group("/api/users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("profile", userController.Get)
		userRoutes.PUT("", userController.Update)
	}

	bookRoutes := r.Group("/api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("", bookController.All)
		bookRoutes.POST("", bookController.Insert)
		bookRoutes.GET("/:id", bookController.Get)
		bookRoutes.PUT("", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	receiptRoutes := r.Group("/api/receipts")
	{
		receiptRoutes.GET("/all", receiptController.All)
		receiptRoutes.POST("", receiptController.Insert)
		receiptRoutes.GET("/:id", receiptController.Show)
		receiptRoutes.PUT("", receiptController.Update)
		receiptRoutes.DELETE("/:id", receiptController.Delete)
	}

	graphqlRoutes := r.Group("/graphql", middleware.AuthorizeJWT(jwtService))
	{
		graphqlRoutes.GET("", graphqlController.GetGraphQl)
		graphqlRoutes.POST("", graphqlController.PostGraphQl)
	}

	graphqlPlaygroundRoutes := r.Group("/playground", middleware.AuthorizeJWT(jwtService))
	{
		graphqlPlaygroundRoutes.GET("", graphqlController.GetGraphQlPlayground)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})

	r.Run()
}
