package main

import (
    "fmt"
    "encoding/json"
    "syscall/js"
    // "crypto/rand"
    "encoding/base64"
    // "github.com/danielhavir/go-xmss"
)

func prettyJson(input string) (string, error) {  
    var raw interface{}
    if err := json.Unmarshal([]byte(input), &raw); err != nil {
            return "", err
    }
    pretty, err := json.MarshalIndent(raw, "", "  ")
    if err != nil {
            return "", err
    }
    return string(pretty), nil
}

func jsonWrapper() js.Func {  
    jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
            if len(args) != 1 {
                    return "Invalid no of arguments passed"
            }
            inputJSON := args[0].String()
            fmt.Printf("input %s\n", inputJSON)
            pretty, err := prettyJson(inputJSON)
            if err != nil {
                    fmt.Printf("unable to convert to json %s\n", err)
                    return err.Error()
            }
            return pretty
    })
    return jsonFunc
}

func Base64UrlEncode(src []byte) []byte {
    return []byte(base64.RawURLEncoding.EncodeToString(src))
}

func Base64UrlDecode(src []byte) ([]byte, error) {
    return base64.RawURLEncoding.DecodeString(string(src))
}

func generate() (string, error) {  
 
    a := make(map[string]string)
	
    a["d"] = string(Base64UrlEncode([]byte("priv")))
    a["x"] = string(Base64UrlEncode([]byte("pub")))

    // params := xmss.SHA2_10_256
    // prv, pub := xmss.GenerateXMSSKeypair(params)
    // a["d"] = string(Base64Encode(*prv))
    // a["x"] = string(Base64Encode(*pub))

    b, err := json.Marshal(a)
    return string(b), err
}

func generateJsonWrapper() js.Func {  
    jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        keypair, err := generate()
        if err != nil {
            fmt.Printf("failed to generate keypair %s\n", err)
            return err.Error()
        }
        return keypair
    })
    return jsonFunc
}

func sign(input string) (string, error) {  
    fmt.Printf("signing message: %s\n", input)
    return "signature...", nil
}

func signJsonWrapper() js.Func {  
    jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
                return "Invalid no of arguments passed"
        }
        input := args[0].String()
        signature, err := sign(input)
        if err != nil {
            fmt.Printf("failed to sign %s\n", err)
            return err.Error()
        }
        return signature
    })
    return jsonFunc
}

func verify(input string) (bool, error) {  
    fmt.Printf("verifying signature: %s\n", input)
    return false, nil
}

func verifyJsonWrapper() js.Func {  
    jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
                return "Invalid no of arguments passed"
        }
        input := args[0].String()
        verification, err := verify(input)
        if err != nil {
            fmt.Printf("failed to verify %s\n", err)
            return err.Error()
        }
        return verification
    })
    return jsonFunc
}

func main() {  
    fmt.Println("Go Web Assembly")
    js.Global().Set("generate", generateJsonWrapper())
    js.Global().Set("sign", signJsonWrapper())
    js.Global().Set("verify", verifyJsonWrapper())
    js.Global().Set("formatJSON", jsonWrapper())

    <-make(chan bool)
}