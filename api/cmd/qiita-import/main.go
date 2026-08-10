package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/service"
)

func main() {
	ctx := context.Background()
	if err := library.InitDB(); err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	db := library.GetDB(ctx)
	var err error
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer sqlDB.Close()

	qiitaClient := external.NewQiitaClient(nil)
	count, err := service.ImportQiitaItems(ctx, db, qiitaClient, qiitaClient)
	if err != nil {
		log.Fatalf("Qiita article import failed: %v", err)
	}
	fmt.Printf("imported %d Qiita articles\n", count)
}
