# Buildlog Agent Notes

このファイルは、Buildlogでの作業中に追加された運用ルールと判断メモを蓄積する。

## Issue / PR運用

- `roadmap.md` を作業順の正本として扱い、未完了Issueを1件ずつ進める。
- 1 Issueにつき `issue/{number}` ブランチと1 Pull Requestを作る。
- PR descriptionには対応Issueを `Closes #123`、`Fixes #123`、`Resolves #123` のいずれかで明記する。マージ時のGitHub標準機能による自動クローズを利用する。
- マージ済みPRに対応するIssueは、対応済みとしてroadmapに反映する。
- 仕様・外部アカウント・決済・DNS・Cloudflareリソースなど、意思決定や権限が必要なものは実装せずPASS理由をroadmapに記録する。

## UIレビュー

- UI変更では390px幅を基準にbefore/afterスクリーンショットを作成し、PR descriptionに添付する。
- 操作を伴う変更では操作前後の動画もPR descriptionに添付する。
- モバイル操作は44px以上のタッチ領域、safe area、キーボード表示時のフォーカス、横スクロール、固定要素との重なりを確認する。

## 実装・検証

- 作業開始前にdev-platformのルール、対象リポジトリのREADME、Issue、既存PRレビューを確認する。
- 他agentの変更を巻き戻さず、ローカル作業ツリーとリモートPRの差分を確認してから編集する。
- `.pnpm-store` や秘密情報をコミットしない。
- 各Issueでformat、lint、test、diff checkを実行し、セルフレビューで指摘がなくなるまで修正する。
- GitHub連携の書き込み権限がない場合は、実施できなかった操作を明記し、勝手にIssue状態を推測しない。
