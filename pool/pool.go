package pool

type GoPool struct {
	MaxLimit int

	tokenChan chan struct{}
}

type GoPoolOption func(*GoPool)

func WithMaxLimit(max int) GoPoolOption {
	return func(gp *GoPool) {
		gp.MaxLimit = max
		gp.tokenChan = make(chan struct{}, gp.MaxLimit)

		for i := 0; i < gp.MaxLimit; i++ {
			gp.tokenChan <- struct{}{}
		}
	}
}

func NewGoPool(options ...GoPoolOption) *GoPool {
	p := &GoPool{}
	for _, o := range options {
		o(p)
	}

	return p
}

// Submit will wait a token, and then execute fn
func (gp *GoPool) Submit(fn func()) {
	token := <-gp.tokenChan // if there are no tokens, we'll block here

	go func() {
		fn()
		gp.tokenChan <- token
	}()
}

// Wait will wait all the tasks executed, and then return
func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.tokenChan
	}

	close(gp.tokenChan)
}

func (gp *GoPool) size() int {
	return len(gp.tokenChan)
}
