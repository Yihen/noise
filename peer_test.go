package noise

import (
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/quick"
)

type testMsg struct {
	Text string
}

func (testMsg) Read(reader payload.Reader) (Message, error) {
	text, err := reader.ReadString()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read test message")
	}

	return testMsg{Text: text}, nil
}

func (m testMsg) Write() []byte {
	return payload.NewWriter(nil).WriteString(m.Text).Bytes()
}

func TestEncodeMessage(t *testing.T) {
	resetOpcodes()

	o := RegisterMessage(Opcode(123), (*testMsg)(nil))
	assert.Equal(t, o, RegisterMessage(o, (*testMsg)(nil)))

	p := newPeer(nil, nil)

	f := func(msg testMsg) bool {
		bytes, err := p.EncodeMessage(msg)
		assert.Nil(t, err)
		assert.Equal(t, append([]byte{byte(o)}, msg.Write()...), bytes)

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestDecodeMessage(t *testing.T) {
	resetOpcodes()

	o := RegisterMessage(Opcode(45), (*testMsg)(nil))
	assert.Equal(t, o, RegisterMessage(o, (*testMsg)(nil)))

	p := newPeer(nil, nil)

	f := func(msg testMsg) bool {
		resultO, resultM, err := p.DecodeMessage(append([]byte{byte(o)}, msg.Write()...))
		assert.Nil(t, err)
		assert.Equal(t, o, resultO)
		assert.Equal(t, msg, resultM)

		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
