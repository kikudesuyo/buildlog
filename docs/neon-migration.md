# Neon DBリージョン移行

## 概要

Neonでは既存ProjectのRegionを変更できないため、移行先Regionに新しいProjectを作成し、旧ProjectのPostgreSQLデータを `pg_dump` / `pg_restore` で移行する。

Neon公式のリージョン移行手順でも、同じ方法で新しいProjectへデータをコピーし、アプリケーションの接続先を切り替える手順が案内されている。

## 移行手順

### 1. 移行先Projectを作成する

Neon Consoleで移行先Regionを指定して新しいProjectを作成する。PostgreSQLのメジャーバージョンは、原則として旧Projectと合わせる。

移行中に旧DBへ書き込まれたデータは、新DBへ自動では反映されない。データの取りこぼしを防ぐには、dumpから接続先切り替えまでの間をメンテナンス時間にするか、論理レプリケーションなど別の方法を検討する。

### 2. 旧DBをdumpする

`pg_dump` は **Direct connection** を使用する。NeonのPooled connectionでは、dump中に実行される `SET` の状態が保持されず、問題になる場合がある。

```bash
pg_dump \
  --format=custom \
  --no-owner \
  --no-privileges \
  --dbname="$OLD_DATABASE_URL" \
  --file=neon.dump
```

### 3. 新DBへrestoreする

```bash
pg_restore \
  --no-owner \
  --no-privileges \
  --clean \
  --if-exists \
  --dbname="$NEW_DATABASE_URL" \
  neon.dump
```

`--no-owner` と `--no-privileges` を指定することで、旧DB固有の所有者・権限に依存せずにrestoreする。`--clean --if-exists` は、移行先に同名オブジェクトがある場合に削除してから復元するための指定である。既存データを残したい場合は、実行前にオプションの意味と対象DBを確認する。

### 4. restore結果を確認する

テーブル一覧を確認する。

```bash
psql "$NEW_DATABASE_URL" -c '\dt public.*'
```

必要に応じて、件数や代表的なデータも確認する。

```bash
psql "$NEW_DATABASE_URL" -c 'SELECT COUNT(*) FROM public.posts;'
psql "$NEW_DATABASE_URL" -c 'SELECT * FROM public.posts LIMIT 10;'
```

### 5. アプリケーションの接続先を切り替える

新DBの接続文字列をアプリケーションの環境変数・Secret Managerなどへ反映し、アプリケーションから読み書きできることを確認する。問題がないことを確認してから旧Projectを削除する。

## はまったこと: `relation "posts" does not exist`

restore後に、次のクエリが失敗した。

```sql
SELECT * FROM posts;
```

エラーは次のとおり。

```text
relation "posts" does not exist
```

しかし、スキーマを明示するとクエリは成功した。

```sql
SELECT * FROM public.posts;
```

### 切り分け

まず、テーブルが存在するスキーマを確認する。

```sql
\dt public.*
```

次に、接続ごとの `search_path` を確認する。

```sql
SHOW search_path;
```

今回の結果は次のとおりだった。

| 接続方式          | `search_path`     |
| ----------------- | ----------------- |
| Pooled connection | 空                |
| Direct connection | `"$user", public` |

### 原因

NeonのPooled connectionを使用していた。NeonのPooled connectionは、接続先ホストのEndpoint IDに `-pooler` が付く。

```text
# Pooled
ep-xxx-pooler.ap-southeast-1.aws.neon.tech

# Direct
ep-xxx.ap-southeast-1.aws.neon.tech
```

Pooled connectionはPgBouncerのトランザクションモードで動作するため、`search_path` のようなセッション設定を後続トランザクションまで保持できない場合がある。その結果、`public` スキーマが検索対象にならず、テーブルが存在するにもかかわらず `relation does not exist` になる。

### 解決策

今回はDirect connectionへ切り替えて解決した。

用途に応じて、次のいずれかを選択する。

- `search_path` を保持したい、またはdump・restore・マイグレーションを実行する場合: Direct connection
- アプリケーションのクエリでスキーマを明示できる場合: `public.posts` のように完全修飾名を使用する
- ロール単位で恒久的に設定したい場合: `ALTER ROLE ... SET search_path ...` を検討する

## 覚えておくこと

`relation does not exist` が出ても、テーブルが本当に存在しないとは限らない。次の順番で確認する。

1. `\dt public.*` でテーブルの存在を確認する
2. `SHOW search_path;` で検索対象スキーマを確認する
3. `SELECT * FROM public.posts;` のようにスキーマを明示してクエリする
4. 接続先ホストに `-pooler` が付いていないか確認する

## 参考

- [Neon: How do I migrate an existing Neon project to a different AWS region?](https://neon.com/faqs/change-region-existing-neon-project)
- [Neon Docs: Migrate data from another Neon project](https://neon.com/docs/import/migrate-from-neon)
- [Neon Docs: Connection pooling](https://neon.com/docs/connect/connection-pooling)
