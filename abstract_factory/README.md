# AbstractFactory
インスタンス生成をすべて工場に任せる

具体的なクラスを指定せず、関連するオブジェクトのファミリを生成できる

コードが関連製品のさまざまなファミリ（e.g. 椅子、机、ソファ）で機能する必要があるが、

それらの製品の具体的なクラスに依存自宅ない場合はこれ。

事前に不明な場合や、将来の拡張性を考慮したい場合とか。

### メリット
- 整合性が必要なオブジェクトを間違いなく生成できる
- 単一責任の原則。関連し合うオブジェクトの集まりを生成できる
- オープンクローズドの原則。使う側の修正を少なく、インスタンス生成を切り替えたい

### デメリット
- コードが本来より複雑になる可能性
- 関連オブジェクトの種類が増えた場合、修正箇所がfactoryクラスに及んでしまう
    - abstract factory インターフェイスを適切に定義する必要がある

### 他パターンとの関係性
- FactoryMethod > AbstractFactory > Prototype > Builder順で使う
- Bridgeで一部を特定の実装でのみ機能させることができる
- サブシステムオブジェクトの作成方法をクライアントから隠したい時のみ、Facadeの代替になる

### factory vs abstractFactory
factory: 生成しなければならないオブジェクトを事前に知ることができない場合
abstractFactory: オブジェクトを選ぶ場合

### 例題1
- ドアを取り付ける
- 木製のドアと木製ドア取り付け職人
- 鉄製のドアと鉄製ドア取り付け職人


### 例題2
- スポーツキット
- 靴とシャツを購入する必要がある
