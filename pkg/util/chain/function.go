package chain

// HandleFunc functions to be executed
// ret: function return
// cont: if execute the next function
// err: executed error, when error is returned, chain will be stopped
type HandleFunc[T any] func() (ret T, cont bool, err error)

// Call makes a call chain of many functions
func Call[T any](chains ...HandleFunc[T]) (res T, err error) {
	for _, chain := range chains {
		var cont bool
		res, cont, err = chain()
		if !cont {
			break
		}
	}
	return
}
