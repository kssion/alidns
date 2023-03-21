## 运行条件

- 下载源代码
```sh
git clone https://github.com/kssion/alidns.git
```

- 编译
*Golang 版本要求 1.13 及以上*
```sh
make all
```

- 设置 AccessKey
```sh
./alidns -AK ACCESS_KEY_ID=ACCESS_KEY_SECRET
```

> - 在阿里云帐户中获取您的 [凭证](https://usercenter.console.aliyun.com/#/manage/ak)并通过它替换以上命令中的 ACCESS_KEY_ID 以及 ACCESS_KEY_SECRET;
> - 设置后 AccessKey 将保存在 `$HOME/.alidns/ali.json`

## 在 Certbot 中使用

-  测试运行
```sh
sudo certbot certonly --dry-run --manual --preferred-challenges dns-01 --manual-auth-hook "/usr/bin/alidns" --manual-cleanup-hook "/usr/bin/alidns" -d "example.com,*.example.com"
```