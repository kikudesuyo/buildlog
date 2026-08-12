# AIエージェントへの指示

## 最重要: 作業完了の定義

Issue・タスクを実装しただけでは作業完了ではありません。

特別な指示がない限り、以下の3つのゲートをすべて通過して初めて作業完了としてください。

### 1. Acceptance Gate

Issue・ユーザー指示の要求を実際に満たしていること。

### 2. Quality Gate

テスト・静的解析・自己レビューを通過し、未解決の問題がないこと。

### 3. Delivery Gate

commit、push、Pull Request作成、必要な検証結果・UI証跡の添付まで完了していること。

**実装をレビューするだけではなく、Issueが本当に解決したことをレビューしてください。**

コード変更、commit、pushだけで作業を終了してはいけません。
特別な指示がない限り、必ずPull Request作成まで行ってください。

---

# 基本作業フロー

Issueごとに、以下の順序で作業してください。

1. 最新の `main` を取得する
2. Issue・ユーザー指示を確認する
3. Acceptance Criteriaを整理する
4. 各Acceptance Criteriaの検証方法を決める
5. 作業branchを作成する
6. 実装する
7. format・lint・test等を実行する
8. Acceptance Criteriaを実際の出力・挙動で検証する
9. 自己レビューする
10. 問題があれば修正し、7〜9を繰り返す
11. UI変更の場合はスクリーンショット・必要に応じて動画を取得する
12. UI証跡そのものを確認し、Acceptance Criteriaを再検証する
13. commitする
14. remoteへpushする
15. Pull Requestを作成する
16. PR descriptionへ検証結果・Acceptance Criteria・UI証跡を記載する
17. 作成したPRをGitHub上で確認する
18. PR URLを取得して最終報告する

1つでも未完了の場合、作業を「完了」と報告してはいけません。

---

# Acceptance Gate

## Issue要件の整理

実装前に、Issue本文・ユーザー指示から達成すべき条件をAcceptance Criteriaとして具体化してください。

例えば、

> 目標を追加ボタンと同じサイズにする

というIssueであれば、

- モバイルで2つのボタンの最終的な表示サイズが同じ
- PCでも意図しないレイアウト崩れがない

など、実際に検証可能な条件へ落とし込んでください。

コードの変更内容をAcceptance Criteriaにしてはいけません。

悪い例:

- `px-4` に変更する
- `min-h-11` を追加する

良い例:

- ブラウザ上で比較対象と同じ高さになっている
- ブラウザ上で比較対象と同じ幅になっている

---

## Acceptance Criteriaの検証

実装後は、各Acceptance Criteriaを必ず `PASS` / `FAIL` で判定してください。

コードを読んだだけで要求を満たしたと判断してはいけません。

可能な限り、以下のような実際の結果を使用してください。

- ブラウザ上のレンダリング結果
- Playwright等によるDOM・寸法の取得
- APIレスポンス
- DBの実データ
- テスト結果
- CLIの出力

1つでも `FAIL` がある場合は、実装へ戻ってください。

すべて `PASS` になるまでPR作成完了・作業完了としてはいけません。

---

## 比較対象が指定されている場合

Issueに以下のような要求がある場合は、比較対象との**最終結果を直接比較**してください。

- 「Aと同じサイズ」
- 「Aと同じ見た目」
- 「Aと同じ挙動」
- 「Aに合わせる」
- 「以前と同じ」
- 「デザイン通り」
- 「○○のようにする」

CSS class、padding、height、width、propsなど、一部の実装値が一致したことだけを根拠に完了としてはいけません。

最終的なレンダリング結果・実寸・挙動を比較してください。

可能な場合は目視だけではなく、Playwright等で機械的に検証してください。

例:

- サイズ: `boundingBox()` / `getBoundingClientRect()`
- 表示状態: locator assertion
- 文言: text assertion
- URL遷移: URL assertion
- API: status code / response body

---

# Quality Gate

## 必須検証

変更内容に応じて、必要なformat・lint・test・build等を実行してください。

最低限、以下を確認してください。

```bash
git diff --check
```

加えてプロジェクトで利用可能な、

- format
- lint
- type check
- unit test
- integration test
- build

などを実行してください。

テストが成功しただけではQuality Gate通過とはみなしません。

---

# 自己レビュー

実装・検証後、変更内容を実装者自身でレビューしてください。

少なくとも以下を確認してください。

## Issueとの整合性

- Issueの要求をすべて満たしているか
- 要求していない変更を含めていないか
- Acceptance Criteriaがすべて `PASS` か
- 「実装した」ことではなく「問題が解決した」ことを確認したか

