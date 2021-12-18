package main_test

import (
	"fmt"
	"testing"
)

func TestMapTypes(t *testing.T) {

	foods := map[string]interface{}{
		"Estilo_musical": "Pagode",
		"Downloads":      10003,
		"Deezer":         "",
		"Midia": struct {
			tipo  string
			price float64
		}{"CD", 10.75},
		"Spotify": true,
	}

	for k, v := range foods {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		case int:
			fmt.Printf("Looks like item %q is a Inteiro, specifically %T\n", k, c)
		case bool:
			fmt.Printf("Looks like item %q is a Boleano, specifically %T\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}
}

func TestParseMap(t *testing.T) {

	data := make(map[string]interface{})
	data["person"] = map[string]interface{}{
		"peter": map[string]interface{}{
			"scores": map[string]interface{}{
				"calculus": 88,
				"algebra":  99,
				"golang":   89,
			},
		},
	}

	result := parseMap(data)

	fmt.Printf("Objeto transformado %q\n", result)
}

func parseMap(aMap map[string]interface{}) interface{} {
	var retVal interface{}

	for _, val := range aMap {
		switch c := val.(type) {
		case map[string]interface{}:
			retVal = parseMap(val.(map[string]interface{}))
			fmt.Printf("Looks like item is a number, specifically %q\n", c)
		//case []interface{}:
		//  retVal = parseArray(val.([]interface{}))
		default:
			//here i would have done aMap["physics"] = 95 if I could access the original map by reference, but that is not possible

			retVal = aMap

		}
	}

	return retVal
}
