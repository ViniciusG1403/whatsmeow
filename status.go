package whatsmeow

import (
	"context"

	waBinary "go.mau.fi/whatsmeow/binary"
)

// handleStandaloneStatus acknowledges WhatsApp story/status stanzas without
// exposing them as application messages. Ignoring status content must happen
// after this protocol ACK; otherwise the stanza remains at the head of the
// offline queue and WhatsApp repeatedly closes the stream.
func (cli *Client) handleStandaloneStatus(ctx context.Context, node *waBinary.Node) {
	cli.sendAck(ctx, node, 0)
}