## 差分

必ず以下を確認してください。

```bash
git diff main...HEAD
git diff --check
```

差分を最初から最後まで読み、

- 意図しない変更
- デバッグコード
- 不要なコメント
- 不要な依存追加
- 関係ないファイル変更

がないことを確認してください。

## 既存コード規則

変更対象だけでなく周辺コードも確認し、

- ディレクトリ構成
- レイヤー責務
- 命名
- エラーハンドリング
- 共通コンポーネント
- 既存の実装パターン

に沿っていることを確認してください。

既存実装と異なる方式を採用する場合は、その理由をPR descriptionに記載してください。

## セキュリティ

以下を確認してください。

- 認証・認可漏れ
- 入力値検証不足
- SQLインジェクション
- XSS
- CSRF
- SSRF
- 機密情報のログ出力
- 機密情報のハードコード
- 危険なファイル操作
- 不要な依存関係追加

ユーザー入力・外部データ・権限境界を変更した場合は、正常系だけでなく拒否されるべきケースも確認してください。

## 影響範囲

変更ファイルだけでなく、

- API利用箇所
- DB
- migration
- UI状態
- 既存テスト
- 呼び出し元
- 呼び出し先

への影響も確認してください。

---

# 実装・検証・レビューサイクル

AI Agentは一度実装して終了せず、以下のサイクルを自律的に繰り返してください。

```text
要件整理
↓
実装
↓
技術検証
↓
Acceptance Criteria検証
↓
自己レビュー
↓
判定
```

判定結果が以下をすべて満たした場合のみ `APPROVED` としてください。

- Acceptance Criteriaがすべて `PASS`
- 必要なテスト・静的解析が成功
- UI変更の場合、実際のブラウザ確認が完了
- 自己レビューに未解決の指摘がない
- セキュリティ上の懸念がない
- 意図しない差分がない

問題が1つでもある場合は実装へ戻ってください。

修正後は、修正した箇所だけでなくAcceptance Criteria全体を再検証してください。

同じ問題を3サイクル連続で解消できない場合は、無理に完了扱いせず、

- 原因
- 試した対応
- 現在の状態
- 必要なユーザー判断

を報告してください。

---

# UI変更時の必須確認

UI変更がある場合、コード・lint・testだけで確認済みとしてはいけません。

必ず実際のブラウザを使用してください。

## ローカル起動

API:

```bash
cd api
make dev
```

Web:

```bash
cd web
pnpm run dev
```

利用可能なブラウザ操作ツールを使用して対象画面を開いてください。

---

## 必須viewport

原則として以下の両方を確認してください。

Mobile:

```text
375px × 812px
```

PC:

```text
1280px × 800px
```

異なるviewportを使用した場合は、実際に使用したサイズをPR descriptionへ記載してください。

単に「モバイル」「PC」とだけ書いてはいけません。

---

## 必須スクリーンショット

UI変更がある場合、少なくとも以下を撮影してください。

- Mobile
- PC

変更したすべての主要ページ・主要状態を確認対象にしてください。

単一ページの変更でも、原則としてMobile・PCの両方を撮影してください。

---

# スクリーンショット取得後の必須レビュー

**スクリーンショットを撮影しただけではUI確認完了ではありません。**

撮影後、必ずスクリーンショット自体を確認してください。

IssueのAcceptance Criteriaと画像を比較し、以下を確認してください。

- 修正対象が期待通り変わっている
- Issueの要求を実際に満たしている
- 比較対象がある場合、画像上でも比較している
- サイズが正しい
- 位置が正しい
- 余白が正しい
- レイアウト崩れがない
- Mobile / PCそれぞれで問題がない
- 無関係な箇所に意図しない変化がない

**Issueを満たしていないことがスクリーンショットから確認できる場合、その画像をそのままPRへ貼って完了扱いにしてはいけません。**

実装へ戻り、

```text
修正
↓
ブラウザ確認
↓
再撮影
↓
画像レビュー
↓
Acceptance Criteria再判定
```

を行ってください。

---

# 動画が必要な変更

以下のような操作を変更した場合は、スクリーンショットに加えて操作前後が確認できる動画を撮影してください。

- クリック
- 入力
- モーダル
- ドロワー
- ページ遷移
- アニメーション
- 追加読み込み
- ドラッグ&ドロップ
- その他インタラクション

単純な見た目だけの変更であれば動画は不要です。

---

# UI証跡のPR添付

スクリーンショット・動画は、GitHub上からレビュー担当者が参照できる状態にしてください。

ローカルファイルパスだけをPR descriptionへ記載してはいけません。

