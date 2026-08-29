-- プロフィールのサブタイトルは管理画面・公開画面ともに使用しないため削除する。
ALTER TABLE profiles
DROP COLUMN IF EXISTS subtitle;
