# go-hello

GoLang の cowsay で "Hello, World!"

## 実行

```sh
aqua i
task
```

### メモ: Windows の場合

aqua のインストールは`winget install aquapro.aqua`が楽です。

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
