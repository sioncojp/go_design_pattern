# Composit
オブジェクトをツリー構造に構成し操作する

ツリー構造を実装するならこれ

### メリット
- 複雑なツリー構造を便利に操作できる
- オープン/クローズド原則を守り、新しい要素を追加できる

### デメリット
- 機能が、大きく異なるクラスに共通のインターフェースを提供するのは難しい
- 理解が難しくなるかもしれない

### 他パターンとの関係性
- 複合ツリーを作るならbuilderも使える
- Chain of Responsibilityは、Compositeとよく組み合わせる。
- Iteratorで、複合ツリーをトラバース出来る
- Visitorで、ツリー全体を操作できる
- Flyweightsで共有リーフノードを作り、RAMを節約できる
- Decoratorはどちらも再帰的な構造に依存 + さまざまなオブジェクトを編成するための構造体を持ってる
    - Prototypeを使うと、複雑な構造を最初から複製することができる

### 例題
- ファイルシステムにはファイルとフォルダがある
- 検索するとき
    - ファイルの場合、ファイルの名前を探す
    - フォルダの場合、フォルダの全てのファイルを調べる
- 下記ディレクトリ構成となる
```
└── Folder2
    ├── File2
    ├── File3
    └── Folder1
        └── File1
```
