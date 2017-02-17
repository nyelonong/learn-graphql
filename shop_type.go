package main

import "github.com/graphql-go/graphql"

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shop",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if shop, ok := p.Source.(Shop); ok {
					return shop.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if shop, ok := p.Source.(Shop); ok {
					return shop.Name, nil
				}
				return nil, nil
			},
		},
		"product": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if shop, ok := p.Source.(Shop); ok {
					var prodList []Product
					for _, prod := range shop.Products {
						prodList = append(prodList, ProductData[prod])
					}
					return prodList, nil
				}
				return nil, nil
			},
		},
	},
})
