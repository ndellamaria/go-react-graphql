package mutations 

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"


	"context"

	"app/data"
	types "app/types"
)

var DeleteTodo = &graphql.Field {
	Type: types.NotTodo,
	Description: "Delete a todo", 
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// get params
		id, _ :=  primitive.ObjectIDFromHex(params.Args["id"].(string))
		
		notTodoCollection := mongo.Client.Database("notTodoDB").Collection("Not_Todos")

		// new_id,_ := primitive.ObjectIDFromHex(id)

		_, err := notTodoCollection.DeleteOne(context.Background(), 
			map[string]primitive.ObjectID{ "_id":  id })

		if err != nil { panic(err) }
		return id, nil
	},
}