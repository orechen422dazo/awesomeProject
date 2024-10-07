# Goの `for` ループの使い方

Goには4種類の `for` ループがあります。それぞれの使い方について、例と共に解説します。

## 1. 伝統的な `for` ループ（C言語と同じ形式）

最も一般的な `for` ループの形式で、初期化部分、条件部分、および後処理部分から成ります。

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}
```

この形式の `for` ループは、配列やスライスなどの要素を一つずつ処理する際に便利です。

## 2. 条件付き `for` ループ（C言語の `while` 文と同じ形式）

条件が真である間ループを繰り返します。初期化や後処理を省略し、条件のみでループを制御します。

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 3 {
        fmt.Println(i)
        i++
    }
}
```

この形式の `for` ループは、条件のみでループを制御したい場合に使用されます。

## 3. 無限ループ

条件なしの `for` 文で、無限にループします。明示的にループを抜ける操作が必要です（例えば `break` 文）。

```go
package main

import "fmt"

func main() {
    i := 0
    for {
        if i >= 3 {
            break
        }
        fmt.Println(i)
        i++
    }
}
```

無限ループは、特定の条件が満たされるまで繰り返し処理を行いたい場合に使用されます。

## 4. 範囲（range）を使った `for` ループ

スライス、配列、マップ、文字列、チャネルの要素を1つずつ処理するための `for` ループです。

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    for index, value := range s {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

この形式の `for` ループは、コレクションの各要素を処理する際に非常に便利です。

---

これら4種類の `for` ループを理解することで、Goでの繰り返し処理の幅が広がります。特に `range` を使ったループは非常に強力で、スライスやマップなどのデータ構造を簡潔に処理するために頻繁に使用されます。