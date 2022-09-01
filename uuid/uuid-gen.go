package uuid

import "github.com/google/uuid"

func GenerateUUID() uuid.UUID {
	id, err := uuid.NewUUID()

	if err != nil {
		return uuid.UUID{}
	}

	return id
}

func ConvertToUUID(id string) uuid.UUID {
	ID, err := uuid.Parse(id)

	if err != nil {
		return uuid.UUID{}
	}

	return ID
}



