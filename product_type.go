package main

import "github.com/graphql-go/graphql"

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if product, ok := p.Source.(Product); ok {
					return product.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if product, ok := p.Source.(Product); ok {
					return product.Name, nil
				}
				return nil, nil
			},
		},
		"price": &graphql.Field{
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if product, ok := p.Source.(Product); ok {
					return product.Price, nil
				}
				return nil, nil
			},
		},
	},
})
