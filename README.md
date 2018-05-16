# logging_ecs
ECSで動作するサーバー用のログ出力機能（標準出力がCloudWatchLogへ出力されるように設定されていること）

## 概要
各ログレベル（DEBUG, INFO, WARN, ERROR）ごとに２つの関数を用意
### logger.[LEVEL] (string)
任意のメッセージを出力
- {"level":"DEBUG","message":"test_debug_1","time":"2018-05-07T17:40:17.86595512+09:00"}
### logger.[LEVEL]WithDetail (string, interface{})
任意のメッセージに加えて詳細情報を出力。独自定義したstructを指定することを想定しているが、stringやintなども可
- {"level":"INFO","message":"test_info_5","time":"2018-05-07T17:40:17.87362231+09:00","detail":{"param_int":123,"param_string":"test_param","param_bool":true}}
- {"level":"INFO","message":"test_info_3","time":"2018-05-07T17:40:17.872672441+09:00","detail":"detail_string"}
- {"level":"INFO","message":"test_info_4","time":"2018-05-07T17:40:17.873198194+09:00","detail":"2018-05-07T17:40:17.86595112+09:00"}

### logger.Error (string, error)
### logger.ErrorWithDetail (string, error, interface{})
エラーログはerrorオブジェクトを指定する。（エラーテキストが出力される）
- {"level":"ERROR","message":"test_error_5","time":"2018-05-07T17:40:17.87915789+09:00","error_text":"Test Error","detail":{"param_int":123,"param_string":"test_param","param_bool":true}}

## ライセンス
MIT LICENSE


