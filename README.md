# hex-gRPC

# ヘキサゴナルアーキテクチャの図
- Domain
- Application
- Framework

- 図に対して右側にDBなどのアプリケーションが触るアダブターが入る(Driven Adapters, Secondary Adapters)
- 図の左側にHTTP, gRPC, CLIなどアプリケーションを触るアダブターが入る(Driving Adapters, Primary Adapters)

- これらのアダブターとアプリケーションロジックの間は**ポート**を介して処理が行われる

## それぞれの層
Domain
- モデル・ビジネスロジック core
Application
- フレームワークからの呼び出しに対応してドメインを呼び出し
Framework
- 各種アダブターが接続するための呼び出し口を提供

- それぞれの外の層は必ず内方向に依存する
  - そのためにDIする必要がある

