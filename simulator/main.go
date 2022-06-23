package main

import (
	"fmt"

	"github.com/kaiqueluz/imersapfsfc2-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[1])
}
