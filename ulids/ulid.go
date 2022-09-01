package ulids

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

func ConvertToString(id uuid.UUID) (string, error) {
	theId, err := uuid.Parse(id.String())
	if err != nil {
		return "", err
	}

	return theId.String(), nil
}