原則としてGitHubから直接参照可能なURLを使用してください。

アップロード機能が利用できない場合は、

```text
docs/screenshots/<issue-number>/
```

へ画像をcommit・pushし、GitHubの絶対URLを使用してください。

例:

```text
https://github.com/<owner>/<repo>/blob/<commit>/docs/screenshots/<issue-number>/mobile.png?raw=1
```

相対パスやローカルパスだけを使用してはいけません。

画像自体をリサイズせず、PR descriptionではHTMLの `width` 属性を使用してください。

Mobile:

```html
<img src="IMAGE_URL" alt="モバイル表示" width="300" />
```

PC:

```html
<img src="IMAGE_URL" alt="PC表示" width="400" />
```

PR作成後、GitHub上で実際に画像・動画が表示・参照できることを確認してください。

確認完了まで一時ファイルを削除してはいけません。

---

# 作業開始時のGit操作

複数Agent・複数ユーザーが同時に開発するため、古いbranchを起点に作業してはいけません。

作業開始時は必ず最新の `main` を取得してください。またworktreeを作成して着手してください。

```bash
git switch main
git pull --ff-only
git switch -c issue/<issue-number>
```

既存branchを利用する場合も、作業開始前にremoteの最新状態を確認してください。

---

# マルチエージェント・並行作業

別Agentによってbranchや作業ツリーが変更されている場合があります。

branchが変わったことだけを理由に作業を停止してはいけません。

状態が想定と異なる場合は、まず以下を確認してください。

```bash
git status --short --branch
git branch -vv
git worktree list
git log --oneline --decorate -10
```

その上で、

- 自分の作業branchが存在する場合は戻って継続する
- branchが失われている場合は最新mainからbranchを作り直す
- 必要に応じて既存commit・差分を移す
- 他Agentの変更を勝手に削除しない
- 他Agentの未commit差分をresetしない
- 他Agentのbranchを上書きしない
- 差分が混在する場合はworktreeまたは別branchへ分離する

安全に分離できない場合のみ、状況を具体的に報告してください。

---

# IssueとPull Requestの対応単位

IssueとPull Requestは原則1対1にしてください。

複数Issueをまとめて依頼された場合でも、Issue番号ごとに、

```text
issue/<number>
```

branchを作成し、個別のPRを作成してください。

1つのbranch・PRに複数Issueの変更を混在させてはいけません。

Issueごとに、

```text
実装
↓
検証
↓
Acceptance確認
↓
レビュー
↓
commit
↓
push
↓
PR
```

まで完了してから次のIssueへ進んでください。

依存関係がある場合は、先行PRをPR descriptionへ明記してください。

---

# commit / push

Quality GateとAcceptance Gateを通過した後にcommitしてください。

その後remoteへpushしてください。

未解決の問題がある状態を「完成版」としてcommit・pushしてはいけません。

---

# Pull Request作成

特別な指示がない限り、実装・検証後は必ずPull Requestを作成してください。

Draftではなく、レビュー可能な状態で作成してください。

```bash
gh pr create --body-file /tmp/pr-body.md
```

`--draft` は使用しないでください。

PR descriptionを直接長いコマンドライン引数へ埋め込まず、Markdownファイルを作成して `--body-file` を使用してください。

---

# Issueの自動クローズ

PR descriptionには必ず以下のいずれかを記載してください。

```text
Closes #123
```

```text
Fixes #123
```

```text
Resolves #123
```

Issue番号を推測してはいけません。

PR作成直後にIssueを手動で閉じてはいけません。
IssueはPRマージ時に自動クローズさせてください。

---

# PR description

最低限、以下を記載してください。

```md
## 概要

変更内容を簡潔に記載する。

Closes #123

## Acceptance Criteria

- [x] 条件1
  - 検証方法:
  - 結果: PASS

- [x] 条件2
  - 検証方法:
  - 結果: PASS

## 変更内容

- xxx
- xxx

## 動作確認

- `command`: 成功
- `command`: 成功

## UI確認

### Mobile

Viewport: `375px × 812px`

<img src="..." alt="モバイル表示" width="300">

### PC

Viewport: `1280px × 800px`

<img src="..." alt="PC表示" width="400">

### 動画

必要な場合のみ記載。

## 自己レビュー

- Issue要件: PASS
- 差分確認: PASS
- 既存コード規則: PASS
- セキュリティ: PASS
- 影響範囲: PASS
- 未解決事項: なし

## 補足

必要な場合のみ記載。
```

UI変更がない場合は、

```text
UI変更なし
```

と明記してください。

---

# PR作成後の確認

PRを作成しただけではDelivery Gate通過ではありません。

