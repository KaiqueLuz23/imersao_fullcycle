package main

import (
	"fmt"

	"github.com/kaiqueluz/imersapfsfc2-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "2",
		ClientID: "2",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[1])
}
