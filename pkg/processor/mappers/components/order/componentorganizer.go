// ordering package is a group of functions associated with component ordering,
// This package seeks to avoid code duplication between ordered components
package order

import (
	"fmt"
)

// SetPageOrder is responsible for validating the component order and adding the order to the page
func SetPageOrder(template *map[string]interface{}, resourceName, sourceKey string) (int, error) {
	defer delete(*template, "order")

	order, ok := (*template)["order"]
	if !ok {
		return 0, fmt.Errorf("could not find field order on %s \"%s\"", resourceName, sourceKey)
	}
	validOrder, ok := order.(float64)
	if !ok || validOrder < 1 {
		return 0, fmt.Errorf("the order field passed on %s \"%s\" is not valid", resourceName, sourceKey)
	}

	return int(validOrder), nil
}
