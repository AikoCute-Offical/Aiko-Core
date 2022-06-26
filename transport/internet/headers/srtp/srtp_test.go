package srtp_test

import (
	"context"
	"testing"

	"github.com/AikoCute-Offical/Aiko-Core/common"
	"github.com/AikoCute-Offical/Aiko-Core/common/buf"
	. "github.com/AikoCute-Offical/Aiko-Core/transport/internet/headers/srtp"
)

func TestSRTPWrite(t *testing.T) {
	content := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	srtpRaw, err := New(context.Background(), &Config{})
	common.Must(err)

	srtp := srtpRaw.(*SRTP)

	payload := buf.New()
	srtp.Serialize(payload.Extend(srtp.Size()))
	payload.Write(content)

	expectedLen := int32(len(content)) + srtp.Size()
	if payload.Len() != expectedLen {
		t.Error("expected ", expectedLen, " of bytes, but got ", payload.Len())
	}
}
