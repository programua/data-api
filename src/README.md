# clean archtectureを基本にしたディレクトリ構成

## controller
いわゆるhandler。リクエストを受けつけ、紐づく処理をコールする。

## usecase
インターフェース。エンティティの操作を行うinteractorを呼び出す。
インプット時、アウトプット時

## interactor
エンティティの作成や操作。usecaseから呼び出される。

## repository
データの永続化を担当する層。DBに対してSave, Update, Deletとかをやる。

## presenter
表示用のデータ加工を担当。例えばJavascriptに返す為にJSON化するとか。

