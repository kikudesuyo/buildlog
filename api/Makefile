.PHONY: build dev dev-auth lint lint-fix docker-build-stage docker-build-prod _docker-build deploy-stage deploy-prod __deploy

GO111MODULE=on
LOAD_ENV=. ./default_env.sh
APP_BINARY ?= app
GCP_PROJECT_DEV ?=
GCP_PROJECT_PROD ?=
API_DOMAIN_STAGE ?=
API_DOMAIN_PROD ?=

build:
	go build -o ./$(APP_BINARY) .

dev:
	$(LOAD_ENV) ; \
	go run main.go

dev-auth:
	gcloud auth application-default login \
	  --scopes=https://www.googleapis.com/auth/cloud-platform

lint:
	@type golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./... --fix

docker-build-stage:
	make _docker-build project=${GCP_PROJECT_DEV}

docker-build-prod:
	make _docker-build project=${GCP_PROJECT_PROD}

_docker-build:
	gcloud builds submit --default-buckets-behavior="regional-user-owned-bucket" --region=asia-northeast1 --config=_cloudbuild.all.yaml --project ${project} .

# 本番環境 Cloud Runへデプロイ
deploy-prod:
	make __deploy project=${GCP_PROJECT_PROD} bucketprefix=prod resourceprefix=prod apidomain=${API_DOMAIN_PROD} is_production=true

# ステージング環境 Cloud Runへデプロイ
deploy-stage:
	make __deploy project=${GCP_PROJECT_DEV} bucketprefix=dev resourceprefix=dev apidomain=${API_DOMAIN_STAGE} is_production=false

__deploy:
	gcloud builds submit \
	  --default-buckets-behavior="regional-user-owned-bucket" \
	  --region=asia-northeast1 \
	  --config=_cloudbuild.all.yaml \
	  --project ${project} \
	  --substitutions=_BUCKET_PREFIX=${bucketprefix},_RESOURCE_PREFIX=${resourceprefix},_API_DOMAIN=${apidomain},_IS_PRODUCTION=${is_production}

# チェックサム作成
migrate_hash:
	cd ../sql && atlas migrate hash

# マイグレーション実行
migrate_apply:
	$(LOAD_ENV) ; \
	cd ../sql && atlas migrate apply --env local

# テストデータ注入
db_seed:
	$(LOAD_ENV) ; \
	cat ../sql/seeds/seed_posts.sql | docker compose -f ../compose.yml exec -T postgres psql -U $${POSTGRES_USER} -d $${POSTGRES_DB}

# マイグレーション適用とテストデータ注入を実行
db_setup: migrate_apply db_seed

# ローカル環境のDBをクリーンアップして再マイグレーション（空のDBを作成）
migrate_fresh:
	$(LOAD_ENV) ; \
	cd ../sql && atlas schema clean --env local --auto-approve ; \
	docker compose -f ../compose.yml exec -T postgres psql -U $${POSTGRES_USER} -d $${POSTGRES_DB} -c "CREATE SCHEMA IF NOT EXISTS public;" ; \
	migrate_apply
