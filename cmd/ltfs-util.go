package cmd

import (
    "sync"
    "syscall"
)

type FileLock struct{
	l sync.Mutex
	fd int
}

func NewFileLock(filename string) (*FileLock, error) {
	var err error
	if filename == "" {
		return nil, errInvalidArgument
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT | syscall.O_RDONLY, 0750)
	if err != nil{
		return nil, err
	}
	return &FileLock{fd: fd}, nil
}

func (m *FileLock) Lock() error {
	m.l.Lock()
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		return err
	}
	return nil
}

func (m *FileLock) Unlock() error {
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		return err
	}
	m.l.Unlock()
	return nil
}

/*
func main() {
	l := NewFileLock("main.go")
	fmt.Println("try  locking...")
	l.Lock()
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock")
}
*/
