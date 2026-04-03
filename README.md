# Digi-Engawa（デジタル縁側）

[![CI](https://github.com/akaitigo/digi-engawa/actions/workflows/ci.yml/badge.svg)](https://github.com/akaitigo/digi-engawa/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

公民館・地域サロンで開催する「デジタル縁側」教室の運営を支援するプラットフォーム。

大きなフォント・やさしい日本語・音声ガイド付きのステップバイステップ教材を提供し、受講者のつまずきをサポーター側にリアルタイム可視化するヘルプボードを搭載。

## 技術スタック

| レイヤー | 技術 |
|---------|------|
| フロントエンド | TypeScript / Next.js 15 (PWA) |
| バックエンド | Go 1.24 (net/http) |
| リアルタイム | WebSocket |
| 音声 | Web Speech API |
| データ | JSON ファイルストア（MVP） |

## アーキテクチャ

```
web/          Next.js PWA (TypeScript)
├── src/app/          App Router ページ
├── src/components/   UIコンポーネント
└── src/types/        型定義

api/          Go REST + WebSocket API
├── cmd/server/       エントリーポイント
├── internal/handler/ HTTPハンドラ
├── internal/model/   データモデル
├── internal/repository/ データ永続化
├── internal/service/ ビジネスロジック
└── internal/ws/      WebSocket Hub

db/migrations/  PostgreSQL マイグレーション（将来用）
docs/adr/       Architecture Decision Records
```

## MVP機能

1. **教材ビューア** — やさしい日本語＋ふりがな＋音声読み上げ付きステップバイステップ表示
2. **困ったボタン＋ヘルプボード** — 受講者の🆘をサポーターにリアルタイム通知（WebSocket）
3. **教室管理ダッシュボード** — 教室の作成・参加者管理・教室コードによるアクセス
4. **進捗トラッキング** — 受講者のステップ進捗をサポーターに可視化

## クイックスタート

```bash
# リポジトリをクローン
git clone git@github.com:akaitigo/digi-engawa.git
cd digi-engawa

# 環境変数を設定
cp .env.example .env

# 依存関係をインストール
make install

# フロントエンド開発サーバー
cd web && npm run dev

# バックエンドサーバー（別ターミナル）
cd api && go run ./cmd/server
```

## 開発

```bash
# 全チェック（format → lint → test → build）
make check

# フロントエンドのみ
cd web && npm run dev         # 開発サーバー
cd web && npm test            # テスト
cd web && npx biome check .   # lint

# バックエンドのみ
cd api && go run ./cmd/server # 起動
cd api && go test ./...       # テスト
```

## アクセシビリティ

- フォントサイズ最小18px
- タッチターゲット最小48px
- ARIA属性による進捗バー・ボタンのアクセシビリティ
- やさしい日本語（ふりがな付き）
- 音声読み上げ（Web Speech API、速度0.8x）

## デモ

<!-- デモGIF/スクリーンショットをここに配置 -->

## ライセンス

[MIT](LICENSE)
