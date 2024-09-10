package oss

import (
	"context"
	"io"
	"log"
	"mime"
	"net/url"
	"path/filepath"
	"time"

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
	contentType := getContentTypeFromFileName(objectName)
	info, err := MinioClient.PutObject(context.Background(), bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	log.Printf("Upload Success, ETag: %s\n", info.ETag)
	return nil
}

func GetFileURL(bucketName, objectName string) (u *url.URL, err error) {
	reqParams := make(url.Values)
	// content-type is image
	contentType := getContentTypeFromFileName(objectName)
	reqParams.Set("content-type", contentType)
	u, err = MinioClient.PresignedGetObject(context.Background(), bucketName, objectName, time.Second*24*60*60, reqParams)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func getContentTypeFromFileName(filename string) string {
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		return "application/octet-stream"
	}
	return contentType
}
