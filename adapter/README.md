![](adapter.png)

# adapter

- 互換性のないクラス同士を組み合わせる
- 継承や移譲
- 既存のクラスを使用したいが、そのインターフェイスが残りのコードと互換性がない場合これ

### メリット
- 単一責任の原則。インターフェイス、データ変換コードをビジネスロジックから分離できる
- オープンクローズド原則。クライアントインターフェイスを解して、アダプタと連携するので、
  既存のクライアントコードを破壊することがない

### デメリット
- unit test しづらい
- 安易に使い回すとスパゲティになるかも
- 小さいパーツに、シンプルに、疎結合にすること

### 他パターンとの関係性
- Bridgeはアプリケーションの一部を互いに独立して開発できるが、Adapterは互換性のないアプリを連携
- Adapterは既存のオブジェクトのインターフェイス変更、Decoratorはインターフェイスを変更せずオブジェクトを拡張
- Adapterはラップされたオブジェクトに異なるインターフェイスを提供、Proxyは同じインターフェイスを提供、Decoratorは拡張されたインターフェイスを提供
- Facadeは既存のオブジェクトの新しいインターフェイスを定義、Adapterは既存のインターフェイスを使用可能にする

### Go の adapter

Go は interface があり、struct の埋め込みがあるので adapter はそれで事足りる

### 例題
- オブジェクト（Lightningポート）
- 同じ機能を異なるインターフェイス（USBポート）で提供するadaptee（Windowsラップトップ）と呼ばれる別のオブジェクト
- クライアントが期待するのと同じインターフェース（Lightningポート）に準拠
- アダプターはLightningコネクターを受け入れ、その信号をUSB形式に変換して、WindowsラップトップのUSBポートに渡す
