```
#!/bin/sh
# 生成PKCS#1的公私钥
openssl genrsa -out pkcs1_private.pem
openssl rsa -in pkcs1_private.pem -RSAPublicKey_out -out pkcs1_public.pem

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

>可以看到当拥有任意一种私钥时，就可以生成所有需要的东西。当拥有一种公钥时，只能生成另一种公钥。





