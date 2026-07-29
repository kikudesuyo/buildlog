# Buildlog Issue Roadmap

更新日: 2026-07-29

GitHub の Open Issue 45件を、依存関係と実装順に整理したロードマップです。各行のIssueを実装タスクとして扱い、同一内容の重複Issueは作成しません。`issue/34` は他agentの変更が作業ツリーに存在するため、最初にレビューして引き継ぎます。

## 実装順

### Phase 0: 既存実装のレビュー・基盤

- [x] #34 [syntax highlighting](https://github.com/kikudesuyo/buildlog/issues/34) — コードブロック表示・言語ラベル・コピー操作・ダークテーマを実装。⏳ PR Created (#90)
- [x] #56 [syntax highlighting for technology articles](https://github.com/kikudesuyo/buildlog/issues/56) — #34と重複。PR #90に統合済みとしてPASS（Issueへのクローズ操作はGitHub連携権限不足）
- [x] #62 [アーキテクチャ修正](https://github.com/kikudesuyo/buildlog/issues/62) — handler/service/routesの責務整理。⏳ PR Created (#91)
- [x] #61 [Go命名修正](https://github.com/kikudesuyo/buildlog/issues/61) — 命名規約と既存コードの整合。⏳ PR Created (#92)
- [x] #5 [認可の設定](https://github.com/kikudesuyo/buildlog/issues/5) — 公開投稿権限といいね/コメント権限の整理。⏳ PR Created (#93)

### Phase 1: 公開コンテンツ配信

- [x] #35 [RSS feed](https://github.com/kikudesuyo/buildlog/issues/35) — RSS/Atom生成とalternate link。⏳ PR Created (#94)
- [x] #36 [rich link preview cards](https://github.com/kikudesuyo/buildlog/issues/36) — OGP取得と安全なカード表示。⏳ PR Created (#95)
- [x] #57 [rich link preview cards v2](https://github.com/kikudesuyo/buildlog/issues/57) — #36と重複、同PRで対応済
- [x] #1 [Qiita連携](https://github.com/kikudesuyo/buildlog/issues/1) — 本文・認証・同期方向・重複処理の仕様が空欄のためPASS
- [x] #28 [Cloudflare R2画像アップロード](https://github.com/kikudesuyo/buildlog/issues/28) — R2バケット/公開ドメイン/認証情報が必要なためPASS
- [ ] #40 [favicon/PWA icon](https://github.com/kikudesuyo/buildlog/issues/40) — 既存faviconがSvelteロゴのためBuildlog用に置換し、manifest/iconを整備

### Phase 2: モバイル公開UI

