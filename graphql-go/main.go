package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Echo struct {
	Text string `json:"text"`
}

var echoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Echo",
	Fields: graphql.Fields{
		"text": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"echo": &graphql.Field{
			Type:        echoType,
			Description: "Get echo of the argument",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				text, _ := params.Args["text"].(string)

				echo := Echo{
					Text: text,
				}

				return echo, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
})

func main() {

	h := handler.New(&handler.Config{
		Schema: &schema,
	})

	http.Handle("/graphql", h)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello World")
	})

	// Display some basic instructions
	fmt.Println("Now server is running on port 8080")

	http.ListenAndServe(":8080", nil)
}
