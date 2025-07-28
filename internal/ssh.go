package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
)

func UploadFile(host, user, password, localFile, remotePath string) error {
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return err
	}
	defer client.Close()
	
	sess, err := client.NewSession()
	if err != nil {
		return err
	}
	defer sess.Close()
	
	src, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer src.Close()
	
	go func() {
		w, _ := sess.StdinPipe()
		defer w.Close()
		info, _ := src.Stat()
		fmt.Fprintf(w, "C0644 %d %s\n", info.Size(), filepath.Base(remotePath))
		io.Copy(w, src)
		fmt.Fprint(w, "\x00")
	}()
	
	return sess.Run(fmt.Sprintf("scp -t %s", remotePath))
}
