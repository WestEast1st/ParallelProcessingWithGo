# Coffman条件
デッドロックのより厳格な定義として、Edgar Coffmanが論文で条件を列挙している。
Coffman条件として知られるこの条件は、デッドロックの検知と予防、訂正を助ける手法の基礎となっている。
## 相互排他
ある並行プロセスがリソースに対して排他的な権利をどの時点においても保持している。
## 条件待ち
ある並行プロセスはリソースの保持と追加のリソースを持ち、同時に行わなければならない。
## 横取り不可
ある並行プロセスによって保持されているリソースは、そのプロセスによって解放される。
## 循環待ち
ある並行プロセスをP1とし、他の連なっている並行プロセスをP2とする。
このP1はP2をまたなければならない、またP2はP1をまたなければならない。

# main.goでの条件の確認
main.goではこの条件が含まれているのかを確認しよう。
1. printSum関数はaとbという変数に排他的アクセス権を必須としているので、この条件を満たしている。("相互排他")
2. printSumはa,bのどちらかを保持しており、もう片方を待っている。("条件待ち")
3. goroutineの横取りを行う方法がない。("横取り不可")
4. printSumの最初の呼び出しで、2番目の呼び出しも待っている。("循環待ち")
