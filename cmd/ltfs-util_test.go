package cmd

import (
	"os"
	"testing"
	"io/ioutil"
	"time"
)

func LockWork(q chan bool, t *testing.T){
		var err error
		var l *FileLock
		l, err = NewFileLock("./testfile")
		if err != nil {
			t.Errorf("Expecting error invalid argument, got %s", err)
		}
		err = l.Lock()
		if err != nil {
			t.Errorf("Expecting error invalid argument, got %s", err)
		}
		time.Sleep(10 * time.Second)
		err = l.Unlock()
		if err != nil {
			t.Errorf("Expecting error invalid argument, got %s", err)
		}
		l, err = NewFileLock("./testfile")
		q <- true
}



func TestNewLock(t *testing.T) {
	ioutil.WriteFile("./testfile", []byte("hello"), os.ModePerm)
	defer os.Remove("./testfile")
	quit := make(chan bool)

	go LockWork(quit, t)
	go LockWork(quit, t)
	q1, q2 := <-quit, <-quit
	_, _ = q1, q2

}