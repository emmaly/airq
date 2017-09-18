package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dustywilson/airnow"
)

func main() {
	key := flag.String("key", "", "AirNow API Key")
	zip := flag.String("zip", "", "ZIP Code")
	radius := flag.Int("radius", 25, "Radius")
	variable := flag.String("var", "aqi", "Return: {aqi,color,category}")
	flag.Parse()

	if key == nil || *key == "" || zip == nil || *zip == "" || radius == nil || *radius < 1 {
		flag.Usage()
		os.Exit(1)
	}

	an := airnow.New(*key)
	ob, err := an.NowByZIP(*zip, *radius)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *variable {
	case "aqi":
		fmt.Println(ob.AQI)
	case "color":
		fmt.Printf("%d,%d,%d\n", ob.Category.Color.R, ob.Category.Color.G, ob.Category.Color.B)
	case "category":
		fmt.Println(ob.Category.Num)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
