package ip

import "testing"

func TestInternalIP(t *testing.T) {
	ip := InternalIP()
	if ip == "" {
		t.Error("not get ip")
	} else {
		t.Log(ip)
	}
}
