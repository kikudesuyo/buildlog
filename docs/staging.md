# Staging環境

## GCP Project

表示名は `buildlog_dev` として構いませんが、GCPのProject IDにはアンダースコアを使用できません。Project IDは `buildlog-dev` を使用します。

このリポジトリでは、Stagingの既定値を `buildlog-dev` としています。既に取得済みの別IDを使う場合は、GitHub Repository Variablesの `GCP_PROJECT_STAGING` で上書きしてください。

## 必須リソース

- Cloud Run service: `buildlog-api-staging`
- Artifact Registry image: `asia-northeast1-docker.pkg.dev/buildlog-dev/buildlog/api-staging`
- Secret Manager secrets:
  - `STAGING_DATABASE_URL`
  - `STAGING_ADMIN_PASSWORD`
  - `STAGING_ADMIN_SESSION_SECRET`
- GitHub Actions用Workload Identity Federation provider
- GitHub Actions用サービスアカウント

StagingのCloud Runには `STAGING_*` prefix付きの環境変数をSecret Managerから注入します。`IS_PRODUCTION=false` の場合、APIはProduction用の環境変数へfallbackしません。Secretの値や接続文字列をGitHub Actionsのログへ出力しないでください。

## Neon

Neonに `staging` ブランチを作成し、その接続文字列をStaging ProjectのSecret Managerにある `STAGING_DATABASE_URL` へ登録します。本番データを含めたくない場合は、migrationとテスト用seedからStaging DBを作成してください。

Staging DBへmigrationを適用する場合は、GitHub Actionsの `Migrate Staging Database` を手動実行します。

## GitHub Repository Variables

Staging用のGCPプロジェクト、WIF Provider、サービスアカウントはWorkflowに固定値として定義しています。これらは秘密情報ではなく、GitHub ActionsはWIF経由でGCPへ認証します。追加のRepository Variablesは不要です。

Actions の `Merge branch to Staging` を手動実行し、workflow のブランチ選択でマージ元を選び、`deploy` を選択します。選択したブランチを `staging` へマージし、`deploy=true` の場合だけ、マージ後の `staging` ブランチをデプロイします。`staging` へのpushだけではデプロイしません。
