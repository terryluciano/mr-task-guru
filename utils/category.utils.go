package utils

import (
	"fmt"
	"strings"
)

var storedCategories = make(map[string]string)

func AddCategory(category string) (map[string]string, error) {

	baseCategory := strings.ToLower(category)

	for _, value := range storedCategories {
		if value == baseCategory {
			return nil, fmt.Errorf("%s Already Exists as a Category", TitleizeString(category))
		}
	}

	if categoryID, err := GenerateRandomID(); err != nil {
		return nil, err
	} else {

		storedCategories[categoryID] = baseCategory

		return map[string]string{categoryID: baseCategory}, nil
	}
}

// TODO: when user removes a category, remove all instances of that category from tasks
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

func GetAllCategories() map[string]string {
	return storedCategories
}

func DoesCategoryExists(category string) bool {
	if _, ok := storedCategories[category]; !ok {
		return false
	} else {
		return true
	}
}
