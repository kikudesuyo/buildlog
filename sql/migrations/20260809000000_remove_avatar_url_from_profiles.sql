-- プロフィール画像はWebアプリケーションの静的ファイルで管理するため、外部URLの保持をやめる。
ALTER TABLE profiles
    DROP COLUMN avatar_url;
