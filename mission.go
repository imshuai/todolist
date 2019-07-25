package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type mission struct {
	Name            string      `json:"name"`
	CreateTime      time.Time   `json:"create_time"`
	DeadLineTime    time.Time   `json:"deadline_time"`
	CompleteTime    time.Time   `json:"complete_time"`
	Type            missionType `json:"type"`
	TaskDescription string      `json:"description"`
	Comment         string      `json:"comment"`
}

type missionType string

var (
	industry  missionType = "工商"
	tax       missionType = "税务"
	insurance missionType = "社保"
	other     missionType = "其它"
)

func (m *mission) Deserialize(v []byte) error {
	t := m
	dec := gob.NewDecoder(bytes.NewBuffer(v))
	err := dec.Decode(t)
	if err != nil {
		return err
	}
	m = t
	return nil
}

func (m *mission) Serialize() ([]byte, error) {
	buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
