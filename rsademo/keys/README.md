```
#!/bin/sh
# 生成PKCS#1的公私钥
openssl genrsa -out pkcs1_private.pem 1024
openssl rsa -in pkcs1_private.pem -RSAPublicKey_out -out pkcs1_public.pem

or 不加RSAPublicKey_out 生成PKCS#8公钥， 如果是下面这样的话，公钥解析就需要用ParsePKIXPublicKey上面带RSAPublicKey_out 用的是
openssl rsa -in pkcs1_private.pem -pubout -out rsa_public_key.pem


# 由PKCS#1的私钥，生成PKCS#8的公私钥
openssl pkcs8 -topk8 -inform PEM -in pkcs1_private.pem -outform PEM -nocrypt -out from_pkcs1_private_to_pkcs8_private.pem
openssl rsa -in pkcs1_private.pem -pubout -out from_pkcs1_private_to_pkcs8_public.pem

# 由PKCS#8的私钥，生成PKCS#1的公私钥
openssl rsa -in from_pkcs1_private_to_pkcs8_private.pem -out from_pkcs8_private_to_pkcs1_private.pem
openssl rsa -in from_pkcs1_private_to_pkcs8_private.pem -RSAPublicKey_out -out from_pkcs8_private_to_pkcs1_public.pem

# 由PKCS1公钥生成PKCS#8公钥:
openssl rsa -RSAPublicKey_in -in pkcs1_public.pem -pubout -out from_pkcs1_public_to_pkcs8_public.pem

# 由PKCS8公钥生成PKCS#1公钥:
openssl rsa -pubin -in from_pkcs1_private_to_pkcs8_public.pem -RSAPublicKey_out -out from_pkcs8_public_to_pkcs1_public.pem
```
x509.ParsePKCS1PublicKey (PKCS#1) and 
x509.ParsePKIXPublicKey (PKCS#8).

可以看到当拥有任意一种私钥时，就可以生成所有需要的东西。当拥有一种公钥时，只能生成另一种公钥。
openssl genrsa -out rsaprivatekey.pem 1024
最大加密明文长度为117位，最大解密密文128位
比如：需要加密String param = "id=1&name=张三",实际上只能加密到：id=1&name. 
使用一对1024位密钥私钥-加密后得到密文：[B@39ba5a14.
解密也只能按照128位：[B@39ba5a14 将1024位公钥解密后 得到id=1&name 

密钥可分为1024，2048，4096等位密钥，位数不同，可加解密明文长度不一。 比如1024位密钥： 可加解密明文长度 len = 1024/8 - 11 = 117字节。
------
Encrypting a File with a Password from the Command Line using OpenSSL
openssl genrsa -des3 -out private_pwd.pem 2048
openssl rsa -in private_pwd.pem -pubout -out rsa_public_key.pem


所以我们说的“密钥”其实是它们两者中的其中一组。但我们说的“密钥长度”一般只是指模值的位长度。目前主流可选值：1024、2048、3072、40961024 位 或 2048 位，位数越长，算法越难被破解
目前主流密钥长度至少都是1024bits以上，低于1024bit的密钥已经不建议使用

openssl rsa -RSAPublicKey_in -in pkcs1_public_loc.pem -pubout -out from_pkcs1_public_to_pkcs8_public_dev.pem

/*
公钥加密、私钥解密、私钥签名、公钥验签。

The encoding/json package marshals maps in sorted key order and structs in the order that the fields are declared.

Although the order used by the encoding/json package is not documented, 
it's safe to assume that maps are marshaled in sorted key order and 
structs are marshaled in field declaration order. 
There are many tests in the standard library and elsewhere that depend on these orders.

*/
