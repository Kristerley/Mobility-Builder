package data

import (
    "time"
)

type Exercise struct {
    ID int64
    CreatedAt time.Time `json:"-"`
    Name string
    MainBodyPart string
    SecondaryBodyParts []string `json:",omitempty"`
    Purpose string // mobility, pain relief etc
    Language string
    URL string
    URLType string
    Version int32
}
