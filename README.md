# go-hello

GoLang の cowsay で "Hello, World!"

## インストール

次のコマンドで
`go-hello`
をインストールできます。

```sh
go install github.com/heiwa4126/go-hello@latest
# または
go install -trimpath -ldflags="-s -w" github.com/heiwa4126/go-hello@latest
```

実行ファイルは Go の`$GOBIN`(未設定の場合は`$GOPATH/bin`)に配置されます。

### 実行例

```console
$ go-hello

Version: v0.0.5
Revision: (unknown)
 _______________
< Hello, World! >
 ---------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## 別の Go プロジェクトから利用

利用するプロジェクトのディレクトリで module を初期化し、`say`パッケージを追加します。

```sh
go mod init example.com/hello-sample
go get github.com/heiwa4126/go-hello/say@latest
```

次の内容を`main.go`を作成して、

```go
package main

import (
    "fmt"

    "github.com/heiwa4126/go-hello/say"
)

func main() {
    fmt.Println(say.Say("Hello from another project!"))
}
```

で、

```sh
go run .
```

## このプロジェクトの開発

```sh
aqua i
task
```

### メモ: Windows の場合

[aqua](https://github.com/aquaproj/aqua)
のインストールは
`winget install aquapro.aqua`
が楽です。

### メモ: govulncheck

govulncheck だけは go のバージョンにうるさいので
go をアップデートしたら再度

```sh
go install golang.org/x/vuln/cmd/govulncheck@latest
```

してください

### その他開発サポート

```sh
task fmt
task check
task build
```

などが便利

## GoReleaser を追加した

```sh
task release
```

で`dist/`以下に Linux 版と Windows 版が生成される。

また、GitHub に対して

```sh
git commit -am 'update something`
git tag v9.9.9
git push --follow-tags
```

で GitHub Releases が生成される。
