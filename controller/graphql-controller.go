package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deepakworldphp86/golang-api/graph"
)

// graphql get
func GetGraphQl(c *gin.Context) {
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

		http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)
		c.JSON(200, gin.H{
			"message": "Hello, World Get",
		})
}

// graphql post
func PostGraphQl(c *gin.Context) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
		c.JSON(200, gin.H{
			"message": "Hello, World Post",
		})
}
