![](factory_method.png)
# FactoryMethod
オブジェクトの生成を分離

### メリット
* 柔軟にオブジェクトを利用できる
* 呼び出すクラスを意識しなくて良い

### デメリット
* 機能を増やすにはfactoryクラスのメンテナンスが必要

### ポイント
* DIContainer = Factoryの集合体
* factoryはconrainerパッケージを用意するなり、配列で管理するなりしたほうが、一箇所に集約されて相互importの心配をしなくてよい

### 例題
* TemplateMethodを応用して、
* 太郎くんはウィングガンダムが作りたいと言った
* 次郎くんはゴッドガンダムが作りたいと言った
* この際、どのガンダム作るかは彼らに任せよう！
