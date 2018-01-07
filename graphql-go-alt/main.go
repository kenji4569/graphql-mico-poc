package main

import (
	"fmt"
	"net/http"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

var Schema = `
	schema {
		query: Query
	}
	type Query {
		echo(text: String!): Echo
	}
	type Echo {
		text: String!
	}
`

type echo struct {
	Text string
}

type Resolver struct{}

type echoResolver struct {
	e *echo
}

func (r *echoResolver) Text() string {
	return r.e.Text
}

func (r *Resolver) Echo(args struct{ Text string }) *echoResolver {
	e := echo{
		Text: args.Text,
	}
	return &echoResolver{&e}
}

var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(Schema, &Resolver{})
}

func main() {
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello World")
	})

	fmt.Println("Now server is running on port 8080")

	http.ListenAndServe(":8080", nil)
}
