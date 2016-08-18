package wire

import (
	"strings"
	"testing"
)

func wirePlainForTest() string {
	return "type:plain|created:1468640833|content:Saturday, 16-Jul-16 03:47:13 UTC hello world log"
}

func wireTopicJSONForTest() string {
	return `topic:awesome|type:json|created:1468640833|content:{"foo": "bar", "awesome": 9001}`
}

func TestParseSingle(t *testing.T) {
	lg := ParseSingle(wirePlainForTest())
	if lg.Type != "plain" {
		t.Errorf("Failed to assign type. Type: %v", lg.Type)
	}
	if lg.Created != int64(1468640833) {
		t.Errorf("Failed to assign created timestamp. Created: %v", lg.Created)
	}
	if lg.Content != "Saturday, 16-Jul-16 03:47:13 UTC hello world log" {
		t.Errorf("Failed to assign content. Content: %v", lg.Content)
	}
}

func TestParseSingleTopicJSON(t *testing.T) {
	lg := ParseSingle(wireTopicJSONForTest())
	if lg.Topic != "awesome" {
		t.Errorf("Failed to assign type. Type: %v", lg.Type)
	}
	if lg.Type != "json" {
		t.Errorf("Failed to assign type. Type: %v", lg.Type)
	}
	if lg.Created != int64(1468640833) {
		t.Errorf("Failed to assign created timestamp. Created: %v", lg.Created)
	}
	if lg.Content != `{"foo": "bar", "awesome": 9001}` {
		t.Errorf("Failed to assign content. Content: %v", lg.Content)
	}
}

func TestParseMultiple(t *testing.T) {
	loglines := strings.Join([]string{wirePlainForTest(), wirePlainForTest(), wirePlainForTest()}, "\n")

	lgs := Parse(loglines)
	for _, lg := range lgs {
		if lg.Type != "plain" {
			t.Errorf("Failed to assign type. Type: %v", lg.Type)
		}
		if lg.Created != int64(1468640833) {
			t.Errorf("Failed to assign created timestamp. Created: %v", lg.Created)
		}
		if lg.Content != "Saturday, 16-Jul-16 03:47:13 UTC hello world log" {
			t.Errorf("Failed to assign content. Content: %v", lg.Content)
		}
	}
}

func TestDifferentContent(t *testing.T) {
	lg := ParseSingle(wirePlainForTest())
	if lg.PlainContent() != lg.Content {
		t.Errorf("Failed to return the correct content: %v", lg.PlainContent())
	}
	if lg.Base64Content() != "U2F0dXJkYXksIDE2LUp1bC0xNiAwMzo0NzoxMyBVVEMgaGVsbG8gd29ybGQgbG9n" {
		t.Errorf("Failed to return the correct content: %v", lg.Base64Content())
	}
}

func TestEncodePlain(t *testing.T) {
	lg := ParseSingle(wirePlainForTest())
	if lg.EncodePlain() != wirePlainForTest() {
		t.Errorf("Failed to encode correctly: %v", lg.EncodePlain())
	}
}

func TestEncodeBase64(t *testing.T) {
	lg := ParseSingle(wirePlainForTest())
	if lg.EncodeBase64() != "type:base64|created:1468640833|content:U2F0dXJkYXksIDE2LUp1bC0xNiAwMzo0NzoxMyBVVEMgaGVsbG8gd29ybGQgbG9n" {
		t.Errorf("Failed to encode correctly: %v", lg.EncodeBase64())
	}
}

func TestEncodeJSON(t *testing.T) {
	lg := ParseSingle(wireTopicJSONForTest())
	if lg.EncodeJSON() != `topic:awesome|type:json|created:1468640833|content:{"foo": "bar", "awesome": 9001}` {
		t.Errorf("Failed to encode correctly: %v", lg.EncodeJSON())
	}
}
