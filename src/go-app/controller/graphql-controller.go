package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deepakworldphp86/golang-api/graph"
	"github.com/99designs/gqlgen/graphql/playground"
    "github.com/deepakworldphp86/golang-api/service"
	//"github.com/dgrijalva/jwt-go"

)

type GraphqlController interface {
	GetGraphQl(context *gin.Context)
	PostGraphQl(context *gin.Context)
	GetGraphQlPlayground(context *gin.Context)
}

type graphqlController struct {
	jwtService  service.JWTService
}

func NewGraphqlController(jwtService service.JWTService) GraphqlController {
	return &graphqlController{
		jwtService:  jwtService,
	}
}

// graphql get
func (g *graphqlController) GetGraphQl(context *gin.Context) {
	    graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
        graphqlHandler.ServeHTTP(context.Writer, context.Request) 
}

// graphql post
func (g *graphqlController) PostGraphQl(context *gin.Context) {
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	graphqlHandler.ServeHTTP(context.Writer, context.Request) 
}

// graphql post
func (g *graphqlController) GetGraphQlPlayground(context *gin.Context) {
	playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(context.Writer, context.Request)
}
