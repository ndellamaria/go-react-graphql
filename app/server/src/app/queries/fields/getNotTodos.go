package queries

import (
	"context"
	"github.com/graphql-go/graphql"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"app/data"
	types "app/types"
)

type todoStruct struct {
	ID primitive.ObjectID `bson:"_id" json:id,omitempty"`
	NAME string `json:"name"`
	DESCRIPTION string `json:"description"`
	COMPLETE bool `json:"compelte"`
}

var GetNotTodos = &graphql.Field {
	Type:	graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		notTodoCollection := mongo.Client.Database("notTodoDB").Collection("Not_Todos")

		todos, err := notTodoCollection.Find(context.Background(), bson.M{})
		if err != nil { panic(err) }

		var todosList []todoStruct

		for todos.Next(context.Background()) {

			var elm todoStruct

			err := todos.Decode(&elm)
			if err != nil { panic(err) }

			todosList = append(todosList, elm)
		}

		if err := todos.Err(); err != nil {
			panic(err)
		}

		// Close the cursor once finished
		todos.Close(context.TODO())

		return todosList, nil
	},
}