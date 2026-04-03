# Digi-Engawa

## コマンド
- ビルド: `make build`
- テスト: `make test`
- lint: `make lint`
- フォーマット: `make format`
- 全チェック: `make check`
- フロントエンド単体: `cd web && npm run dev`
- バックエンド単体: `cd api && go run ./cmd/server`

## ワークフロー
1. research.md を作成（調査結果の記録）
2. plan.md を作成（実装計画。人間承認まで実装禁止）
3. 承認後に実装開始。plan.md のtodoを進捗管理に使用

## ルール
- ADR: docs/adr/ 参照。新規決定はADRを書いてから実装
- テスト: 機能追加時は必ずテストを同時に書く
- lint設定の変更禁止（ADR必須）

## 構造
- web/ — Next.js PWA（TypeScript）
- api/ — Go REST + WebSocket API
- db/migrations/ — PostgreSQLマイグレーション

## 禁止事項
- any型(TS) / console.log / print文のコミット
- TODO コメントのコミット（Issue化すること）
- .env・credentials のコミット
- lint設定の無効化

## 状態管理
- git log + GitHub Issues でセッション間の状態を管理
- セッション開始: `bash .claude/startup.sh`
