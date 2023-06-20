package ftp

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/jlaffaye/ftp"
)

type FTPServer struct {
	Conn *ftp.ServerConn
}

func NewFTPServer(cf Config) *FTPServer {

	conn, err := ftp.Dial(cf.Addr)
	if err != nil {
		panic("could dial to server")
		// Send an email to admin
	}

	err = conn.Login(cf.User, cf.Password)
	if err != nil {
		panic("could dial to server")
		// Send an email to admin
	}

	fmt.Println("Successful connection with FTP server")

	return &FTPServer{
		Conn: conn,
	}
}

func (s *FTPServer) CreateDir(path string) error {
	return s.Conn.MakeDir(path)
}

func (s *FTPServer) DelDir(path string) error {
	return s.Conn.RemoveDirRecur(path)
}

func (s *FTPServer) AddFile(path string, file *multipart.FileHeader) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error Here")
		return "", err
	}

	err = s.Conn.ChangeDir(path)
	if err != nil {
		return "", err
	}

	err = s.Conn.Stor(file.Filename, fileContent)
	if err != nil {
		s.Conn.Quit()
		return "", err
	}

	return fmt.Sprintf("File with the name %s has been successfuly store it", file.Filename), nil
}

func (s *FTPServer) ReadFiles(path string) ([]string, error) {
	return s.Conn.NameList(path)
}

func (s *FTPServer) GetFile(path string, fileName string) ([]byte, error) {

	file, err := s.Conn.Retr(fmt.Sprintf("%s/%s", path, fileName))
	if err != nil {
		return []byte(""), err
	}

	defer file.Close()

	return io.ReadAll(file)
}

func (s *FTPServer) DeleteFile(path, fileName string) error {
	return s.Conn.Delete(fmt.Sprintf("%s/%s", path, fileName))
}
