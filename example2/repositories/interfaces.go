package repositories

type Proxy interface {
	Call(target string)  ([]byte, error)
}
