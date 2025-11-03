package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func generateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func getExecutablePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not determine executable path")
	}
	return filepath.Abs(filepath.Dir(filename))
}

func getAwsS3Client() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return s3.New(sess)
}

func getFilenameFromRequest(r *http.Request) string {
	return filepath.Base(r.URL.Path)
}

func getFileNameFromExtension(filename string) (string, string) {
	ext := filepath.Ext(filename)
	return strings.TrimSuffix(filename, ext), ext
}

func getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

var (
	mu sync.Mutex
)

func getLock() *sync.Mutex {
	mu.Lock()
	return &mu
}