package service

import (
	"golang_test_task/model"
	"math/rand"
	"time"
)

var statuses = [4]string{"done", "in_process", "error", "new"}

func Get(id string) (*model.Task, error) {
	// Implementing connection to gRPC, HTTP or any other microservice..
	m := model.Task{
		ID:     id,
		Status: randomStatus(),
	}

	return &m, nil
}

func randomStatus() string {
	// Just randomizer
	rand.Seed(time.Now().UnixNano())
	randomNumber := 0 + rand.Intn(4-0)

	return statuses[randomNumber]
}
