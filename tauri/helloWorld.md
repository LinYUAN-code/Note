Get Start Hello World ---2022.9.5


## for Mac
```
xcode-select --install # 会安装command-line 开发的一些基本应用 such: clang gcc git webkit？....
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh # 脚本安装rust
# peek whatever u like
source $HOME/.cargo/env # 设置环境变量

pnpm/npm create tauri-app # 使用tauri脚手架创建一个工程
# peek whatever u like
cd xxx
pnpm/npm i
pnpm/npm tauri dev
Done!

# annoy rust error: proc macro panicked message: The `distDir` configuration is set to `"../dist"` but this path doesn't exist
cd xxx
mkdir dist
# fixed but dont't know why
```
a question: 我们怎么没有安装类似webview的东西就可以跑tauri了

## for Windows



1. OK 环境搭建好了 但是现在我们的窗口点击关闭就直接关闭了，我想要它跑到TaskBar上去，也就是让其成为SystemTray
    - 给我们的SystemTray 加上Hide Quit
    - 给我们的SystemTray 加上About 点击弹出窗口展示作者信息
        - 使用vite-plugin-mpa and tauri multiWindow
        - 两种方式创建窗口
            - 通过tauri.conf.json 静态配置默认创建一个窗口,然后关闭按钮控制为隐藏
            - 在代码中动态创建窗口
                - ques: 这两种方式有什么区别
                    - 静态创建的窗口，从程序一开始就存在会一直占用资源，对于我们这个场景是没必要的，因为用户不会频繁的点击窗口创建
        - 在代码中显现窗口（backend，frontend）

2. 主窗口点击关闭 是变为最小化，只能在taskBar上右键进行关闭
    - 自定义title-bar
        