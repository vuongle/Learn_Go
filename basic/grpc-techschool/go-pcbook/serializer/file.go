package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtoBufToBinaryFile(message proto.Message, filename string) error {
	// serialize a proto message to binary
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can't serialize to binary: %w", err)
	}

	// write binary to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can't write binary to file: %w", err)
	}

	return nil
}

func WriteProtoBufToJSONFile(message proto.Message, filename string) error {
	// serialize a proto message to json string
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("can't serialize to json: %w", err)
	}

	// write json string to file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can't write json string to file: %w", err)
	}

	return nil
}

func ReadProtoBufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can't read binary from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("can't deserialize from binary: %w", err)
	}

	return nil
}

func ReadProtoBufFromJSONFile() {

}
