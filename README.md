# tinyurl
Tiny URL implement, written in Go.

实现时需要思考的一些问题：[zhihu](https://www.zhihu.com/question/20103344)

## 0 Quick Start

短地址生成服务，可作为独立组件使用。示例请参考其配合 [binacs.cn](https://github.com/BinacsLee/server) 的使用。

> 服务试用链接：[binacs.cn/toys/tinyurl](https://binacs.cn/toys/tinyurl)

## 1 Algorithms

总体来说有以下两种方案：

1. 将原始 URL 编码后作为输入直接生成一段短字符串

2. 为原始 URL 分配全局唯一 ID ，随后将 ID 编码生成短字符串

### 1.1 编码原始 URL

1. 将长链接 md5 生成 32 位签名串，分为 4 段，每段 8 个字节;
2. 对这四段循环处理，取 8 个字节，将其转化为 16 进制串与 0x3fffffff(30位1) 与操作;
3. 30 位分成 6 段，每 5 位的数字作为字母表的下标索引取得特定字符，最终获得 6 位字符串;
4. 共获得 4 个 6 位串，取里面的任一个作为长链接 url 的短链接 tiny-url 地址。

本实现直接返回 4 个 6 位串的第一个。

**有碰撞几率；共可以分配 62^6 (56 800 235 584) 个短链接。**

### 1.2 为原始 URL 分配 ID

1. 为长链接分配全局唯一 ID （分布式发号器）;
2. 将 ID 直接转为 62 进制，获取对应短链接。

这样的实现，短链接长度会变化。

## 2 More

其实也有直接生成随机字符串，存储 k-v 对应关系的办法，并不优雅。


