package drivers

import (
	"fmt"
)

const (
	TxIndexFrom    = 1
	TxIndexTo      = 2
	TxIndexPrivacy = 3
)

//用于存储地址相关的hash列表，key=TxAddrHash:addr:height*100000 + index
//地址下面所有的交易
func CalcTxAddrHashKey(addr string, heightindex string) []byte {
	return []byte(fmt.Sprintf("TxAddrHash:%s:%s", addr, heightindex))
}

//用于存储地址相关的hash列表，key=TxAddrHash:addr:flag:height*100000 + index
//地址下面某个分类的交易
func CalcTxAddrDirHashKey(addr string, flag int32, heightindex string) []byte {
	return []byte(fmt.Sprintf("TxAddrDirHash:%s:%d:%s", addr, flag, heightindex))
}

func CalcPrivacyTxHashKey(flag int32, heightindex string) []byte {
	return []byte(fmt.Sprintf("PrivacyTxDirHash:%d:%s", flag, heightindex))
}

func CalcPrivacyTxHashKeyPrefix(flag int32) []byte {
	return []byte(fmt.Sprintf("PrivacyTxDirHash:%d", flag))
}