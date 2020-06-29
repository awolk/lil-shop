package model

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(uuid uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(uuid.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	s, err := graphql.UnmarshalString(v)
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.Parse(s)
}
