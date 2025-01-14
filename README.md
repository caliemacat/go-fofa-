# go调用fofa爬取使用自己的代理池 
clash 使用的yaml文件 没有完善 后面考虑再补充

使用https://github.com/honmashironeko/ProxyCat 这位师傅所写的代理池软件 并且在这里面 format参数修改成了符合这位师傅软件的格式 直接可使用

-key 传入fofa密钥参数  -s 指定查找的数量 默认5000条

-extra 无F点 使用fofa-view 导出xlsx文件 然后使用这个提取 传入xlsx文件路径

-test 测试端口是否开放 (先测试端口是否开放 再-alive 测试代理是否能使用)

-alive 测试代理是否能使用访问

-format 用于将测试成功的代理 转化为对应的格式 可以使用这位师傅工具 将代理 转化为对应的格式 传入相关文件的地址

 -h 帮助
