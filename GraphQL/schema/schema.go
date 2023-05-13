package schema

import "github.com/graphql-go/graphql"

var (
	productType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Product",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.Float,
				},
			},
		},
	)

	queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"product": &graphql.Field{
					Type:        productType,
					Description: "Get a product by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						// Here you can implement the logic to fetch a product by ID from your microservice
						id, ok := params.Args["id"].(string)
						if !ok {
							return nil, nil
						}

						// Mocked response for demonstration purposes
						return map[string]interface{}{
							"id":    id,
							"name":  "Sample Product",
							"price": 9.99,
						}, nil
					},
				},
			},
		},
	)

	Schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
)
