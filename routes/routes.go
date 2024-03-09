package routes

import (
	gql "root/graphql"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

type Input struct {
	Query         string                 `query:"query"`
	OperationName string                 `query:"operationName"`
	Variables     map[string]interface{} `query:"variables"`
}

var schema, _ = gql.CreateGraphQLSchema()

func Hello(ctx *fiber.Ctx) error {
	var input Input
        if err := ctx.BodyParser(&input); err != nil {
            return ctx.
                Status(fiber.StatusInternalServerError).
                SendString("Cannot parse body: " + err.Error())
        }

        result := graphql.Do(graphql.Params{
            Schema:         schema,
            RequestString:  input.Query,
            OperationName:  input.OperationName,
            VariableValues: input.Variables,
        })

        ctx.Set("Content-Type", "application/graphql-response+json")
        return ctx.JSON(result)
}

func Create(ctx *fiber.Ctx) error {
	var input Input
        if err := ctx.BodyParser(&input); err != nil {
            return ctx.
                Status(fiber.StatusInternalServerError).
                SendString("Cannot parse body: " + err.Error())
        }
    
        result := graphql.Do(graphql.Params{
            Schema:         schema,
            RequestString:  input.Query,
            OperationName:  input.OperationName,
            VariableValues: input.Variables,
        })
    
        ctx.Set("Content-Type", "application/graphql-response+json")
        return ctx.JSON(result)
}