# kifuwarabe-gtp-a1

kifuwarabe-gtp からブランチしたやつ（＾～＾） a1 というのは実験番号で 適当に付けるぜ（＾～＾）

Computer go.  

Base: [https://github.com/bleu48/GoGo](https://github.com/bleu48/GoGo)  

## 考え方

TODO 以下を実装  

1. この作問エンジン（以降 `作問エンジンa1` と呼称）は、自身が考える前に　局面（ポジション）などを `quest-1.toml` ファイルとして `shared/qフォルダー` へ書き出す
   * `quest-1.toml` は、 `quest-2.toml`, `quest-3.toml` のように複数あるものとする

quest-1.toml:  

```toml
# This file is auto generated
[Quest]

BoardSize = 10

Step = 1
Phase = 'x'

BoardData = '''
	+ + + + + + + + + + + +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ + + + + + + + + + + +
'''

[Trick]

Rows = ['A', 'C', 'E', 'G', 'J', 'L', 'N', 'P', 'R', 'T']
Columns = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
```

2. 別の思考エンジン `思考エンジンa2` は `shared/q` フォルダーの中にある `quest-x.toml` を読み取ったあと削除します。  
   そして思考したあと、次の一手を追記した `ans-x.toml` ファイルを `shared/a/1` フォルダーの中に作ります。フォルダーの数字は何手目を打ったかを表します

ans-1.toml:  

```toml
# This file is auto generated
[Quest]

BoardSize = 10

Step = 2
Phase = 'w'

BoardData = '''
	+ + + + + + + + + + + +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . x . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ . . . . . . . . . . +
	+ + + + + + + + + + + +
'''

[Ans]

Bestmove = 'C3'

[Trick]

Rows = ['A', 'C', 'E', 'G', 'J', 'L', 'N', 'P', 'R', 'T']
Columns = [1, 3, 5, 7, 9, 11, 13, 15, 17, 19]
```

3. 作問エンジンa1 は `shared/a/1` フォルダーの中の `ans-*.toml` を全て読取り、その中から１つ選んで次の一手を指す。
   `ans-*.toml` ファイルが存在しなかったなら、投了する。

## Documents

* Set up
  * [on Windows](./doc/set-up-app-on-windows.md)
* Run
  * [on Windows](./doc/run-app-on-windows.md)

## Dependent

* [gtp-engine-to-nngs](https://github.com/muzudho/gtp-engine-to-nngs)

