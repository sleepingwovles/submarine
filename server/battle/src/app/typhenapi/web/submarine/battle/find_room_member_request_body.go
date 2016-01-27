// This file was generated by typhen-api

package battle

import (
	"app/typhenapi/core"
	"errors"
	"fmt"
	"net/url"
)

var _ = errors.New

type FindRoomMemberRequestBody struct {
	RoomKey string `codec:"room_key"`
}

// Coerce the fields.
func (t *FindRoomMemberRequestBody) Coerce() error {
	return nil
}

// Bytes creates the byte array.
func (t *FindRoomMemberRequestBody) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// QueryString returns the query string.
func (t *FindRoomMemberRequestBody) QueryString() string {
	queryString := fmt.Sprintf("room_key=%v", t.RoomKey)
	return url.QueryEscape(queryString)
}
