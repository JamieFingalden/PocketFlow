# Web管理界面项目

## 项目结构

```
web-admin/
├── public/              # 静态资源
│   ├── index.html
│   └── favicon.ico
├── src/                 # 源代码
│   ├── components/      # 公共组件
│   ├── views/           # 页面视图
│   ├── api/             # API接口
│   ├── utils/           # 工具函数
│   ├── assets/          # 静态资源
│   ├── router/          # 路由配置
│   ├── store/           # 状态管理
│   ├── styles/          # 样式文件
│   └── App.vue          # 根组件
├── package.json         # 依赖管理
├── tsconfig.json        # TypeScript配置
├── vite.config.ts       # Vite配置
└── README.md            # 项目说明
```

## 技术栈

- Vue3 + TypeScript
- Element Plus UI框架
- Vue Router
- Pinia状态管理
- Axios HTTP客户端
- Vite构建工具

## 功能模块

1. 用户管理 (增删改查)
2. 账单管理 (查看所有用户账单)
3. 数据统计 (全局数据报表)
4. 分类管理 (系统分类配置)
5. 系统配置 (参数设置)
6. 权限管理 (角色和权限分配)
7. 日志管理 (操作日志查看)

## 开发环境

1. 安装Node.js 16+和npm 8+
2. 安装Yarn (可选): `npm install -g yarn`
3. 进入项目目录安装依赖: `npm install` 或 `yarn install`
4. 启动开发服务器: `npm run dev` 或 `yarn dev`

## 第三方库

- Element Plus: 管理界面UI框架
- Axios: HTTP客户端
- Vue Router: 路由管理
- Pinia: 状态管理
- ECharts: 数据可视化
- Day.js: 日期处理
- Lodash: 工具函数库