# TYP_ENGINEER(エンジニア用タイピングアプリ)

・タイピング画面↓

<img width="652" alt="スクリーンショット 2023-02-20 15 07 31" src="https://user-images.githubusercontent.com/71370709/220164632-a14bdbd4-c1eb-4650-bb6f-c5e3d7d91281.png">

## 本プロジェクトの発足理由

このアプリ開発者はエンジニアなのにもかかわらず英語が壊滅的なのです。
(5W1Hすら怪しい。。。)
最低限単語は覚えないとと思い、シンプルに勉強した結果
全然覚えられないのですね。

そういえば生まれてから23年、小中高英語から逃げ続けてきた人生です。
そりゃ、普通に勉強してもダメだよなと。。。

そして考えました。
「エンジニアだからアプリでも作って問題解決すればええやん」
なのでエンジニア用英語タイピングアプリを作ってみました。

それが本プロジェクト
**typ-engineerです！！**

## 使用技術,その他

フロントエンドはVue.jsを使用し、バックエンドはGolangでAPIサーバを構築しました。

- フロントエンド : Vue.js
  - Version 2.6.14
- バックエンド : Golang（APIサーバ）
  - Version 1.18
- データベース : MySQL
  - Version 8.0.29
- 仮想環境 : Docker
- [単語の引用元](https://progeigo.org/learning/essential-words-600-plus/)
- SPAで構築

フロントエンドのリポジトリは[こちら](https://github.com/develop-suda/typ-engineer-front)

## 技術選定の経緯

- Golang採用理由
  - [tenntenn(上田拓也)](https://twitter.com/tenntenn)さん講師の[Golang初心者向けハンズオン](https://techplay.jp/event/705479)に参加したのがきっかけ
  - 今までの業務でPHPをメインに触っており、静的型付け言語を使ったことがなかったので勉強したかった
  - 本アプリを開発するにあたりWebで作成することを決め、ここ最近のバックエンドの流行りがGolangなので触りたかった
  - （バックエンド大好き）
  - （裏でこそこそ動いているのがたまらん）

- Vue採用理由
  - フロントはReact,Vueの2つを検討
  - どちらも軽く触ったところで、Vueの方が簡単に理解できたので採用
  - （フロントエンドに対するモチベーションが上がらず上記の理由で採用）
  - （普通だったら世界的にみて採用率が一番高いReactだよなあ）

- MySQL採用理由
  - MySQLとPostgreSQLの2つを検討し、MySQLは業務で触ったことがなかったので採用
  - 世界シェア1位なので
  - [DeveloperSurvey2022](https://survey.stackoverflow.co/2022/#databases)から

- AWS採用理由（デプロイできてないけど）
  - 今までの業務でAWSを触ったことがなかったので勉強したかった
  - Golangを採用しているので、GCPの方が相性良さそうだがAWSが圧倒的なシェアなのでAWSを採用
  - こちらも[DeveloperSurvey2022](https://survey.stackoverflow.co/2022/#cloud-platforms)から

## 改善点

- デプロイされていない
  - これは致命的2023年中にAWSデプロイを目指す
  - [CloudTechで勉強](https://kws-cloud-tech.com/)

- テストコードがない
  - これも由々しき事態

- コードが美しくない
  - なんか冗長なんだよな。。。トランザクション貼るところはボロボロ。。。
  - リファクタリング必須

- CSSが絶望的

## 転職活動中です!!（2023年2月21日）

### 興味があったら是非ご連絡ください

- メールアドレス : developsuda@gmail.com
- [Twitter](https://twitter.com/fumi_elephant)
- [wantedly](https://www.wantedly.com/id/fumi_elephant)
