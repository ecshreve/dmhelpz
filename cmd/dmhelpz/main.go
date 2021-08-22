package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
	"github.com/ecshreve/dmhelpz/pkg/gql/gen"
	"github.com/kr/pretty"
)

func main() {
	ctx := context.Background()

	client := &gen.Client{
		Client: client.NewClient(http.DefaultClient, "https://www.dnd5eapi.co/graphql"),
	}

	m, err := client.GetMonster(ctx, "Aboleth")
	if err != nil {
		log.Fatal(err)
	}
	pretty.Print(m)
}
