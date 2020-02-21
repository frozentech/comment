package model

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

// GenerateUUID ...
func GenerateUUID(args ...string) string {
	var (
		token string
		id    = uuid.UUID{}
	)

	if len(args) > 0 {

		for _, value := range args {
			token = token + value
		}

		id = uuid.NewV3(id, token)
		return id.String()
	}

	id = uuid.Must(uuid.NewV4())
	return id.String()
}

// JSON ...
func JSON(model interface{}) string {
	body, _ := json.Marshal(model)
	return string(body)
}
