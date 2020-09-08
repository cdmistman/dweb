package ssh

func (serv *Service) Start() {
	for _, sess := range serv.Config {
		go sess.Run()
	}
}
