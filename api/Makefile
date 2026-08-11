.PHONY: build dev dev-auth lint lint-fix docker-build-prod _docker-build deploy-prod deploy-staging __deploy migrate_apply_prod migrate_apply_staging

GO111MODULE=on
LOAD_ENV=. ./default_env.sh
APP_BINARY ?= app
GCP_PROJECT_PROD ?= buildlog-503612
GCP_PROJECT_STAGING ?= buildlog-dev
API_DOMAIN_PROD ?=
API_DOMAIN_STAGING ?=

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

docker-build-prod:
	make _docker-build project=${GCP_PROJECT_PROD}

_docker-build:
	gcloud builds submit --default-buckets-behavior="regional-user-owned-bucket" --region=asia-northeast1 --config=_cloudbuild.all.yaml --project ${project} .

ALLOWED_ORIGINS ?=

# 本番環境 Cloud Runへデプロイ
deploy-prod:
	make __deploy project=${GCP_PROJECT_PROD} bucketprefix=prod resourceprefix=prod apidomain=${API_DOMAIN_PROD} is_production=true allowed_origins="${ALLOWED_ORIGINS}" database_env=DATABASE_URL database_secret=DATABASE_URL password_env=ADMIN_PASSWORD password_secret=ADMIN_PASSWORD session_secret_env=ADMIN_SESSION_SECRET session_secret=ADMIN_SESSION_SECRET

# Staging環境のCloud Runへデプロイ
deploy-staging:
	@test -n "${GCP_PROJECT_STAGING}" || (echo "GCP_PROJECT_STAGING is required" && exit 1)
	make __deploy project=${GCP_PROJECT_STAGING} bucketprefix=staging resourceprefix=staging apidomain=${API_DOMAIN_STAGING} is_production=false allowed_origins="${ALLOWED_ORIGINS}" database_env=STAGING_DATABASE_URL database_secret=STAGING_DATABASE_URL password_env=STAGING_ADMIN_PASSWORD password_secret=STAGING_ADMIN_PASSWORD session_secret_env=STAGING_ADMIN_SESSION_SECRET session_secret=STAGING_ADMIN_SESSION_SECRET

__deploy:
	gcloud builds submit \
	  --default-buckets-behavior="regional-user-owned-bucket" \
	  --region=asia-northeast1 \
		--config=_cloudbuild.all.yaml \
		--project ${project} \
		--substitutions=_RESOURCE_PREFIX=${resourceprefix},_API_DOMAIN=${apidomain},_IS_PRODUCTION=${is_production},_ALLOWED_ORIGINS="${allowed_origins}",_DATABASE_URL_ENV=${database_env},_DATABASE_URL_SECRET=${database_secret},_ADMIN_PASSWORD_ENV=${password_env},_ADMIN_PASSWORD_SECRET=${password_secret},_ADMIN_SESSION_SECRET_ENV=${session_secret_env},_ADMIN_SESSION_SECRET_SECRET=${session_secret}

# チェックサム作成
migrate_hash:
	cd ../sql && atlas migrate hash

# マイグレーション実行
migrate_apply:
	$(LOAD_ENV) ; \
	cd ../sql && atlas migrate apply --env local

# 本番環境マイグレーション実行
migrate_apply_prod:
	@PROD_DATABASE_URL=$$(gcloud secrets versions access latest --secret=DATABASE_URL --project=${GCP_PROJECT_PROD}) ; \
	cd ../sql && PROD_DATABASE_URL="$$PROD_DATABASE_URL" atlas migrate apply --env prod

# Staging環境のマイグレーション実行
migrate_apply_staging:
	@STAGING_DATABASE_URL=$$(gcloud secrets versions access latest --secret=STAGING_DATABASE_URL --project=${GCP_PROJECT_STAGING}) ; \
	cd ../sql && STAGING_DATABASE_URL="$$STAGING_DATABASE_URL" atlas migrate apply --env staging

# テストデータ注入
db_seed:
	$(LOAD_ENV) ; \
	cat ../sql/seeds/seed_posts.sql ../sql/seeds/seed_comments.sql ../sql/seeds/seed_goals.sql | docker compose -f ../compose.yml exec -T postgres psql -U $${POSTGRES_USER} -d $${POSTGRES_DB}

# マイグレーション適用とテストデータ注入を実行
db_setup: migrate_apply db_seed

# ローカル環境のDBをクリーンアップして再マイグレーション（空のDBを作成）
migrate_fresh:
	$(LOAD_ENV) ; \
	cd ../sql && atlas schema clean --env local --auto-approve ; \
	docker compose -f ../compose.yml exec -T postgres psql -U $${POSTGRES_USER} -d $${POSTGRES_DB} -c "CREATE SCHEMA IF NOT EXISTS public;" ; \
	migrate_apply
