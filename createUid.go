package main

import (
	"fmt"

	"github.com/google/uuid"
)

func createRandomUid() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating UUID: ", err)
	}
	return id.String()
}
