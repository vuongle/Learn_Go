package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vuongle/grpc/sample"
	"github.com/vuongle/grpc/serializer"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	// test case 1
	err := serializer.WriteProtoBufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	// test case 2
	laptop2 := sample.NewLaptop()
	err = serializer.ReadProtoBufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	// test case 3
	err = serializer.WriteProtoBufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
