# Uniapp移动端项目

## 项目结构

```
uniapp/
├── pages/              # 页面目录
│   ├── index/         # 首页
│   ├── bills/         # 账单管理
│   ├── statistics/    # 数据统计
│   ├── profile/       # 个人中心
│   └── login/         # 登录注册
├── components/         # 自定义组件
├── static/             # 静态资源
├── uni_modules/        # uni_modules组件
├── utils/              # 工具函数
├── api/                # 接口请求
├── store/              # 状态管理
├── manifest.json       # 应用配置
├── pages.json          # 页面配置
├── App.vue             # 应用入口
└── main.js             # 入口文件
```

## 技术栈

- Uniapp跨平台框架
- Vue3 Composition API
- HBuilderX开发工具
- uni_modules组件库
- uView UI组件库
- Pinia状态管理

## 功能模块

1. 用户认证 (登录/注册)
2. 账单管理 (增删改查)
3. OCR拍照记账
4. 数据统计图表
5. 预算管理
6. 分类管理
7. 个人设置

## 开发环境

1. 安装HBuilderX开发工具
2. 安装微信开发者工具(用于微信小程序调试)
3. 安装Android Studio(用于Android调试)
4. 安装Xcode(用于iOS调试)
5. 安装Node.js和npm/yarn
6. 使用Vite开发服务器 (端口: 10086)

## 项目配置

1. 修改`manifest.json`配置应用基本信息
2. 修改`pages.json`配置页面路由
3. 配置不同平台的编译条件
4. Vite开发服务器配置在`vite.config.js`中 (端口: 10086)

## 启动开发服务器

```bash
# 安装依赖
npm install

# 启动Vite开发服务器 (端口: 10086)
npm run dev:h5

# 或者使用yarn
yarn dev:h5
```

## 编译发布

1. 微信小程序: 使用HBuilderX发行到微信小程序
2. Android App: 使用HBuilderX打包为Android应用
3. iOS App: 使用HBuilderX打包为iOS应用
4. H5: 使用HBuilderX发行H5版本