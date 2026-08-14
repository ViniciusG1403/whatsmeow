package whatsmeow

import (
	"context"
	"testing"

	waBinary "go.mau.fi/whatsmeow/binary"
	"go.mau.fi/whatsmeow/store"
)

func TestStreamErrorReconnectIsMarkedExpectedSynchronously(t *testing.T) {
	client := NewClient(&store.Device{}, nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	client.reconnectAfterStreamError(ctx, "stream-error-test")

	if !client.isExpectedDisconnect() {
		t.Fatal("stream error reconnect must suppress the immediate EOF auto-reconnect race")
	}
}

func TestClientRegistersStandaloneStatusHandler(t *testing.T) {
	client := NewClient(&store.Device{}, nil)

	if client.nodeHandlers["status"] == nil {
		t.Fatal("standalone status stanzas must be handled so their protocol ACK is sent")
	}
}

func TestStandaloneStatusMediaUsesPlainAck(t *testing.T) {
	node := &waBinary.Node{
		Tag: "status",
		Attrs: waBinary.Attrs{
			"id":   "poison-status-id",
			"from": "status@broadcast",
			"type": "media",
		},
	}

	attrs, suppressed := buildAckAttrs(node, 0)
	if !suppressed {
		t.Fatal("expected media type to be suppressed for a standalone status ACK")
	}
	if _, exists := attrs["type"]; exists {
		t.Fatalf("expected plain status ACK, got type=%v", attrs["type"])
	}
	if attrs["class"] != "status" || attrs["id"] != "poison-status-id" {
		t.Fatalf("plain ACK lost status identity: %#v", attrs)
	}
}
