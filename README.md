# go-hello

GoLang の cowsay で "Hello, World!"

## インストール

次のコマンドでインストールできます。

```sh
go install github.com/heiwa4126/go-hello@latest
# または
go install -trimpath -ldflags="-s -w" github.com/heiwa4126/go-hello@latest
```

実行ファイルは Go の`$GOBIN`(未設定の場合は`$GOPATH/bin`)に配置されます。

## 実行

```sh
aqua i
task
```

### 実行例

```console
$ task

task: [run] go run main.go

Version: dev
Revision: unknown
 _______________
< Hello, World! >
 ---------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
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

## 開発中は

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

で Releases が生成される。
