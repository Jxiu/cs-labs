# hello_vue3
## 环境
### nodejs
- 版本
  - nodejs 23.2.0
#### npm镜像
nrm工具：nrm是一个npm源管理器，允许你快速地在npm源间切换。
```bash
npm install -g nrm
nrm ls
# 测试速度
nrm test
# 使用淘宝镜像
nrm use taobao
```
##### 项目配置
- .npmrc文件
- 配置淘宝镜像
```txt
registry=https://artifactory.sf-express.com/artifactory/api/npm/npm/
```
#### 多环境管理-nvm
nodejs 版本管理工具
```bash
# nvm 安装
nvm -v
nvm list available
nvm install 23.2.0
# check
nvm list
# 切换版本
nvm use 22.11.0
```
##### 项目配置
- 新建.nvmrc文件
- 配置node版本
```txt
v23.2.0
```
- vscode 插件`vsc-nvm`
## 插件
##### vscode 插件
- Vue - Official
- Ant Design Vue helper
- vsc-nvm (nvm for vscode)