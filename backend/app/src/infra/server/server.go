package server

import (
	"app/envvars"
	"app/infra/graph/generated"
	"app/infra/graph/resolver"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func App() {
	var PORT = envvars.ServerEnv.PORT

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err := e.Start(":" + PORT)
	if err != nil {
		log.Fatalln(err)
	}
}
