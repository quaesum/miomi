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

func AddAnimalsFile(ctx context.Context, fileSize int64, fileName string, file io.Reader) (int64, error) {
	fmt.Println(fileSize, fileName)
	err := uploadFile(ctx, fileSize, "animals", fileName, file)
	if err != nil {
		return 0, err
	}

	return mysql.CreateFile(ctx, fmt.Sprintf("/animals/%s", fileName))
}

func AddNewsFile(ctx context.Context, fileSize int64, fileName string, file io.Reader) (int64, error) {
	err := uploadFile(ctx, fileSize, "news", fileName, file)
	if err != nil {
		return 0, err
	}

	return mysql.CreateFile(ctx, fmt.Sprintf("/news/%s", fileName))
}

func uploadFile(ctx context.Context, fileSize int64, bucket, newName string, file io.Reader) error {
	/*	tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	*/
	minioClient, err := minio.New("miomi.by:9000", &minio.Options{
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

func GetFilenameById(ctx context.Context, id int64) (*entity.PhotoRequest, error) {
	return mysql.GetUrlByID(ctx, id)
}
