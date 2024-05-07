package serializer_test

import (
	"math/rand"
	"testing"
	"time"

	grpctest "github.com/kiritocyanpine/grpctest/pb/proto"
	"github.com/kiritocyanpine/grpctest/sample"
	"github.com/kiritocyanpine/grpctest/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func TestFileSerialization(t *testing.T) {
	t.Parallel()

	binfile := "../temp/" + randSeq(10) + ".bin"

	lap := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(lap, binfile)

	require.NoError(t, err)
}

func TestProtobufWritenRead(t *testing.T) {
	t.Parallel()

	binfile := "../temp/" + randSeq(10) + ".bin"

	lap := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(lap, binfile)

	require.NoError(t, err)

	lap2 := &grpctest.Laptop{}

	err2 := serializer.ReadProtobufFromBinFile(binfile, lap2)

	require.NoError(t, err2)

	require.True(t, proto.Equal(lap, lap2))
}
