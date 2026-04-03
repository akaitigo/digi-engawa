# Harvest: digi-engawa

## メトリクス
| 項目 | 値 |
|------|-----|
| コミット数 | 6 |
| ADR数 | 1 |
| CLAUDE.md行数 | 35 |
| CI | YES (.github/workflows/ci.yml) |
| lefthook.yml | YES |
| Go テストパッケージ | 4 (handler, repository, service, ws) |
| Web テストファイル | 11 |
| Web テスト数 | 44 |
| Issue数（closed） | 4 |
| PR数（merged） | 4 |
| v1.0.0 タグ | YES |

## 振り返り

### うまくいったこと
- **モノレポ構成**: web/ + api/ の分離により、CIジョブを並列実行でき、ビルド時間を短縮
- **biome v2.4.10 対応**: スキーマバージョン・ルール名変更への対応をCIで早期検出
- **アクセシビリティ優先**: 48px+ タッチターゲット、ARIA属性、ふりがな、音声読み上げをMVPから組み込み
- **テスト充実**: Go 全レイヤー + Web 全コンポーネントのテストを初回PRから含めた

### 課題・学び
- **biome設定の互換性**: テンプレートの biome.json がv2.0.0で、CI環境のv2.4.10と非互換。テンプレート側のバージョン固定 or 最新追従が必要
- **Go HTTP ルート競合**: `/api/classrooms/join/{code}` と `/api/classrooms/{id}/help-requests` が Go 1.22+ の ServeMux で競合。REST設計時にワイルドカードパスの衝突を事前チェックすべき
- **JSON ファイルストア**: MVP では十分だが、並行書き込みでのデータ整合性に限界。PostgreSQL 移行を早期に計画すべき

### テンプレート改善提案
1. **biome.json テンプレート**: スキーマバージョンを `latest` ではなく固定バージョンで管理し、`npx biome migrate` をCI前に実行する手順を追加
2. **Layer-2 web-app テンプレート**: Next.js App Router + vitest + @testing-library のボイラープレートを追加（jsx: automatic 設定含む）
3. **Makefile テンプレート（モノレポ用）**: web/ + api/ の統合ビルド・テスト・harvestターゲットをテンプレート化
4. **ルート設計チェックリスト**: RESTful API設計時のパスパターン競合チェック手順を Layer-1 Go テンプレートに追加
