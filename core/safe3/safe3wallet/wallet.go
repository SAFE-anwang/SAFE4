package safe3wallet

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/binary"
    "errors"
    "github.com/btcsuite/btcutil/base58"
    "github.com/ethereum/go-ethereum/berkeleydb"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/crypto/secp256k1"
    "github.com/status-im/keycard-go/hexutils"
)

type KeyPair map[string]string

type MasterKey struct {
    cryptedKey string
    salt string
    derivationMethod uint32
    deriveIterations uint32
    otherDerivationParameters string
}

func GetKeyFromWallet(walletName string, password string) (KeyPair, error) {
    db, err := berkeleydb.NewDB()
    if err != nil {
        return nil, err
    }

    err = db.Open(walletName, "main", berkeleydb.DbBtree, berkeleydb.DbRdOnly)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    cursor, err := db.Cursor()
    if err != nil {
        return nil, err
    }

    retPair := make(KeyPair)
    keys := make(KeyPair)
    ckeys := make(KeyPair)
    mkeys := make(map[uint32]*MasterKey)
    for true {
        dKey, dValue, err := cursor.Get(berkeleydb.CrsNext)
        if err != nil {
            break
        }
        //fmt.Printf("%s, %x, %d, %x, %d\n", string(dKey), dKey, len(dKey), dValue, len(dValue))
        var strType string
        var buf []byte
        for i := 0; i < int(dKey[0]); i++ {
            buf = append(buf, dKey[i + 1])
        }
        strType = string(buf)
        if strType == "key" {
            pubkey := hexutils.BytesToHex(dKey[5:])
            privKey := hexutils.BytesToHex(dValue[1:215])
            keys[pubkey] = privKey
        } else if strType == "ckey" {
            pubkey := hexutils.BytesToHex(dKey[6:])
            cprivkey := hexutils.BytesToHex(dValue[1:])
            ckeys[pubkey] = cprivkey
        } else if strType == "mkey" {
            id := binary.LittleEndian.Uint32(dKey[5:])
            cryptedKey := hexutils.BytesToHex(dValue[1:49])
            salt := hexutils.BytesToHex(dValue[50:58])
            derivationMethod := binary.LittleEndian.Uint32(dValue[58:62])
            deriveIterations := binary.LittleEndian.Uint32(dValue[62:66])
            otherDerivationParameters := hexutils.BytesToHex(dValue[66:])
            mkeys[id] = &MasterKey{cryptedKey: cryptedKey, salt: salt, derivationMethod: derivationMethod, deriveIterations: deriveIterations, otherDerivationParameters: otherDerivationParameters}
        }
    }

    for pubkey, privkey := range keys {
        //fmt.Printf("key, pubkey: %s, privkey: %s\n", pubkey, privkey)
        key, err := crypto.ToECDSA(secp256k1.LoadKey(hexutils.HexToBytes(privkey)))
        if err != nil {
            return nil, err
        }
        if len(pubkey) != 33 {
            retPair[hexutils.BytesToHex(crypto.CompressPubkey(&key.PublicKey))] = hexutils.BytesToHex(key.D.Bytes())
        } else {
            retPair[pubkey] = hexutils.BytesToHex(key.D.Bytes())
        }
    }

    for _, masterkey := range mkeys {
        //fmt.Printf("mkey, id: %d, masterKey: %s, %s, %d, %d, %s\n", id, masterkey.cryptedKey, masterkey.salt, masterkey.derivationMethod, masterkey.deriveIterations, masterkey.otherDerivationParameters)
        aesKey, aesIV := bytesToKeyAES256CBC([]byte(password), hexutils.HexToBytes(masterkey.salt), masterkey.deriveIterations)
        //fmt.Printf("aes, key: %s, iv: %s\n", hexutils.BytesToHex(aesKey), hexutils.BytesToHex(aesIV))
        plainText, err := decrypt(hexutils.HexToBytes(masterkey.cryptedKey), aesKey, aesIV)
        if err != nil {
            return nil, err
        }
        //fmt.Printf("cipherText: %s, plainText: %s\n", masterkey.cryptedKey, hexutils.BytesToHex(plainText))
        for pubkey, cprivkey := range ckeys {
            //fmt.Printf("ckey, pubkey: %s, cprivkey: %s\n", pubkey, cprivkey)
            iv := sha256.Sum256(hexutils.HexToBytes(pubkey))
            iv = sha256.Sum256(iv[:])
            //fmt.Printf("key: %s, iv: %s\n", cprivkey, hexutils.BytesToHex(iv[:]))
            privkey, err := decrypt(hexutils.HexToBytes(cprivkey), plainText, iv[:16])
            if err != nil {
                return nil, err
            }
            //fmt.Printf("privkey: %s\n", hexutils.BytesToHex(privkey))
            retPair[pubkey] = hexutils.BytesToHex(privkey)
        }
    }
    //for pubkey, privkey := range retPair {
    //    fmt.Printf("key, pubkey: %s, privkey: %s\n", pubkey, privkey)
    //}

    return retPair, nil
}

// BytesToKeyAES256CBC implements the SHA256 version of EVP_BytesToKey using AES CBC
func bytesToKeyAES256CBC(data []byte, salt []byte, round uint32) ([]byte, []byte) {
    var lastHash []byte
    h := sha512.New()
    h.Reset()
    // concatenate lastHash, data and salt and write them to the hash
    h.Write(append(lastHash, append(data, salt...)...))
    // passing nil to Sum() will return the current hash value
    lastHash = h.Sum(nil)
    for i := uint32(1); i < round; i++ {
        h.Reset()
        // concatenate lastHash, data and salt and write them to the hash
        h.Write(lastHash)
        // passing nil to Sum() will return the current hash value
        lastHash = h.Sum(nil)
    }
    return lastHash[:32], lastHash[32:48]
}

func decrypt(cipherText []byte, aesKey []byte, aesIV []byte) ([]byte, error) {
    aesBlock, err := aes.NewCipher(aesKey)
    if err != nil {
        return nil, err
    }
    decrypter := cipher.NewCBCDecrypter(aesBlock, aesIV)
    paddedPlaintext := make([]byte, len(cipherText))
    decrypter.CryptBlocks(paddedPlaintext, cipherText)
    plaintext := pkcs7Unpad(paddedPlaintext)
    if plaintext == nil {
        return nil, errors.New("could not decrypt key with given password")
    }
    return plaintext, err
}

func pkcs7Unpad(in []byte) []byte {
    if len(in) == 0 {
        return nil
    }

    padding := in[len(in)-1]
    if int(padding) > len(in) || padding > aes.BlockSize {
        return nil
    } else if padding == 0 {
        return nil
    }

    for i := len(in) - 1; i > len(in)-int(padding)-1; i-- {
        if in[i] != padding {
            return nil
        }
    }
    return in[:len(in)-int(padding)]
}

func ParseKey(key string) []byte {
    //fmt.Printf("%s\n", hexutils.BytesToHex(base58.Decode(key)))
    //fmt.Printf("%s\n", hexutils.BytesToHex(base58.Decode(key)[1:33]))
    return base58.Decode(key)[1:33]
}