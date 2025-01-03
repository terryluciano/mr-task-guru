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
	Name     string
	Category string
	Minutes  int
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
	task.Date = time.Now()

	storedTasks[taskId] = task

	fmt.Println(storedTasks)

	return nil
}

func RemoveTask(id string) error {
	if _, ok := storedTasks[id]; !ok {
		return fmt.Errorf("task with id %s not found", id)
	}

	delete(storedTasks, id)
	return nil
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

func UpdateTask(id string, name *string, minutes *int, category *string, status *Status) error {
	if _, ok := storedTasks[id]; !ok {
		return fmt.Errorf("task with id %s not found", id)
	}

	task := storedTasks[id]

	if status != nil {
		if *status < 0 || *status > 3 {
			return fmt.Errorf("invalid status")
		}
		task.Status = *status
	}

	if name != nil {
		task.Name = *name
	}

	if minutes != nil {
		task.Minutes = *minutes
	}

	if category != nil {
		task.Category = *category
	}
	storedTasks[id] = task
	return nil
}
