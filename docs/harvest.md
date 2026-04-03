# Harvest: digi-engawa

## メトリクス
| 項目 | 値 |
|------|-----|
| コミット数 | 8 |
| ADR数 | 1 |
| CLAUDE.md行数 | 35 |
| CI | YES (.github/workflows/ci.yml) |
| lefthook.yml | YES |
| Go パッケージ数 | 8 (handler, repository, service, ws, id, middleware, model, cmd) |
| Go ファイル数 | 31 |
| Web テストファイル | 11 |
| Web テスト数 | 44 |
| Web ファイル数 | 36 |
| Issue数（closed） | 4 |
| PR数（merged） | 5 |
| v1.0.0 タグ | YES |

## レビュー修正サマリー（PR #17）

3エージェント並列レビュー（Go品質・フロントエンド品質・セキュリティ監査）を実施し、全指摘を1PRで修正。

| 重要度 | 検出数 | 修正数 |
|--------|--------|--------|
| P0/Critical | 6 | 6 |
| P1/Major | 15 | 15 |
| P2/Minor | 14 | 14 |

### 主な修正内容
- **セキュリティ**: CORS middleware、セキュリティヘッダ、MaxBytesReader（1MB制限）、内部エラー非露出
- **WebSocket**: サーバー側ハンドラ実装（gorilla/websocket）、フロントエンド再接続（exponential backoff）
- **並行性安全**: Client.Close() を sync.Once に変更、Hub.Broadcast ロック最小化
- **設計**: id パッケージ新設（panic除去）、ProgressService 層追加（責務分離）、Repository ディープコピー
- **フロントエンド**: fetch エラーハンドリング全ページ追加、setTimeout cleanup、API_BASE集約

## 振り返り

### うまくいったこと
- **モノレポ構成**: web/ + api/ の分離により、CIジョブを並列実行でき、ビルド時間を短縮
- **biome v2.4.10 対応**: スキーマバージョン・ルール名変更への対応をCIで早期検出
- **アクセシビリティ優先**: 48px+ タッチターゲット、ARIA属性、ふりがな、音声読み上げをMVPから組み込み
- **テスト充実**: Go 全レイヤー + Web 全コンポーネントのテストを初回PRから含めた
- **3エージェント並列レビュー**: Go品質・フロント品質・セキュリティの3観点を同時レビューし、全35指摘を1PRで修正完了

### 課題・学び
- **biome設定の互換性**: テンプレートの biome.json がv2.0.0で、CI環境のv2.4.10と非互換。テンプレート側のバージョン固定が必要
- **Go HTTP ルート競合**: `/api/classrooms/join/{code}` と `/api/classrooms/{id}/help-requests` が Go 1.22+ の ServeMux で競合。REST設計時にワイルドカードパスの衝突を事前チェックすべき
- **JSON ファイルストア**: MVP では十分だが、並行書き込みでのデータ整合性に限界。非アトミック書き込み問題はレビューで検出された
- **panic in ID generation**: `crypto/rand.Read` のエラーを panic で処理していた。MVP速度優先でも panic は避けるべき
- **WebSocketサーバー側欠落**: フロントがws://に接続するコードはあったがサーバー側ハンドラが未実装だった。E2E動作確認の重要性

### テンプレート改善提案
1. **biome.json テンプレート**: スキーマバージョンを固定バージョンで管理し、`npx biome migrate` をCI前に実行する手順を追加
2. **Layer-2 web-app テンプレート**: Next.js App Router + vitest + @testing-library のボイラープレートを追加（jsx: automatic 設定含む）
3. **Makefile テンプレート（モノレポ用）**: web/ + api/ の統合ビルド・テスト・harvestターゲットをテンプレート化
4. **ルート設計チェックリスト**: RESTful API設計時のパスパターン競合チェック手順を Layer-1 Go テンプレートに追加
5. **セキュリティミドルウェアテンプレート**: CORS・セキュリティヘッダ・MaxBytesReader を Layer-1 Go に追加（初期スキャフォールドから含める）
6. **id パッケージテンプレート**: panic を使わない安全なID生成を Layer-1 Go に追加
