package main

import (
    "fmt"
    "crypto/rand"
    "encoding/json"
    "encoding/base64"
    "github.com/danielhavir/go-xmss"
)

var params = xmss.SHA2_10_256

func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

func generate() (string, error) {
    jwk := make(map[string]string)
    prv, pub := xmss.GenerateXMSSKeypair(params)
    jwk["kty"] = "OKP"
    jwk["alg"] = "xmss.SHA2_10_256"
    jwk["d"] = string(Base64UrlEncode(*prv))
    jwk["x"] = string(Base64UrlEncode(*pub))
    sjwk, err := json.Marshal(jwk)
    return string(sjwk), err
}

func sign(kp string, msg []byte) (string, error) {
    a := map[string]string{}
    json.Unmarshal([]byte(kp), &a)
    d, err := Base64UrlDecode([]byte(a["d"]))
    if err != nil {
        fmt.Println("failed to convert private key")
    }
    var prv xmss.PrivateXMSS = d
    var sig xmss.SignatureXMSS
    sig = *prv.Sign(params, msg)
    return string(Base64UrlEncode([]byte(sig))), nil
}

func verify( message []byte, signature string, jwk string) (bool, error){
    
    jwkm := map[string]string{}
    json.Unmarshal([]byte(jwk), &jwkm)
    x, err := Base64UrlDecode([]byte(jwkm["x"]))
    if err != nil {
        fmt.Println("failed to decode public key")
    }
    var pub xmss.PublicXMSS = x
    m := make([]byte, params.SignBytes()+len(message))
    sig, err2 := Base64UrlDecode([]byte(signature))
    if err2 != nil {
        fmt.Println("failed to decode signature")
    }
    v := xmss.Verify(params, m, sig, pub)
    return v, nil
}

func main() {

    // generate
    jwk, err1 := generate()
    fmt.Println(jwk)
    if err1 != nil {
        fmt.Println(err1)
    }

    // sign
    msg := make([]byte, 32)
	rand.Read(msg)
    sig, err2 := sign(jwk, msg)
    fmt.Println(sig)
    if err2 != nil {
        fmt.Println(err2)
    }

    // verify
    v, err3 := verify(msg, sig, jwk)
    fmt.Println(v)
    if err3 != nil {
        fmt.Println(err3)
    }
}
