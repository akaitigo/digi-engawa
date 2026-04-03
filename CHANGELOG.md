# Changelog

## [1.0.0] - 2026-04-04

### Added
- **教材ビューア** (F1): やさしい日本語＋ふりがな＋音声読み上げ付きステップバイステップ教材表示
  - StepViewer, FuriganaText, SpeechButton, ProgressBar コンポーネント
  - Material CRUD API + Step 管理
  - Web Speech API による日本語音声読み上げ（速度0.8x）

- **困ったボタン＋ヘルプボード** (F2): 受講者の SOS をサポーターにリアルタイム通知
  - WebSocket Hub による教室単位のブロードキャスト
  - HelpRequest API（pending → in_progress → resolved ステータス遷移）
  - HelpButton, HelpBoard, HelpRequestCard コンポーネント

- **教室管理ダッシュボード** (F3): 教室の作成・参加者管理・教材配布
  - Classroom CRUD API + 6文字英数字の教室コード
  - Participant 管理（受講者・サポーター・主催者ロール）
  - ClassroomCard, ClassroomForm, ParticipantList コンポーネント
  - 教室コードによる参加機能

- **進捗トラッキング** (F4): 受講者のステップ進捗をサポーターに可視化
  - Progress API（upsert + 教室単位取得）
  - WebSocket による進捗リアルタイム更新
  - ProgressOverview コンポーネント

### Infrastructure
- Next.js 15 PWA フロントエンド + Go 1.24 バックエンド（モノレポ）
- GitHub Actions CI（lint, typecheck, test, build）
- Biome v2 linter/formatter + golangci-lint
- Lefthook による pre-commit hooks
- Dependabot 設定
- ADR-0001: モノレポ構成の決定
