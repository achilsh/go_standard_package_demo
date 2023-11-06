module std_pkg_demo

go 1.20

require golang.org/x/text v0.13.0 // indirect

// 声明该模块的某个版本有问题，引用该模块方，如果使用了有问题版本。就会在go list 是提醒有问题，需要手动去更新新版本。
// 这样做的目的是为了：当有些版本出现问题时， 模块维护者可以更加主动的通知模块使用方。
// retract v1.18.7 //其中版本就是 打的tag
