package mobile

func (e *Engine) StopNode() error {

	if e.listener != nil {
		e.listener.Close()
	}

	e.manager.DisconnectAll()

	return nil
}