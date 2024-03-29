# Proxy
別のオブジェクトの代替またはプレースホルダーを提供できる構造設計パターン

1. 遅延初期化（仮想プロキシ）。たまにしか必要ない場合でも、常に稼働しているためにシステムリソースを浪費する重いサービスオブジェクトがある場合

2. アクセス制御（保護プロキシ）。特定のクライアントのみがサービスオブジェクトを使用できるようにする場合

3. リモートサービス（リモートプロキシ）のローカル実行。サービスオブジェクトがリモートサーバーにある場合

4. ロギングリクエスト（ロギングプロキシ）。サービスオブジェクトへのリクエストの履歴を保持する場合

5. リクエスト結果のキャッシュ（プロキシのキャッシュ）。特に結果が非常に大きい場合に、クライアント要求の結果をキャッシュし、このキャッシュのライフサイクルを管理する必要がある場合

6. スマートリファレンス。重いオブジェクトを使用するクライアントがなくなったら、それを却下できるようにする必要がある場合

### メリット
- クライアントが知らないうちにサービスオブジェクトを制御できます。
- クライアントが気にしない場合は、サービスオブジェクトのライフサイクルを管理できます。
- プロキシは、サービスオブジェクトの準備ができていないか、利用できない場合でも機能します。
- オープン/クローズド原則。サービスやクライアントを変更せずに、新しいプロキシを導入できます。

### デメリット
- 多くの新しいクラスを導入する必要があるため、コードはより複雑になる可能性があります。
- サービスからの応答が遅れる場合があります。

### 他パターンとの関係性
- Adapter...ラップされたオブジェクトに異なるインターフェイスを提供
- Proxy...同じインターフェイスを提供
- Decorator...拡張されたインターフェイス
- Facade...複雑なエンティティをバッファリングし、それ自体で初期化するという点で似てる

### 例題
- NginxなどのWebサーバー
- アプリケーションサーバーへの制御されたアクセスを提供
- レート制限
- リクエストのキャッシュ