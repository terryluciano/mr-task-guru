package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type Status int

const (
	Inactive Status = iota
	Active
	Complete
	Incomplete
)

type Task struct {
	ID       string
	Name     string `json:"name"`
	Category string `json:"category"`
	Minutes  int    `json:"minutes"`
	Status   Status
	Date     time.Time
}

var storedTasks = make(map[string]Task)

func statusToString(s Status) string {
	switch s {
	case Inactive:
		return "Inactive"
	case Active:
		return "Active"
	case Complete:
		return "Complete"
	case Incomplete:
		return "Incomplete"
	default:
		return "Unknown"
	}
}

func formatTask(task Task) map[string]interface{} {
	return map[string]interface{}{
		"id":       task.ID,
		"name":     task.Name,
		"category": task.Category,
		"minutes":  task.Minutes,
		"status":   statusToString(task.Status),
		"date":     task.Date.Format("2006-01-02 15:04:05"),
	}
}

func AddTask(task Task) error {

	currentTime := time.Now().UnixNano()
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {

		return fmt.Errorf("could not generate task id: %w", err)
	}

	taskId := strconv.Itoa(int(currentTime)) + "-" + randomNumber.String()

	task.ID = taskId
	task.Status = Inactive
	task.Date = time.Now()

	storedTasks[taskId] = task

	fmt.Println(storedTasks)

	return nil
}

func RemoveTask() {

}

func GetTask(id string) map[string]interface{} {
	if _, ok := storedTasks[id]; !ok {
		return nil
	}
	return formatTask(storedTasks[id])
}

func GetAllTask() map[string]interface{} {
	formattedTasks := make(map[string]interface{})

	for _, task := range storedTasks {
		formattedTasks[task.ID] = formatTask(task)
	}

	return formattedTasks
}

func UpdateTask() {

}
