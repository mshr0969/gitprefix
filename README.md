# gitprefix
gitのprefixを表示するCLIツール

## Usage
### 簡単な使い方
```bash
go run main.go
```
以下のように表示される
```bash
Commit Message Prefixes:
feat: featureの略。機能の追加や更新の際に使用する。
fix: bug fixの略。バグの修正の際に使用する。
docs: documentationの略。ドキュメントの変更の際に使用する。
style: コードの意味に影響を与えない変更の際に使用する。(インデント、セミコロンの追加など)
refactor: コードのリファクタリングの際に使用する。
perf: performanceの略。パフォーマンスの向上の際に使用する。
test: テストの追加や修正の際に使用する。
chore: その他、雑務的な変更の際に使用する。
```

### どこからでも使えるようにする
#### バイナリをビルド
```bash
go build -o gitprefix
```

#### バイナリを適切な場所に配置
```bash
sudo mv gitprefix /usr/local/bin/
```

#### 必要があれば権限の付与
```bash
sudo chmod +x /usr/local/bin/gitprefix
```

#### 動かす
```bash
gitprefix
```
