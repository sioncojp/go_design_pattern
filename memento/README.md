# Memento
実装の詳細を明らかにすることなく、オブジェクトの以前の状態を保存および復元できる動作設計パターン

オブジェクトの状態のスナップショットを作成して、オブジェクトの以前の状態を復元できるようにする場合

### メリット
- カプセル化に違反することなく、オブジェクトの状態のスナップショットを作成できます。
- caretakerにoriginatorの状態の履歴を維持させることにより、originatorのコードを単純化することができます。

### デメリット
- クライアントがメモリを頻繁に作成する場合、アプリは大量のRAMを消費する可能性があります。
- caretakerは、廃止されたmementoを破壊できる。originatorのライフサイクルを追跡する必要があります。
- PHP、Python、JavaScriptなどのほとんどの動的プログラミング言語は、メモリ内の状態が変更されないことを保証できません。

### 他パターンとの関係性
- 「元に戻す」を実装するときは、CommandとMementoを一緒に使用
- MementoをIteratorと一緒に使用して、現在の反復状態をキャプチャし、必要に応じてロールバックできる
- Prototypeは、Mementoのより簡単な代替手段になる場合がある

### 例題
- オブジェクトの状態のスナップショットを保存
- これらのスナップショットを使用して、オブジェクトを前の状態に戻す