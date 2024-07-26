package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	marshaler := protojson.MarshalOptions{
		Multiline:         true,
		Indent:            "  ",
		UseProtoNames:     true,
		EmitUnpopulated:   true,
		EmitDefaultValues: true,
		UseEnumNumbers:    false,
	}

	data, err := marshaler.Marshal(message)
	if err != nil {
		return nil, err
	}

	return data, nil
}
