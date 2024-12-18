# Port Releaser

使用中のポートを解放するためのシンプルなツールです。

## インストール

```bash
go install github.com/codenamev/port-releaser@latest
```

## 使用方法

```bash
port-releaser -p 3000
```

## パラメータ

- `-p`：解放したいポート番号

## 動作原理

1. `netstat` コマンドを使用して、指定されたポートを使用しているプロセスを特定
2. `taskkill` コマンドを使用してプロセスを終了
3. プロセスが正常に終了できない場合、`-f` パラメータを使用して強制終了

## 注意事項

- 一部のプロセスを終了するには管理者権限が必要
- 重要なシステムプロセスでないことを確認してから使用してください

## ライセンス

本プロジェクトはクリエイティブ・コモンズ 表示-非営利 4.0 国際ライセンス（CC BY-NC 4.0）の下で提供されています。

あなたは以下の条件に従う限り、自由に：
- 共有 — どのようなメディアやフォーマットでも資料を複製、再配布できます
- 翻案 — 資料をリミックスし、改変し、新しい作品を作り出すことができます

以下の条件に従う必要があります：
- 表示 — 適切なクレジットを表示し、ライセンスへのリンクを提供し、変更があったらその旨を示す必要があります
- 非営利 — 著作者の許可なく、資料を営利目的で使用することはできません

詳細：https://creativecommons.org/licenses/by-nc/4.0/deed.ja