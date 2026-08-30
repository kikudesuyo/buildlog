# Gemini PR Reviewer

Pull Requestの差分とGoのテスト・vet・lint結果をGoogle Gemini Developer APIへ渡し、Structured Outputでレビュー結果を取得してPRコメントへ投稿します。85点未満の場合はworkflowを失敗させます。Vertex AIは使用しません。

## GitHub設定

GitHub ActionsのRepository Variablesに次を登録します。

- `GCP_REVIEW_PROJECT`: `GEMINI_API_KEY` Secretを保存するGoogle Cloudプロジェクト
- `GCP_REVIEW_WIF_PROVIDER`: このリポジトリから利用できるWorkload Identity Federation provider
- `GCP_REVIEW_SERVICE_ACCOUNT`: `GEMINI_API_KEY`だけに`roles/secretmanager.secretAccessor`を持つ専用サービスアカウント
- `GEMINI_MODEL`: 任意。未設定時は `gemini-2.5-flash`

必要なVariablesが未設定の間はReviewer jobをスキップします。設定後にPull Requestへpushするとレビューが実行されます。

Google Cloud Secret ManagerにGemini Developer APIの `GEMINI_API_KEY` を保存し、専用サービスアカウントのアクセスを対象Secretだけに制限してください。ActionsのOIDC信頼条件もこのリポジトリ・必要なブランチに限定します。CLIは `https://generativelanguage.googleapis.com/v1beta/models/<model>:generateContent` を `x-goog-api-key: $GEMINI_API_KEY` で呼び出します。

APIキーはGitHub Secretsへ登録せず、workflowのログにも出力しません。レビュー対象のGoコマンドはキー取得後に実行するため、workflowを変更できる権限のないForkからのPull Requestではジョブ自体を実行しません。Forkのレビューが必要な場合は、差分を信頼できるブランチへ取り込んでから実行する運用にしてください。

## 評価基準

100点満点で、Correctness 30 / Maintainability 20 / Test quality 20 / Security 10 / Performance 10 / Readability 10です。

- 85点以上: `PASS`
- 70〜84点: `NEEDS_IMPROVEMENT`
- 69点以下: `FAIL`

現段階ではコードの自動修正や再レビューのループは行いません。これらは別Issueで扱います。
