package cmd

type AccessLockDispatcher struct {
	pool chan *worker
	queue chan interface{}
	workers []*worker
	wg sync.WaitGroup
	quit chan struct{}
}


func() initDriveLock(drivenumber int) error{
	chan := make(chan bool, drivenumber)
}


func(d *AccessLockDispatcher) GetDriveLock(bucket, object string, offset int64, length int64) (err error){
	w.start()
	for{
		select {
		case v := d<-queue:

		}
	}
	return 0
}

func ReleaseDriveLock(){
	<-accessLocks.lock
}