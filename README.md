# ciphers.go


#### 加密算法 (ciphers.Enciphers)
|  算法名称   |  支持参数 | 特点 |
|  ----  | ----  |  ----  |
| RSA  | `RSA` , `RSA 1024` , `RSA 2048` , `RSA 3072` , `RSA 4096` | 广泛使用的非对称加密算法 |
| SM2  | `SM2` , `SM2 C1C3C2` , `SM2 C1C2C3` | 中国自主研发的非对称加密算法 |
| AES  | `AES` , `AES CBC` | 广泛使用的对称加密算法 |

#### 编码方式 (ciphers.Encoders)
|  算法名称   |  支持参数 |
|  ----  | ----  |
| BASE58  | `BASE58` , `BASE58 BTC` , `BASE58 Flickr` |
| BASE64  | `BASE64` |
| HEX  | `HEX` |

#### 散列算法 (ciphers.Hashers)
|  算法名称   |  支持参数 |
|  ----  | ----  |
| MD5  | `MD5` |
| SHA256  | `SHA256` |

#### 签名算法 (ciphers.Signers)
|  算法名称   |  支持参数 |
|  ----  | ----  |
| ECDSA  | `ECDSA` , `ECDSA P224` , `ECDSA P256` , `ECDSA P384` , `ECDSA P521` |

#### 秘钥对生成器 (ciphers.KeypairGenerators)
|  算法名称   |  支持参数 |
|  ----  | ----  |
| ECDSA  | `ECDSA` , `ECDSA P224` , `ECDSA P256` , `ECDSA P384` , `ECDSA P521` |
| RSA  | `RSA` , `RSA 1024` , `RSA 2048` , `RSA 3072` , `RSA 4096` |
| SM2  | `SM2` , `SM2 C1C3C2` , `SM2 C1C2C3` |
