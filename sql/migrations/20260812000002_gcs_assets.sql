-- 静的画像はGCSのオブジェクトIDをDBに保持し、APIで公開URLへ変換する。
ALTER TABLE profiles
    ADD COLUMN IF NOT EXISTS avatar_url TEXT NOT NULL DEFAULT 'profile/profile.jpg';

UPDATE profiles
SET avatar_url = 'profile/profile.jpg';

UPDATE apps
SET icon_url = CASE icon_url
    WHEN '/whichway-icon.svg' THEN 'apps/whichway-icon.svg'
    WHEN '/mahjong-icon.svg' THEN 'apps/mahjong-icon.svg'
    WHEN '/pratan-icon.svg' THEN 'apps/pratan-icon.svg'
    WHEN '/economeye-icon.svg' THEN 'apps/economeye-icon.svg'
    ELSE icon_url
END;
