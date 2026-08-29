package main

import (
	"context"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/kikudesuyo/buildlog/api/library"
)

var assets = map[string]string{
	"profile.jpg":        "profile/profile.jpg",
	"whichway-icon.svg":  "apps/whichway-icon.svg",
	"mahjong-icon.svg":   "apps/mahjong-icon.svg",
	"pratan-icon.svg":    "apps/pratan-icon.svg",
	"economeye-icon.svg": "apps/economeye-icon.svg",
}

func main() {
	ctx := context.Background()
	bucketName := library.Env("GCS_BUCKET")
	if bucketName == "" {
		panic("GCS_BUCKET is not set")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Errorf("create storage client: %w", err))
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	if err := ensureBucket(ctx, bucket, bucketName, library.Env("GCP_PROJECT")); err != nil {
		panic(err)
	}

	for sourceName, objectName := range assets {
		if err := upload(ctx, bucket, sourceName, objectName); err != nil {
			panic(err)
		}
		fmt.Printf("%s -> %s\n", sourceName, library.PublicAssetURL(objectName))
	}
}

func ensureBucket(ctx context.Context, bucket *storage.BucketHandle, bucketName, projectID string) error {
	_, err := bucket.Attrs(ctx)
	if err == nil {
		return nil
	}
	if projectID == "" {
		projectID = "buildlog-local"
	}
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		return fmt.Errorf("create bucket %q: %w", bucketName, err)
	}
	return nil
}

func upload(ctx context.Context, bucket *storage.BucketHandle, sourceName, objectName string) error {
	sourcePath := filepath.Join("..", "web", "static", sourceName)
	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("open %q: %w", sourcePath, err)
	}
	defer file.Close()

	writer := bucket.Object(objectName).NewWriter(ctx)
	writer.ContentType = mime.TypeByExtension(filepath.Ext(sourceName))
	writer.CacheControl = "public, max-age=31536000, immutable"
	if _, err := io.Copy(writer, file); err != nil {
		_ = writer.Close()
		return fmt.Errorf("upload %q: %w", sourceName, err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("finalize upload %q: %w", sourceName, err)
	}
	return nil
}
