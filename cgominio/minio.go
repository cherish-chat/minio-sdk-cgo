package cgominio

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type minioConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	BucketName      string `json:"bucketName"`
	Ssl             bool   `json:"ssl"`
	BucketUrl       string `json:"bucketUrl"`
	Region          string `json:"region"`
}

var config = minioConfig{}

var client *minio.Client

type Minio struct {
}

func (m *Minio) initClient() {
	tmpClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyId, config.SecretAccessKey, ""),
		Secure: config.Ssl,
		Region: config.Region,
	})
	if err != nil {
		log.Printf("init minio client error: %v", err)
	}
	client = tmpClient
}

// NewMinio 创建minio类
// @param configStr: json字符串 endpoint:str accessKeyId:str secretAccessKey:str bucketName:str ssl:bool bucketUrl:str region:str
func NewMinio(configStr string) *Minio {
	// json decode config
	conf := minioConfig{}
	err := json.Unmarshal([]byte(configStr), &conf)
	if err != nil {
		log.Printf("NewMinio json decode config error: %v", err)
	}
	config = conf
	m := &Minio{}
	m.initClient()
	return m
}

// ExistObject 判断文件是否存在
func (m *Minio) ExistObject(key string) bool {
	_, err := client.StatObject(context.Background(), config.BucketName, key, minio.StatObjectOptions{})
	if err != nil {
		e, ok := err.(minio.ErrorResponse)
		if ok && e.Code == "NoSuchKey" {
			return false
		}
		return false
	}
	return true
}

// PubObjectResult 上传文件返回结果
type PubObjectResult struct {
	Url   string `json:"url"`
	Error string `json:"error"`
}

func marshal(o any) string {
	if o == nil {
		return ""
	}
	b, _ := json.Marshal(o)
	return string(b)
}

// PutObject 上传文件
func (m *Minio) PutObject(objectName string, data []byte) (result string) {
	_, err := client.PutObject(context.Background(), config.BucketName, objectName, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("PutObject error: %v", err)
		return marshal(PubObjectResult{Error: err.Error()})
	}
	return marshal(PubObjectResult{Url: config.BucketUrl + "/" + objectName})
}
