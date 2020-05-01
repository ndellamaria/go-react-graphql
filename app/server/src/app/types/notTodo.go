package types

import (
	"github.com/graphql-go/graphql"
)

var NotTodo = graphql.NewObject(graphql.ObjectConfig {
	Name:"NotTodo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		}, 
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"complete": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

