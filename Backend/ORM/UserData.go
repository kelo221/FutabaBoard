package DataModels

import (
	"time"
)

type UserInfo struct {
	IpAddress string `json:"-"`
	UserHash  string `json:"UserHash"`
}

type ImagePrivilege struct {
	IpAddress    string    `json:"-"`
	HashedKey    string    `json:"HashedKey"`
	StartingTime time.Time `json:"StartingTime"`
}

type Bans struct {
	IP               string    `json:"IP"`
	ExpiringTimeUnix time.Time `json:"ExpiringTimeUnix" gorm:"autoCreateTime"`
	Reason           string    `json:"Reason"`
}

// ImageUpload Contains the filename and the hash of the image.
type ImageUpload struct {
	Filename     string `json:"Filename"`
	ImageHash    string `json:"ImageHash,omitempty"`
	OmitFilename bool   `json:"-"`
}

type ImageBans struct {
	ImageHash          string `json:"ImageHash,omitempty"`
	ContentDescription string `json:"ContentDescription"`
}
