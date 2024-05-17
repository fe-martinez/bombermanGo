package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func CreateRandomUid() string {
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating UUID: ", err)
	}
	return id.String()
}
