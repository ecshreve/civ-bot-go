package bot

func (s *DiscordSession) Open() error {
	mock.Input(interface{}(s.Open))
	return nil
}

func (s *DiscordSession) Close() error {
	mock.Input(interface{}(s.Close))
	return nil
}

func (s *DiscordSession) AddHandler(handler interface{}) func() {
	mock.Input(interface{}(s.AddHandler), handler)
	return func() {}
}
