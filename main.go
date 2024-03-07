package main

import (
	"fmt"
	"go-ssh-sftp/sftp/impl"
)

const HOST = "<------>"
const SSH_USERNAME = "<---->"
const SSH_PASSWORD = "<---->"
const SSH_KEY = "<---->"
const SRCFILE = "<---->"
const DESTFILE = "<---->"
const FILENAME = "test-file.txt"

func main(){

	remoteSftp, err := impl.SftpFileRemoteToLocal(
		HOST,
		SSH_USERNAME,
		SSH_PASSWORD,
		SSH_KEY,
		SRCFILE + FILENAME,
		DESTFILE + FILENAME,
	)
	if err != nil {
		fmt.Println(err)
	}

	err = remoteSftp.DownloadFile()
	if err != nil {
		fmt.Println(err)
	}
}