package parameterHelper

import (
	"testing"
)

func TestParseISO8601Date(t *testing.T) {

	iso8601Date := "2014-12-22T10:33:40Z"

	if _, err := ParseFromISO8601Time(iso8601Date); err != nil {
		t.Error("TestParseRFC2616Date failed")
	}
}

func TestParseRFC2616Date(t *testing.T) {
	rfc2616Date := "Wed, 16 Jan 2013 19:01:18 GMT"

	if _, err := ParseFromRFC2616Date(rfc2616Date); err != nil {
		t.Error("TestParseRFC2616Date failed")
	}
}

func TestParseToISO8601Time(t *testing.T) {
	iso8601Date := "2014-12-22T10:33:40Z"
	isoDt, _ := ParseFromISO8601Time(iso8601Date)

	if ParseToISO8601Time(isoDt) != iso8601Date {
		t.Error("TestParseToISO8601Time faield")
	}
}

func TestParseToRFC2616Date(t *testing.T) {
	rfc2616Date := "Wed, 16 Jan 2013 19:01:18 GMT"
	rfcDt, _ := ParseFromRFC2616Date(rfc2616Date)

	if ParseToRFC2616Date(rfcDt) != rfc2616Date {
		t.Error("TestParseToRFC2616Date faield")
	}
}

func TestGetSignatureNonce(t *testing.T) {
	if GetSignatureNonce() == GetSignatureNonce() {
		t.Error("TestGetSignatureNonce failed")
	}
}
