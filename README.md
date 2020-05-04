# ECC PoC

## Create the key pair

OpenSSL contains a large set of pre-defined curves that can be used. The full list of built-in curves can be obtained through the following command:

```bash
$ openssl ecparam -list_curves
```

Keys can be generated from the ecparam command, either through a pre-existing parameters file or directly by selecting the name of the curve. To generate a private/public key pair from the name of a curve:

```bash
$ openssl ecparam -genkey -name secp384r1 -outform der -noout -out secp256k1.der
```

## Read and use the key pair

Run the reference golang software:

```bash
$ go build && ./ecc-test
signature: (0x79a66d07cd87617061f99154eaffae207a3e34d0c7cab914d0b46c7609ee42ea1b2f3ac08e8d0a931de940011975a230, 0xf80898d3206c96540766581371358b1d4d4c6
9a9ca60ece99b28eb1c30c28c3a1f08d78a2184442204d3c6681b5a12d6)
signature verified: true
```