package typhenapi

import (
	"bytes"
	"errors"
	"testing"
)

func TestNewMessage(t *testing.T) {
	serializer := new(JSONSerializer)
	message, err := NewMessage(serializer, 1, &testType{"Foobar"})

	if err != nil {
		t.Error(err)
		return
	}
	if message.Type != 1 {
		t.Errorf("message.Type is expected to equal 1: %v", message.Type)
		return
	}

	deserialized := &testType{}
	deserializedErr := serializer.Deserialize(message.Body, deserialized)

	if deserializedErr != nil {
		t.Error(deserializedErr)
		return
	}
	if deserialized.Message != "Foobar" {
		t.Errorf("deserialized.Message is expected to equal Foobar: %v", deserialized.Message)
		return
	}
}

func TestNewMessageFromBytes(t *testing.T) {
	data := []byte{0xF0, 0xFF, 0x00, 0x00, 2, 3, 5, 7, 11}
	message, err := NewMessageFromBytes(data)

	if err != nil {
		t.Error(err)
		return
	}
	if message.Type != 65520 {
		t.Errorf("message.Type is expected to equal 65520: %v", message.Type)
		return
	}
	if !bytes.Equal([]byte{2, 3, 5, 7, 11}, message.Body) {
		t.Errorf("message.Body is expected to equal %v: %v", []byte{2, 3, 5, 7, 11}, message.Body)
		return
	}
}

func TestMessageBytes(t *testing.T) {
	serializer := new(JSONSerializer)
	messageA, _ := NewMessage(serializer, 1, &testType{"Foobar"})
	message, _ := NewMessageFromBytes(messageA.Bytes())

	deserialized := &testType{}
	err := serializer.Deserialize(message.Body, deserialized)

	if err != nil {
		t.Error(err)
		return
	}
	if deserialized.Message != "Foobar" {
		t.Errorf("deserialized.Message is expected to equal Foobar: %v", deserialized.Message)
		return
	}
}

type testType struct {
	Message string `json:"message"`
}

func (t *testType) Coerce() error {
	if t.Message == "" {
		return errors.New("Message is empty")
	}

	return nil
}

func (t *testType) Bytes(serializer Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
