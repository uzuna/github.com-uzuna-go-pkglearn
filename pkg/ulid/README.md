# ULID: Universally Unique Lexicographically Sortable Identifier

- ソート可能
- 任意のエントロピーソースを使うことのできるID
- URLセーフなcharactorのみ


## 利用ケース

同一の機能を持つアプリケーションが、並列に独自のDBを持ち稼働する環境で、相互にデータを融通可能なIDを生成したい

- 生成頻度は <1000 / day を想定する
- 推測出来ても問題ないが同一インスタンス内、並行稼働するApp間すべてで重複してはならない
    - hostname/ipaddressをもとにする + pidをSeedに使う