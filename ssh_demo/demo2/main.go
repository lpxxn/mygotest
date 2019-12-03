package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strings"
	"syscall"

	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	host             = "cafetest1dev"  // テスト用ホスト
	defaultSshConfig = "~/.ssh/config" // ssh/config
)

func main() {
	// current user
	usr, _ := user.Current()
	confFile := strings.Replace(defaultSshConfig, "~", usr.HomeDir, 1)
	// .ssh/configを開いてDecode
	f, _ := os.Open(confFile)
	cfg, _ := ssh_config.Decode(f)
	// 情報の取得
	user, _ := cfg.Get(host, "User")
	addr, _ := cfg.Get(host, "HostName")
	port, _ := cfg.Get(host, "Port")
	port = "22"
	keyPath, _ := cfg.Get(host, "IdentityFile")
	proxyCommand, _ := cfg.Get(host, "ProxyCommand")
	// keyPathを置換
	keyPath = strings.Replace(keyPath, "~", usr.HomeDir, 1)
	// ProxyCommandを置換
	proxyCommand = strings.Replace(proxyCommand, "%h", addr, -1)
	proxyCommand = strings.Replace(proxyCommand, "%p", port, -1)
	// sshConfigの生成
	auth := []ssh.AuthMethod{}
	key, _ := ioutil.ReadFile(keyPath)
	signer, _ := ssh.ParsePrivateKey(key)
	auth = append(auth, ssh.PublicKeys(signer))
	sshConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // FIXME
	}
	fmt.Println(addr)
	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(keyPath)
	fmt.Println(proxyCommand)
	// net.Pipeの作成
	cli, srv := net.Pipe()
	cmd := exec.Command("bash", "-c", proxyCommand)
	// 標準入出力の指定
	cmd.Stdin = srv
	cmd.Stdout = srv
	cmd.Stderr = os.Stderr
	// コマンドの実行
	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}

	// // proxy経由での接続を実施
	conn, incomingChannels, incomingRequests, err := ssh.NewClientConn(cli, net.JoinHostPort(addr, port), sshConfig)
	if err != nil {
		os.Exit(1)
	}
	// proxy経由でのsshClientの作成
	client := ssh.NewClient(conn, incomingChannels, incomingRequests)
	// client, _ := ssh.Dial("tcp", net.JoinHostPort(addr, port), sshConfig)
	// ↓共通コード
	// Create Session
	session, err := client.NewSession()
	defer session.Close()
	// キー入力を接続先が認識できる形式に変換する(ここがキモ)
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		fmt.Println(err)
	}
	defer terminal.Restore(fd, state)
	// ターミナルサイズの取得
	w, h, err := terminal.GetSize(fd)
	if err != nil {
		fmt.Println(err)
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", h, w, modes)
	if err != nil {
		fmt.Println(err)
	}
	// log := new(bytes.Buffer)
	logFile, _ := os.OpenFile("./ssh_term_with_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	session.Stdout = io.MultiWriter(os.Stdout, logFile)
	session.Stderr = io.MultiWriter(os.Stderr, logFile)
	session.Stdin = os.Stdin
	err = session.Shell()
	if err != nil {
		fmt.Println(err)
	}
	// ターミナルサイズの変更検知・処理
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGWINCH)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGWINCH:
				fd := int(os.Stdout.Fd())
				w, h, _ = terminal.GetSize(fd)
				session.WindowChange(h, w)
			}
		}
	}()

	session.Start("cd baseinfo")
	err = session.Wait()
	fmt.Println("after wait")
	if err != nil {
		fmt.Println(err)
	}
}
/*
https://orebibou.com/2019/06/golang%E3%81%A7-ssh-config%E3%82%92%E8%AA%AD%E3%81%BF%E8%BE%BC%E3%82%93%E3%81%A7ssh%E6%8E%A5%E7%B6%9A%E3%81%99%E3%82%8Bproxycommand%E5%AF%BE%E5%BF%9C/
 */
