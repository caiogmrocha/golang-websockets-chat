package utils

import (
	"errors"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ExtractObjectID(objectID primitive.ObjectID) (string, error) {
	regex := regexp.MustCompile(`ObjectID\(\"(.*)\"\)`)

	matches := regex.FindStringSubmatch(objectID.String())

	if len(matches) < 2 {
		return "", errors.New("ERROR_GENERATING_TOKEN")
	}

	userID := matches[1]

	return userID, nil
}
