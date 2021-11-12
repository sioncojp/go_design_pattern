# go_design_pattern

## デザインパターンとは？

* オブジェクト指向（OOP）を駆使
* 特定の設計上の問題を解決
* GoF が成功したデザインを紹介

## メリット

* 問題解決の迅速化
* パターンに名前がついてることでコミュニケーションがしやすい

## デメリット

* 全てにおいて最善ではない（取捨選択をきっちりしよう）


## Creational Patterns
- [abstract_factory](./abstract_factory)
- [builder](./builder)
- [factory_method](./factory_method)
- [prototype](./prototype)
- [singleton](./singleton)

## Structual Patterns
- [adapter](./adapter)
- [bridge](./bridge)
- [composite](./composite)
- [decorator](./decorator)
- [facade](./facade)
- [flyweight](./flyweight)
- [proxy](./proxy)

## Behavior Patterns
- [chain_of_repository](./chain_of_repository)
- [iterator](./iterator)
- [Memento](./memento)
- [state](./state)
- [template_method](./template_method)
- [command](./command)
- [mediator](./mediator)
- [observer](./observer)
- [strategy](./strategy)
- [visitor](./visitor)

## 実装例
|pattern|example|
|:---|:---|
|abstract_factory|スポーツキット（靴とシャツを購入する）|
|builder|スポーツカーとファミリーカー。同じ車だけど速度が違う|
|factory_method|ak47, musketなど銃を作る。powerが違う|
|prototype|OSファイルシステム。フォルダとファイル。inode|
|singleton|本の貸し出しリスト|
|adapter|LightningポートをPCのUSBポートで提供する|
|bridge|MacとWindowsの2種類のコンピューターでEpsonとHPのプリンターと互換性を持つ|
|composite|ファイルシステム。ファイルとフォルダがある|
|decorator|ピザのトッピング。それによって値段が変わる|
|facade|クレジットカードを使ってピザを注文。<br> アカウントの確認 -> セキュリティPINを確認 -> クレジット/デビットカードの残高 -> 元帳エントリ作成 -> 通知を送信|
|flyweight|CounterStrikeのゲーム。TとCTでドレスが違い、5vs5 = 10人分用意する|
|proxy|nginx|
|chain_of_repository|病院アプリ。患者が到着 -> 受付 -> 医者 -> 薬室 -> レジに行く。順序がある|
|iterator|本棚|
|memento|オブジェクトの状態のスナップショットを保存。revertできる|
|state|自動販売機。アイテムを選択->アイテムを追加->お金を挿入する->ディスペンスアイテム|
|template_method|ワンタイムパスワード（OTP）機能。ロジックは同じだが、SMS、電子メールで送信する先が違う|
|command|テレビをつけるとき。リモートのONボタン。実際のテレビのオンボタン|
|mediator|駅の交通システム。出発した列車は駅に通知し、キュー内の次の列車が到着させる|
|observer|ECサイト。<br>顧客はある頻度で在庫チェックし、サイトは在庫があるものをレコメンド。顧客はサブスクライブ|
|strategy|インメモリキャッシュ。LRU/FIFO/LFU|
|visitor|三角、四角、円を管理するとき、新しいメソッドを追加するとき|
