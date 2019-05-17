package callerinfo

import "testing"

func TestRetrieveCallInfo(t *testing.T) {
	rev := RetrieveCallInfo()
	t.Logf("%#v", rev)
}
