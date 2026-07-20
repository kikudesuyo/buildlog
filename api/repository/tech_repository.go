package repository

import "github.com/kikudesuyo/buildlog/api/entity"

func GetTechFeed() entity.TechFeed {
	return entity.TechFeed{
		FeaturedArticle: entity.FeaturedTechArticle{Title: "静かなるシステムの構築：Rustによるメモリ安全性の再考", Excerpt: "現代のソフトウェア開発において、安全性はもはやオプションではありません。所有権モデルがもたらす新しい秩序と、開発者の認知負荷を軽減するための抽象化について考察します。", Category: "Architecture", ReadTime: "12分で読める", Date: "2024.11.10"},
		Articles: []entity.TechArticle{
			{ID: "tech-1", Title: "インタフェースの沈黙：ミニマリズムUIの実装戦略", Excerpt: "情報を削ぎ落とすことで、ユーザーの集中力を最大化する。CSS Container Queriesを活用した、文脈に応じた適応型レイアウトの設計。", Category: "Development", ReadTime: "8 min read", Date: "2024.11.02", Views: "856 views"},
			{ID: "tech-2", Title: "生成AI時代のプロダクト倫理：透明性の設計", Excerpt: "LLMをプロダクトに組み込む際、どのようにして人間中心の設計を維持するか。データセットの偏りとアルゴリズムの透明性に関する実務的アプローチ。", Category: "Data Science", ReadTime: "15 min read", Date: "2024.10.28", Views: "2,105 views"},
			{ID: "tech-3", Title: "今週のライブラリ選定：Headless UIとアクセシビリティの追求", Excerpt: "スタイリングを強制しないコンポーネントが、いかにして長期的なメンテナンス性を向上させるか。Radix UIとTailwindの組み合わせ事例を詳解。", Category: "Newsletter", ReadTime: "5 min read", Date: "2024.10.20", Views: "542 views", IsNewsletter: true},
		},
	}
}
