# Digi-Engawa（デジタル縁側）

公民館・地域サロンで開催する「デジタル縁側」教室の運営を支援するプラットフォーム。

## セットアップ

```bash
cp .env.example .env
make install
```

## 開発

```bash
# フロントエンド
cd web && npm run dev

# バックエンド
cd api && go run ./cmd/server
```

## チェック

```bash
make check
```
