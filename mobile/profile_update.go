package mobile

import (
    "connect6/data"
    "errors"
    "strings"
)

func (e *Engine) UpdateName(name string) error {
    if e.identity == nil {
        return errors.New("engine not initialized")
    }

    name = strings.TrimSpace(name)

    if name == "" {
        return errors.New("name cannot be empty")
    }

    if len(name) > 30 {
        return errors.New("name too long")
    }

    e.identity.Name = name

    return data.SaveIdentity(e.identity)
}