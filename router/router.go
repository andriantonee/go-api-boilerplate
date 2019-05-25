package router

import (
	"go-api-base/app/user"

	"github.com/gin-gonic/gin"
)

// Config is a container for config (E.g: env) and all service
type Config struct {
	UserService user.Service
}

// New is a constructor to wrap all handler and it's usage
func New(config *Config) *gin.Engine {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 	)
	// }))
	// router.Use(gin.Recovery())

	userHandlerInstance := userHandler{
		userService: config.UserService,
	}

	api := router.Group("/api")
	{
		api.POST("/user", userHandlerInstance.register)
		api.POST("/auth/login", userHandlerInstance.login)
	}

	return router
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

func newErrorResponse(code int, err string) errorResponse {
	return errorResponse{
		Code:  code,
		Error: err,
	}
}

// Custom Validator Message
// func badRequestValidator(
// 	ginContext *gin.Context,
// 	request interface{},
// 	validationErrors *validator.ValidationErrors,
// ) {
// 	invalidFields := make([]map[string]string, 0)

// 	for _, e := range *validationErrors {
// 		field, _ := reflect.TypeOf(request).FieldByName(e.Name)
// 		jsonFieldName := field.Tag.Get("json")
// 		if jsonFieldName == "" {
// 			jsonFieldName = e.Name
// 		}

// 		errors := map[string]string{}
// 		errors[jsonFieldName] = e.Tag

// 		invalidFields = append(invalidFields, errors)
// 	}

// 	ginContext.JSON(http.StatusOK, invalidFields)
// }