- [ ] #53 [mobile header](https://github.com/kikudesuyo/buildlog/issues/53) — モバイルヘッダーの重なり/overflow修正
- [ ] #66 [折りたたみナビ](https://github.com/kikudesuyo/buildlog/issues/66) — ドロワー導入
- [ ] #67 [safe area/touch target](https://github.com/kikudesuyo/buildlog/issues/67) — 44px操作領域・safe area
- [ ] #68 [検索モーダル](https://github.com/kikudesuyo/buildlog/issues/68) — フォーカス・キーボード・閉じる操作
- [ ] #69 [テーマ切替](https://github.com/kikudesuyo/buildlog/issues/69) — 状態表示・初期ちらつき
- [ ] #41 [responsive typography](https://github.com/kikudesuyo/buildlog/issues/41) — 文字サイズ/ヘッダー総合調整
- [ ] #54 [category chips](https://github.com/kikudesuyo/buildlog/issues/54) — 横スクロールフィルタ
- [ ] #74 [category chip accessibility](https://github.com/kikudesuyo/buildlog/issues/74) — 選択状態・自動可視化・通知
- [ ] #55 [sticky footer](https://github.com/kikudesuyo/buildlog/issues/55) — 短いページのfooter位置
- [ ] #83 [footer/RSS links](https://github.com/kikudesuyo/buildlog/issues/83) — リンク折返しとRSS導線
- [ ] #63 [optimistic like](https://github.com/kikudesuyo/buildlog/issues/63) — いいね即時反映
- [ ] #73 [like rollback](https://github.com/kikudesuyo/buildlog/issues/73) — 失敗時ロールバック/再試行
- [ ] #70 [Diary long cards](https://github.com/kikudesuyo/buildlog/issues/70) — 長文カード折りたたみ
- [ ] #71 [Diary admin actions](https://github.com/kikudesuyo/buildlog/issues/71) — 編集/削除ラベル・確認
- [ ] #72 [Diary pagination states](https://github.com/kikudesuyo/buildlog/issues/72) — loading/error/end状態
- [ ] #75 [Tech empty states](https://github.com/kikudesuyo/buildlog/issues/75) — 空状態/再試行
- [ ] #76 [Tech hierarchy](https://github.com/kikudesuyo/buildlog/issues/76) — Featured/通常記事の階層
- [ ] #77 [Tech detail typography](https://github.com/kikudesuyo/buildlog/issues/77) — 詳細タイトル/本文のレスポンシブ
- [ ] #78 [Tech Markdown/code UX](https://github.com/kikudesuyo/buildlog/issues/78) — Markdown/コード/コピー
- [ ] #79 [Profile navigation](https://github.com/kikudesuyo/buildlog/issues/79) — 目次・画像フォールバック
- [ ] #80 [Profile contact](https://github.com/kikudesuyo/buildlog/issues/80) — mailto/コピー明確化
- [ ] #81 [Apps cards](https://github.com/kikudesuyo/buildlog/issues/81) — 画像/説明/外部リンク整理
- [ ] #82 [Apps states](https://github.com/kikudesuyo/buildlog/issues/82) — loading/画像失敗/空状態

### Phase 3: 管理画面・執筆体験

- [ ] #42 [admin UI](https://github.com/kikudesuyo/buildlog/issues/42) — エディタ全体スクロール/スマホ文字サイズ/設定位置
- [ ] #84 [admin drawer](https://github.com/kikudesuyo/buildlog/issues/84) — フォーカス・戻る・背面スクロール
- [ ] #85 [editor action bar](https://github.com/kikudesuyo/buildlog/issues/85) — 固定保存アクション
- [ ] #86 [editor title](https://github.com/kikudesuyo/buildlog/issues/86) — 可変サイズ・文字数
- [ ] #87 [editor sections](https://github.com/kikudesuyo/buildlog/issues/87) — 本文/設定の移動
- [ ] #88 [Markdown toolbar](https://github.com/kikudesuyo/buildlog/issues/88) — textarea連動
- [ ] #89 [inline tags](https://github.com/kikudesuyo/buildlog/issues/89) — prompt廃止

### Phase 4: 管理機能・可視化・運用

- [ ] #32 [投稿履歴カレンダー](https://github.com/kikudesuyo/buildlog/issues/32) — 投稿履歴の可視化
- [ ] #33 [analytics dashboard](https://github.com/kikudesuyo/buildlog/issues/33) — 集計APIと管理ダッシュボード
- [x] #17 [custom domain purchase](https://github.com/kikudesuyo/buildlog/issues/17) — 購入対象ドメイン・決済権限が必要なためPASS
- [x] #18 [custom domain configuration](https://github.com/kikudesuyo/buildlog/issues/18) — DNS/Cloud Run/Pagesのアカウント権限と購入ドメインが必要なためPASS

## 実装ルール

1. 1 Issueにつき1ブランチ `issue/{number}`、1 Pull Requestで対応する。
2. UI変更は390pxを基準に、変更前/変更後のスクリーンショットをPR descriptionに添付する。
3. 操作を伴う変更は、操作前後の動画をPR descriptionに添付する。
4. 既存の未コミット変更は巻き戻さず、重複実装を避ける。
5. 仕様が決めきれないIssue（Qiita連携、認可、ドメイン購入など）は、実装せずpassとして理由を記録する。
6. 各Issueでformat、lint、test、差分確認、セルフレビューを実行する。
7. PR descriptionに `Closes #123` / `Fixes #123` / `Resolves #123` を記載し、マージ時にGitHub標準機能で対象Issueを自動クローズする。
