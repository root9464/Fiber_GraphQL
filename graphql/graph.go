package graphql

import (
	"log"
	"root/database"
	"root/database/models"

	"github.com/graphql-go/graphql"
)

func CreateGraphQLSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "world", nil
            },
        },
    }
	//любая залупа меняющая данные (create/update/delete)
    mutationFields := graphql.Fields{
        "createUser": &graphql.Field{
            Type: graphql.String,
            Args: graphql.FieldConfigArgument{
                "name": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
                "password": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                name := p.Args["name"].(string)
                password := p.Args["password"].(string)
                database.Db.DB.Create(&models.Data{Name: name, Password: password})
                return "User created", nil
            },
        },
    }

    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    mutationType := graphql.NewObject(graphql.ObjectConfig{
        Name:   "Mutation",
        Fields: mutationFields,
    })
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery), Mutation: mutationType}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

	return schema, nil
}