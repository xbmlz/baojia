package oss

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

const (
	// 7FJO2bkJ9H24eDlP6JaE
	// MxOcLfwNLaEq8HD4EkMAksuUgxHAYHKYvxvWPZaI
	endpoint        = "193.112.175.178:9000"
	accessKeyID     = "7FJO2bkJ9H24eDlP6JaE"
	secretAccessKey = "MxOcLfwNLaEq8HD4EkMAksuUgxHAYHKYvxvWPZaI"
	useSSL          = false
)

func InitMinioClient() {
	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (err error) {
	info, err := MinioClient.PutObject(context.Background(), bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	log.Printf("Upload Success, ETag: %s\n", info.ETag)
	return nil
}
