package impl

import (
	"io"
	"log"
	"os"

	f "go-ssh-sftp/sftp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sftpFileRemoteToLocalImpl struct {
	host 				string
	username			string
	password			string
	srcFile				string
	destinationFile 	string
	sshConfig			*ssh.ClientConfig
}


func SftpFileRemoteToLocal(host, username, password, sshKey, srcFile, destinationFile string) (f.ISftpFileRemoteToLocal, error) {
	// initial connection
	init, err := initialConnectionWithPublicKey(host, username, sshKey)
	// init, err := initialConnection(host, username, password)
	if err != nil {
		return nil, err
	}

	return &sftpFileRemoteToLocalImpl{
		host:            host,
		username:        username,
		password:        password,
		srcFile:         srcFile,
		destinationFile: destinationFile,
		sshConfig: init,
	}, nil
}


// InitialConnection implements sftp.ISftpFileRemoteToLocal.
func initialConnection(host, username, password string) (*ssh.ClientConfig, error) {
    sshConfig := &ssh.ClientConfig{
        User: username,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
    }

	return sshConfig, nil
}

// InitialConnection implements sftp.ISftpFileRemoteToLocal.
func initialConnectionWithPublicKey(host, username, ssh_key string) (*ssh.ClientConfig, error) {
    sshConfig := &ssh.ClientConfig{
        User: username,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            PublicKeyFile(ssh_key),
        },
    }

	return sshConfig, nil
}

// DownloadFile implements sftp.ISftpFileRemoteToLocal.
func (s *sftpFileRemoteToLocalImpl) DownloadFile() error {
	
	client, err := ssh.Dial("tcp", s.host, s.sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		return err
	}

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return err
	}

	fSource, err := sftpClient.Open(s.srcFile)
	if err != nil {
		return err
	}

	defer fSource.Close()

	fDestination, err := os.Create(s.destinationFile)
	if err != nil {
		return err
	}

	defer fDestination.Close()

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		return err
	}

	log.Println("File copied.")
	return nil
}

func PublicKeyFile(file string) ssh.AuthMethod {
    buffer, err := os.ReadFile(file)
    if err != nil {
        return nil
    }

    key, err := ssh.ParsePrivateKey(buffer)
    if err != nil {
        return nil
    }

    return ssh.PublicKeys(key)
}