my implement the great awesome repository [danhawkins/go-vite-react-example](https://github.com/danhawkins/go-vite-react-example/tree/main). 

Info: [https://dev.to/danhawkins/embed-vite-react-in-golang-binary-with-live-reload-1k4d](https://dev.to/danhawkins/embed-vite-react-in-golang-binary-with-live-reload-1k4d)

# Note
いまのところコンパイルはビルドしているOSのみに対応（Dockerからwin/macへクロスコンパイルする方法がわからない）

# Diff

- switching to [gin](https://gin-gonic.com/) from `echo`
- use [biomejs](https://biomejs.dev/) instead of `eslint` and `prettier`
- substitute [bun](https://bun.sh/) for `yarn`, bundler, pacakge manager
- sqlite3
- open browser

# usage
`make dev`