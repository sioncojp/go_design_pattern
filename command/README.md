# Command
要求を、要求に関するすべての情報を含むスタンドアロンオブジェクトに変換する動作設計パターン

操作でオブジェクトをパラメータ化する場合

### メリット
- 単一責任の原則。操作を呼び出すクラスを、実行するクラスから切り離すことができます。
- オープン/クローズド原則。既存のクライアントコードを壊すことなく、アプリに新しいコマンドを導入できます。
- 元に戻す/やり直しを実装できます。
- 操作の遅延実行を実装できます。
- 単純なコマンドのセットを複雑なコマンドにアセンブルできます。

### デメリット
- 送信者と受信者の間にまったく新しいレイヤーを導入するため、コードがより複雑になる可能性があります。

### 他パターンとの関係性
- これらは似てる
    - ChainOfReponsibility...リクエストを処理するまで、リクエストを順番に渡す
    - Command...送信者と受信者の間に単方向接続を確立
    - Mediator...メディエーターオブジェクトを介して間接的に通信するよう強制
    - Observer...リクエストの受信を動的にサブスクライブおよびサブスクライブ解除できる
- Chain of ResponsibilityのハンドラーとしてCommandが使える
- 「元に戻す」を実装するときは、CommandとMementoを一緒に使用できる
- CommandとStrategyは両方を使用してオブジェクトを何らかのアクションでパラメーター化できるため、似ているように見える場合があるが違う
    - Command...任意の操作をオブジェクトに変換
    - Strategy...同じことを行うさまざまな方法を説明し、単一のコンテキストクラス内でこれらのアルゴリズムを交換
- PrototypeはCommandのコピーを履歴に保存する必要がある場合に使える
- VisitorはCommandの強力なパターンとして使える

### 例題
- テレビをつけるとき
  - リモートのONボタン
  - 実際のテレビのオンボタン
    