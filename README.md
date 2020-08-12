# muon getting started

muonのGetting StartedっぽいHello, World

[ImVexed/muon: GPU based Electron on a diet](https://github.com/ImVexed/muon)

## 実行方法

binに必要なDLLとかをいれて、`$ PATH=$PATH:./bin go run main.go`

少なくともWindowsではカレントディレクトリかPATHに入っているディレクトリからDLLを探すのでこのテクニックが使えます。
もしBuildするとしてもos.SetEnvで解決できるかも？

