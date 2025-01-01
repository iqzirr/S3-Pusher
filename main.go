// FileUploader.go MinIO example
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Define flags
var minioObjectName = flag.String("objectname", "", "minio object name")
var minioBucketName = flag.String("bucket", "", "Bucket Name")
var minioEndpoint = flag.String("endpoint", "", "S3 Endpoint (IP or URL)")
var minioSSLMode = flag.Bool("ssl", false, "SSL (True/false),if https > use true (Default false)")
var minioFilePath = os.Args[1:] // Exclude the program name
var minioAccessKey = flag.String("accesskey", "", "Your S3 access key")
var minioSecretKey = flag.String("secretkey", "", "Your S3 secret key")

func main() {
	// Check if the flag is provided
	if *minioObjectName == "" {
		fmt.Println("object name is empty, please fill it first... (-h for help)")
		return
	} else if *minioBucketName == "" {
		fmt.Println("Bucket is empty, please fill it first... (-h for help)")
		return
	} else if *minioAccessKey == "" {
		fmt.Println("Access key is empty, please fill it first... (-h for help)")
		return
	} else if *minioSecretKey == "" {
		fmt.Println("Secret key is empty, please fill it first... (-h for help)")
		return
	} else if *minioEndpoint == "" {
		fmt.Println("Endpoint is empty, please fill it first... (-h for help)")
		return
	} else if len(minioFilePath) == 0 {
		fmt.Println("Please specify your file location (-h for help)")
		return
	} else {
		minioPush()
	}
}

func minioPush() {
	ctx := context.Background()
	endpoint := *minioEndpoint
	useSSL := *minioSSLMode
	accessKey := *minioAccessKey
	secretKey := *minioSecretKey
	bucketName := *minioBucketName

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := *minioObjectName
	contentType := "application/octet-stream"
	filePath := strings.Join(minioFilePath, " ")

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
