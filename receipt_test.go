// Copyright (c) 2026 Roteia
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package whatsmeow

import (
	"testing"

	waBinary "go.mau.fi/whatsmeow/binary"
)

func TestBuildAckAttrsSuppressesStatusMediaType(t *testing.T) {
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
		t.Fatal("expected status media type to be suppressed")
	}
	if got := attrs["class"]; got != "status" {
		t.Fatalf("expected status class, got %v", got)
	}
	if got := attrs["id"]; got != "poison-status-id" {
		t.Fatalf("expected original stanza ID, got %v", got)
	}
	if _, ok := attrs["type"]; ok {
		t.Fatalf("expected a plain ACK without type, got %v", attrs["type"])
	}
}

func TestBuildAckAttrsKeepsOtherReceiptTypes(t *testing.T) {
	node := &waBinary.Node{
		Tag: "receipt",
		Attrs: waBinary.Attrs{
			"id":   "receipt-id",
			"from": "5511999999999@s.whatsapp.net",
			"type": "read",
		},
	}

	attrs, suppressed := buildAckAttrs(node, 0)

	if suppressed {
		t.Fatal("did not expect a normal receipt type to be suppressed")
	}
	if got := attrs["type"]; got != "read" {
		t.Fatalf("expected receipt type to be preserved, got %v", got)
	}
}
