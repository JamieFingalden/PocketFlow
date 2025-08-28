# iOS移动端项目

## 项目结构

```
ios/
├── PocketFlow/           # 主应用目录
│   ├── AppDelegate.swift
│   ├── SceneDelegate.swift
│   ├── ViewController/   # 视图控制器
│   ├── View/            # 自定义视图
│   ├── Model/           # 数据模型
│   ├── Service/         # 网络服务
│   ├── Util/            # 工具类
│   └── Resources/       # 资源文件
│       ├── Assets.xcassets
│       ├── Base.lproj
│       └── Info.plist
├── PocketFlow.xcodeproj  # Xcode项目文件
└── Podfile              # CocoaPods依赖管理
```

## 技术栈

- Swift 5.0+
- SwiftUI/UIKit
- iOS 13.0+
- Xcode 14.0+
- CocoaPods/Swift Package Manager

## 功能模块

1. 用户认证 (登录/注册)
2. 账单管理 (增删改查)
3. OCR拍照记账
4. 数据统计图表
5. 预算管理
6. 分类管理
7. 个人设置

## 开发环境

1. 安装Xcode 14.0或更高版本
2. 安装CocoaPods: `sudo gem install cocoapods`
3. 进入项目目录执行: `pod install`
4. 打开`PocketFlow.xcworkspace`文件

## 第三方库

- Alamofire: 网络请求
- Kingfisher: 图片加载
- SwiftyJSON: JSON解析
- Charts: 图表绘制
- SnapKit: 自动布局