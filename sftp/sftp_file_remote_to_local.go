package sftp

type ISftpFileRemoteToLocal interface {
	DownloadFile() error
}