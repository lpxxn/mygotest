package exec_demo

import "testing"

func TestRemoteCmd(t *testing.T) {
	// /home/ec2-user/baseinfo
	err := RunSSHCmd("cafe1sandboxdev", "cd /home/ec2-user/baseinfo && ls; pwd", "")
	if err != nil {
		t.Fatal(err)
	}
}
