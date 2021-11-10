# State
内部状態が変化したときにオブジェクトの動作を変更できるようにする動作設計パターン

現在の状態に応じて動作が異なるオブジェクトがあり、状態の数が膨大で、状態固有のコードが頻繁に変更される場合

### メリット
- 単一責任の原則。特定の状態に関連するコードを個別のクラスに編成します。
- オープン/クローズド原則。既存の状態クラスやコンテキストを変更せずに、新しい状態を導入します。
- かさばるステートマシンの条件を排除することにより、コンテキストのコードを簡素化します。

### デメリット
- ステートマシンの状態が少ないか、ほとんど変更されない場合、やりすぎになる可能性

### 他パターンとの関係性
- Bridge、State、Strategy（およびある程度Adapter）の構造は非常に似ているが、他のオブジェクトに作業を委任する構成に基づいてる
- StateはStrategyの延長

### 例題
- 自動販売機
- 4つの状態があるとする 
  - hasItem 
  - noItem
  - itemRequested
  - hasMoney
    
- また4つのアクションがある
  - アイテムを選択
  - アイテムを追加
  - お金を挿入する
  - ディスペンスアイテム
    
- Stateはオブジェクトがさまざまな状態になる可能性があり、に次の状態に切り替わる