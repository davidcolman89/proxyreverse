package repositories

type proxyRepo struct {

}

func NewProxyRepo() Proxy{
	return proxyRepo{}
}

func (r proxyRepo) Call()  {

}