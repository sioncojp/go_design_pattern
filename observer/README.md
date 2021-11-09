# Observer
監視しているオブジェクトに発生するイベントについて。

複数のオブジェクトに通知するサブスクリプションメカニズムを定義できる動作設計パターン

1つのオブジェクトの状態を変更するために他のオブジェクトを変更する必要があり、

実際のオブジェクトのセットが事前に不明であるか、動的に変更される場合に適用

### メリット
- オープン/クローズド原則。パブリッシャーのコードを変更せずに新しいサブスクライバークラスを導入できる
- 実行時にオブジェクト間の関係を確立

### デメリット
-  サブスクライバーはランダムな順序で通知

### 他パターンとの関係性
- MediatorとObserverの違いそんなにない

### 例題
- ECサイト
- 商品が在庫切れになることがある。
- 在庫切れになった特定の商品に興味を持っているお客様もいるかもしれない。
- 問題が３つある
    - 顧客はある頻度で在庫チェック
    - ECサイトは在庫のある新しいアイテムをレコメンド
    - 顧客は関心のあるアイテムをサブスクライブ。そのアイテムが利用可能かどうかを通知。 複数のお客様が同じ製品を購読することができます

- observerパターンの主なコンポーネントは下記
    - 件名、何かが起こったときにイベントを公開するインスタンス
    - サブジェクトイベントをサブスクライブし、イベントが発生すると通知を受け取るオブザーバー。