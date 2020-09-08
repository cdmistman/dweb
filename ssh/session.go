package ssh

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

var (
	DwebVersion    string
	DwebSSHVersion string
)

func (sess *Session) Run() {
	log.Printf("Starting session %s\n", sess.Name)

	server := ssh.ServerConfig{}
	server.ServerVersion = fmt.Sprintf("SSH-2.0-dweb%s-ssh%s", DwebVersion, DwebSSHVersion)

	if sess.AuthRequired != nil {
		server.NoClientAuth = !(*sess.AuthRequired)
	} else {
		server.NoClientAuth = false
	}

	if sess.MaxAuthTries != nil {
		server.MaxAuthTries = *(sess.MaxAuthTries)
	}

	server.PublicKeyCallback = func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
		return sess.publicKeyCallback(conn, key)
	}

	server.KeyboardInteractiveCallback = func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
		return sess.keyboardInteractiveCallback(conn, client)
	}

	server.AuthLogCallback = func(conn ssh.ConnMetadata, method string, err error) {
		sess.authLogCallback(conn, method, err)
	}

	server.BannerCallback = func(conn ssh.ConnMetadata) string {
		return sess.bannerCallback(conn)
	}
}

func (sess *Session) publicKeyCallback(conn ssh.ConnMetadata, key ssh.PublicKey) (perms *ssh.Permissions, err error) {
	return
}

func (sess *Session) keyboardInteractiveCallback(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (perms *ssh.Permissions, err error) {
	return
}

func (sess *Session) authLogCallback(conn ssh.ConnMetadata, method string, err error) {
}

func (sess *Session) bannerCallback(conn ssh.ConnMetadata) string {
	return ""
}
