package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deepakworldphp86/golang-api/graph"
	"github.com/99designs/gqlgen/graphql/playground"

)

// graphql get
func GetGraphQl(c *gin.Context) {
	    graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
        graphqlHandler.ServeHTTP(c.Writer, c.Request) 
}

// graphql post
func PostGraphQl(c *gin.Context) {
	graphqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	graphqlHandler.ServeHTTP(c.Writer, c.Request) 
}

// graphql post
func GetGraphQlPlayground(c *gin.Context) {
	playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(c.Writer, c.Request)
}
