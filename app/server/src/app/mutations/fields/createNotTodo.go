package mutations 

import (
	"github.com/graphql-go/graphql"
	"context"

	"app/data"
	types "app/types"
)

type todoStruct struct {
	NAME string `json:"name"`
	DESCRIPTION string 	`json:"description"`
	COMPLETE bool `json:"complete"`
}

var CreateNotTodo = &graphql.Field {
	Type: types.NotTodo,
	Description: "Create a not Todo", 
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// get params
		name, _ := params.Args["name"].(string)
		description, _ := params.Args["description"].(string)
		complete := false
		notTodoCollection := mongo.Client.Database("notTodoDB").Collection("Not_Todos")

		_, err := notTodoCollection.InsertOne(context.Background(), 
			map[string]interface{}{ "name":name, "description": description, "complete": complete })

		if err != nil { panic(err) }

		return todoStruct{ name, description, complete}, nil
	},
}