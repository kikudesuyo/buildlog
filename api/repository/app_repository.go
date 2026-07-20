package repository

import "github.com/kikudesuyo/buildlog/api/entity"

func ListAppProjects() []entity.AppProject {
	return []entity.AppProject{
		{ID: "focus-canvas", Name: "Focus Canvas", Category: "Productivity / Editorial", Tags: []string{"React", "Three.js"}, Description: "集中力を高めるために設計された執筆環境。環境音とタイポグラフィが調和し、ユーザーを深いフロー状態へと導きます。", Icon: "edit_note", DemoURL: "#", CodeURL: "#"},
		{ID: "chroma-pulse", Name: "Chroma Pulse", Category: "Data Viz / Audio", Tags: []string{"Web Audio", "Canvas"}, Description: "音声波形を美しい幾何学的なビジュアルへと変換する実験。リアルタイムでのオーディオ解析と、それを補完する滑らかなアニメーション。", Icon: "monitoring", DemoURL: "#", CodeURL: "#"},
		{ID: "zenith-shell", Name: "Zenith Shell", Category: "Utility / CLI", Tags: []string{"Rust", "WASM"}, Description: "モダンなワークフローに最適化されたコマンドライン・ユーティリティ。Rustによる高速な処理と、シンプルかつ強力なインターフェースを提供します。", Icon: "terminal", DemoURL: "#", CodeURL: "#"},
	}
}
