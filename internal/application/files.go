package application

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	cr "github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"madmax/internal/application/db/mysql"
	"madmax/internal/entity"
)

func AddFile(ctx context.Context, fileSize int64, fileName string, file io.Reader, target string) (int64, error) {
	fmt.Println(fileSize, fileName)
	err := uploadFile(ctx, fileSize, target, fileName, file)
	if err != nil {
		return 0, err
	}
	return mysql.CreateFile(ctx, fmt.Sprintf("/%s/%s", target, fileName))

}

func uploadFile(ctx context.Context, fileSize int64, bucket, newName string, file io.Reader) error {
	/*	tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	*/
	//miomi.by:9000
	minioClient, err := minio.New("localhost:9000", &minio.Options{
		Creds:  cr.NewStaticV4("admin", "admin3000", ""),
		Secure: false,
		//Transport: tr,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = minioClient.PutObject(context.Background(), bucket, newName, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetFileNameAndId(ctx context.Context) ([]entity.PhotoRequest, error) {
	return mysql.GetUrlAndId(ctx)
}
