package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {

	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not unmarshall protobuff message to bytes due to %s", err)
	}

	return os.WriteFile(filename, data, 0655)
}

func ReadProtobufFromBinFile(filename string, msg proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read file %s", err)
	}

	return proto.Unmarshal(data, msg)
}
