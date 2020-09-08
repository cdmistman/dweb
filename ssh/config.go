package ssh

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Service struct {
	Config []Session
}

type Session struct {
	Name         string      `json:"name"`
	Mode         SessionMode `json:"mode"`
	Port         int64       `json:"port"`
	AuthRequired *bool       `json:"auth required"`
	MaxAuthTries *int        `json:"max auth tries"`

	// only when (mode & ModeDockerized == ModeDockerized)
	Location *string  `json:"location"`
	Links    *[]Link  `json:"links"`
	Images   *[]Image `json:"images"`
	Users    *[]User  `json:"users"`
}

type SessionMode int

const (
	ModeNative SessionMode = iota
	ModeDockerized
)

func (mode *SessionMode) UnmarshalJSON(bb []byte) error {
	var str string
	if err := json.Unmarshal(bb, &str); err != nil {
		return err
	}

	switch strings.ToLower(str) {
	case "native":
		*mode = ModeNative
	case "dockerized":
		*mode = ModeDockerized
	default:
		return errors.New(fmt.Sprintf("Unrecognized session mode: %s", str))
	}
	return nil
}

func (mode SessionMode) MarshalJSON() ([]byte, error) {
	var str string
	switch mode {
	case ModeNative:
		str = "native"
	case ModeDockerized:
		str = "dockerized"
	default:
		return nil, errors.New(fmt.Sprintf("Unrecognized session mode: %s", str))
	}
	return json.Marshal(str)
}

type Link struct {
	Local  string `json:"local"`
	Target string `json:"target"`
}

type Image struct {
	Name  string         `json:"name"`
	Src   *string        `json:"src"`
	Image *[]Instruction `json:"image"`
}

type Instruction struct {
	Instruction string `json:"instruction"`
	Args        string `json:"args"`
}

type User struct {
	Name         string `json:"name"`
	Passwordless bool   `json:"passwordless"`
}
