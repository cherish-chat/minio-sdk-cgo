package cgominio

import (
	"testing"
)

func TestNewMinio(t *testing.T) {
	// endpoint:str accessKeyId:str secretAccessKey:str bucketName:str ssl:bool bucketUrl:str region:str
	minio := NewMinio(marshal(map[string]any{
		"endpoint":        "42.194.149.177:9000",
		"accessKeyId":     "ehpDDDeYekG3SS3q",
		"secretAccessKey": "vvOn6zY22uQpwDKK9elbvToXu1J2RXzt",
		"bucketName":      "xxim",
		"ssl":             false,
		"bucketUrl":       "http://42.194.149.177:9000/xxim",
		"region":          "",
	}))
	existObject := minio.ExistObject("111")
	t.Logf("existObject: %v", existObject)
}
