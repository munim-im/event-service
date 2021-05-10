package utils

import (
	"event-service/pb"
	"fmt"
)

func ValidateEventInput(input *pb.EventInputParams) error {
	if input.Environment == "" {
		return fmt.Errorf("Environement should not be empty")
	}
	if input.Component == "" {
		return fmt.Errorf("Component should not be empty")
	}
	if input.Email == "" {
		return fmt.Errorf("Email should not be empty")
	}
	if input.MessageString == "" {
		return fmt.Errorf("Message string should not be empty")
	}
	if input.Data == nil {
		return fmt.Errorf("Data should not be empty")
	}
	_, err := input.Data.MarshalJSON()
	if err != nil {
		return err
	}
	return nil
}