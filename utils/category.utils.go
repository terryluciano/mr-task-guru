package utils

import (
	"fmt"
	"strings"
)

var storedCategories = make(map[string]string)

func AddCategory(category string) error {

	baseCategory := strings.ToLower(category)

	for _, value := range storedCategories {
		if value == baseCategory {
			return fmt.Errorf("%s Already Exists as a Category", TitleizeString(category))
		}
	}

	if categoryID, err := GenerateRandomID(); err != nil {
		return err
	} else {

		storedCategories[categoryID] = baseCategory

		return nil
	}
}

func RemoveCategory(id string) error {

	if _, ok := storedCategories[id]; !ok {
		return fmt.Errorf("Category does not exists with ID: %s", id)
	}

	delete(storedCategories, id)

	return nil
}

func GetCategory(id string) (map[string]string, error) {
	if category, ok := storedCategories[id]; !ok {
		return nil, fmt.Errorf("Category does not exists with ID: %s", id)
	} else {
		return map[string]string{id: category}, nil
	}
}

// func GetAllCategories() map[string]interface{} {}