`gh pr view` 等を使用して、実際のPRを確認してください。

最低限以下を確認します。

- PRが作成されている
- PR URLを取得できる
- titleが正しい
- descriptionが登録されている
- Markdownの改行が壊れていない
- `Closes #...` が記載されている
- Acceptance Criteriaが記載されている
- 検証結果が記載されている
- UI変更の場合、Mobile / PC画像がある
- 必要な場合、動画がある
- 画像・動画URLが `https://` で参照可能
- GitHub上で画像・動画が実際に表示できる

**PR URLを取得するまでPR作成完了としてはいけません。**

---

# migration変更時の検証

PR作成前の確認で本番DBへmigrationを適用してはいけません。

migrationを変更した場合は最低限以下を実行してください。

```bash
cd sql && atlas migrate validate --dir "file://migrations"
```

```bash
cd api && GOCACHE=/tmp/buildlog-go-cache go test ./...
```

```bash
git diff --check
```

実行コマンドと結果をPR descriptionへ記載してください。

既知warningや失敗を隠してはいけません。

本番migrationは承認済みのデプロイ・運用手順、またはGitHub Actionsからのみ実行してください。

---

# UI変更時の完了チェック

UI変更がある場合、以下をすべて確認してください。

- [ ] IssueからAcceptance Criteriaを抽出した
- [ ] Acceptance Criteriaの検証方法を決めた
- [ ] Mobileで対象画面を確認した
- [ ] PCで対象画面を確認した
- [ ] Mobileスクリーンショットを撮影した
- [ ] PCスクリーンショットを撮影した
- [ ] スクリーンショット自体をレビューした
- [ ] 比較対象がある場合は実際に比較した
- [ ] Acceptance CriteriaがすべてPASS
- [ ] 操作変更がある場合は動画を撮影した
- [ ] UI証跡をGitHubから参照可能にした
- [ ] PR descriptionへUI証跡を追加した
- [ ] PRを作成した
- [ ] GitHub上でdescriptionを確認した
- [ ] GitHub上で画像・動画を確認した
- [ ] PR URLを取得した
- [ ] 必要なテスト結果を記載した
- [ ] 最終確認後に一時ファイルを削除した

1つでも未完了の場合、UI変更タスクを「完了」と報告してはいけません。

---

# 作業を停止してよい条件

以下の場合を除き、PR作成まで自律的に進めてください。

- 必要な認証情報がなく継続できない
- GitHubへのpush権限がない
- GitHubへのPR作成権限がない
- ブラウザを利用できない
- 必要なローカルサービスを起動できない
- UI証跡をGitHubから参照可能な状態にできない
- ユーザー判断なしでは安全に進められない重大な仕様の曖昧さがある

branchが別Agentによって変更されたことだけでは停止してはいけません。

停止する場合も、可能なところまで作業を進めた上で、

- 完了していること
- 未完了のこと
- 実行できない理由
- 試した対応
- ユーザー側で必要な対応

を具体的に報告してください。

「PRを作った方がよい」
「スクリーンショットを撮った方がよい」

など、残作業を提案するだけで終了してはいけません。

---

# GitHub文書

以下は日本語で記載してください。

- Issue
- Pull Request title
- Pull Request description
- GitHubコメント

長いMarkdown本文は `.md` ファイルを作成し、GitHub CLIの `--body-file` を利用してください。

---

# 最終報告

作業終了時には最低限以下を報告してください。

- 対応したIssue
- 実装内容
- Acceptance Criteriaと結果
- 実行したテスト・検証
- UI確認内容
- 自己レビュー結果
- commit
- PR URL

UI変更の場合は、

- Mobile screenshot添付済み
- PC screenshot添付済み
- 必要な場合は動画添付済み
- GitHub上で表示確認済み

であることも明記してください。

---

# 絶対にしてはいけないこと

以下は禁止します。

- 実装しただけで完了報告する
- testが通っただけでIssue解決と判断する
- CSS classや実装値だけを見てUI要件を満たしたと判断する
- スクリーンショットを撮影しただけでUI確認済みとする
- スクリーンショットに問題が見えているのに完了扱いする
- Acceptance CriteriaがFAILのままPRを完成扱いする
- 自己レビューを省略する
- PRを作成せず「作成した方がよい」と提案して終了する
- UI証跡が必要なのにローカルパスだけを記載する
- 他Agentの変更を勝手にreset・削除する
- Issue番号を推測する
- 未解決の重大な問題を隠す

**最終的な判断基準は「コードを書いたか」ではなく、「Issueで要求された状態が実際に実現されているか」です。**
