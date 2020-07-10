package exec_demo

import "testing"

func TestRemoteCmd(t *testing.T) {
	// /home/ec2-user/baseinfo
	// 会话[:窗口:面板]
	// work:0  work:1  work:baseinfo
	err := RunSSHCmd("cafe1sandboxdev", "cd /home/ec2-user/baseinfo && ls; pwd&& tmux capture-pane -t work:2 && tmux show-buffer")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoteCmd2(t *testing.T) {
	// /home/ec2-user/baseinfo
	// 会话[:窗口:面板]
	// work:0  work:1  work:baseinfo
	err := RunSSHCmd("cafe1sandboxdev", `tmux send-keys -t work:3 "aaaaaaa" C-m && sleep 0.1s && tmux capture-pane -t work:3 && tmux show-buffer`)
	if err != nil {
		t.Fatal(err)
	}
}
