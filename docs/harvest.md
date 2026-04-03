# Harvest: digi-engawa

## メトリクス
| 項目 | 値 |
|------|-----|
| コミット数 | 10 |
| ADR数 | 1 |
| CLAUDE.md行数 | 35 |
| CI | YES (.github/workflows/ci.yml) |
| lefthook.yml | YES |
| Go パッケージ数 | 8 (handler, repository, service, ws, id, middleware, model, cmd) |
| Go ファイル数 | 32 |
| Web テストファイル | 11 |
| Web テスト数 | 44 |
| Web ファイル数 | 36 |
| Issue数（closed） | 4 |
| PR数（merged） | 7 |
| v1.0.0 タグ | YES |

## レビューサイクル

### Round 1（PR #17）
3エージェント並列レビュー（Go品質・フロントエンド品質・セキュリティ監査）を実施。

| 重要度 | 検出数 | 修正数 |
|--------|--------|--------|
| P0/Critical | 6 | 6 |
| P1/Major | 15 | 15 |
| P2/Minor | 14 | 14 |

### Round 2（PR #18）
修正の検証 + 残存・新規問題の再レビュー。

| 重要度 | 検出数 | 修正数 |
|--------|--------|--------|
| P0/Critical | 4 | 4 |
| P1/Major | 5 | 5 |
| P2/Minor | 3 | 2 |

### 累計修正
- **Round 1**: 35件検出 → 35件修正
- **Round 2**: 12件検出（うち新規7件 + 不完全修正5件）→ 11件修正
- **残存**: 認証なし（設計判断としてMVPでは許容）、NewRouterMuxのドキュメント化

### Round 2 主な修正内容
- **WebSocket CheckOrigin**: CORS_ORIGIN環境変数で正規のオリジン検証（CSWSH防止）
- **WebSocket classroomID検証**: 存在しないclassroomへのJoin拒否
- **goroutine対称cleanup**: sync.Onceで両goroutineから安全にcleanup
- **Hub接続制限**: 100/room, 1000 globalでDoS防止
- **WebSocket ReadLimit**: 4096バイトでメモリ枯渇防止
- **HTTPタイムアウト**: ReadHeader 10s, Read/Write 30s, Idle 120s（Slowloris対策）
- **CORS OPTIONS検証**: 不正Origin からのpreflightに403
- **CSPヘッダ**: `default-src 'none'; frame-ancestors 'none'`
- **アトミックファイル書き込み**: temp+renameパターンで全repository（データ破損防止）
- **ファイル権限**: 0o644 → 0o600（owner-only）
- **エラーメッセージ完全マスク**: 全ハンドラで内部エラー非露出

## 振り返り

### うまくいったこと
- **モノレポ構成**: web/ + api/ の分離により、CIジョブを並列実行でき、ビルド時間を短縮
- **アクセシビリティ優先**: 48px+ タッチターゲット、ARIA属性、ふりがな、音声読み上げをMVPから組み込み
- **テスト充実**: Go 全レイヤー + Web 全コンポーネントのテストを初回PRから含めた
- **レビュー→修正→再レビューサイクル**: 2ラウンドで47件の問題を検出・修正。再レビューで修正の不完全さ（エラーマスク漏れ、WebSocket CheckOrigin）を検出できた

### 課題・学び
- **修正の不完全性**: Round 1でエラーメッセージ非露出を修正したが、一部ハンドラ（classroom, help_request）で漏れていた。「全箇所」の修正は機械的なチェックリストが必要
- **WebSocket セキュリティ**: CheckOriginとCORS middlewareは別の防御層。WebSocketはCORSプリフライトが効かないため、CheckOriginが唯一の防御線
- **goroutine cleanup**: 複数goroutineが共有リソース（conn, client）を持つ場合、sync.Once + 共通cleanup関数パターンが安全
- **接続数制限**: WebSocketサーバーは接続数制限がないとgoroutine枯渇DoSに脆弱。Hub層で制限するのが適切
- **アトミック書き込み**: os.WriteFileは非アトミック。temp+renameパターンは必須知識

### テンプレート改善提案
1. **biome.json テンプレート**: スキーマバージョンを固定バージョンで管理
2. **Layer-2 web-app テンプレート**: Next.js App Router + vitest + @testing-library ボイラープレート
3. **Makefile テンプレート（モノレポ用）**: web/ + api/ の統合ビルド・テスト・harvestターゲット
4. **ルート設計チェックリスト**: RESTful API設計時のパスパターン競合チェック
5. **セキュリティミドルウェアテンプレート**: CORS・セキュリティヘッダ・MaxBytesReader・CSP・タイムアウトを初期スキャフォールドから含める
6. **id パッケージテンプレート**: panic を使わない安全なID生成
7. **WebSocket テンプレート**: CheckOrigin検証・ReadLimit・接続数制限・goroutine対称cleanupを含むハンドラ雛形
8. **Repository テンプレート**: atomicWriteFile + 0o600権限 + ディープコピーを標準化
