package mobile

func (e *Engine) GetProfile() *MyProfile {

    profile, err := e.Profile()
    if err != nil {
        return &MyProfile{}
    }

    return &profile
}