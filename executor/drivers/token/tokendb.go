package token

import (
	"code.aliyun.com/chain33/chain33/types"
	dbm "code.aliyun.com/chain33/chain33/common/db"
	"code.aliyun.com/chain33/chain33/account"
	"fmt"
)

type TokenDB struct {
	token types.Token
}

func NewTokenDB(preCreate *types.TokenPreCreate, creater string) *TokenDB {
	t := &TokenDB{}
	t.token.Name = preCreate.GetName()
	t.token.Symbol = preCreate.GetSymbol()
	t.token.Introduction = preCreate.GetIntroduction()
	t.token.Total = preCreate.GetTotal()
	t.token.Price = preCreate.GetPrice()
	//token可以由自己进行创建，也可以通过委托给其他地址进行创建
	t.token.Owner = preCreate.GetOwner()
	t.token.Creater = creater
	t.token.Status = types.TokenStatusPreCreated
	return t
}

func (t *TokenDB) Save(db dbm.KVDB, key []byte) {
	set := t.GetKVSet(key)
	for i := 0; i < len(set); i++ {
		db.Set(set[i].GetKey(), set[i].Value)
	}
}

func (t *TokenDB) getLogs(ty int32, status int32) []*types.ReceiptLog{
	var log []*types.ReceiptLog
	value := types.Encode(&types.ReceiptToken{t.token.Symbol, t.token.Owner, t.token.Status})
	log = append(log, &types.ReceiptLog{ty, value})

	return log
}

//key:mavl-create-token-addr-xxx or mavl-token-xxx <-----> value:token
func (t *TokenDB) GetKVSet(key []byte) (kvset []*types.KeyValue) {
	value := types.Encode(&t.token)
	kvset = append(kvset, &types.KeyValue{key, value})
	return kvset
}

func getTokenFromDB(db dbm.KVDB, symbol string, owner string) (token *types.Token, err error) {
	key := tokenAddrKey(symbol, owner)
	value, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	if err = types.Decode(value, token); err != nil {
		return nil, err
	}
	return token, nil
}

func deleteTokenDB(db dbm.KVDB, symbol string) {
	key := tokenKey(symbol)
	db.Set(key, nil)
}

type TokenAction struct {
	coinsAccount *account.AccountDB
	db        dbm.KVDB
	txhash    []byte
	fromaddr  string
	toaddr    string
	blocktime int64
	height    int64
	execaddr  string
}

func NewTokenAction(t *Token, toaddr string, tx *types.Transaction) *TokenAction {
	hash := tx.Hash()
	fromaddr := account.PubKeyToAddress(tx.GetSignature().GetPubkey()).String()
	return &TokenAction{t.GetCoinsAccount(), t.GetDB(), hash, fromaddr, toaddr,
	t.GetBlockTime(), t.GetHeight(), t.GetAddr()}
}

func (action *TokenAction) preCreate(token *types.TokenPreCreate) (*types.Receipt, error) {
	if len(token.GetName()) > types.TokenNameLenLimit {
		return nil, types.ErrTokenNameLen
	} else if len(token.GetIntroduction()) > types.TokenIntroLenLimit {
		return nil, types.ErrTokenIntroLen
	} else if len(token.GetSymbol()) > types.TokenSymbolLenLimit {
		return nil, types.ErrTokenSymbolLen
	}

	if checkTokenExist(token.GetSymbol(), action.db) {
		return nil, types.ErrTokenExist
	}

	receipt, err := action.coinsAccount.ExecFrozen(action.fromaddr, action.execaddr, token.GetPrice())
	if err != nil {
		tokenlog.Error("token precreate ", "addr", action.fromaddr, "execaddr", action.execaddr, "amount", token.GetTotal())
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	tokendb := NewTokenDB(token, action.fromaddr)
	key := tokenAddrKey(tokendb.token.Symbol, tokendb.token.Owner)
	tokendb.Save(action.db, key)
	logs = append(logs, receipt.Logs...)
	logs = append(logs, tokendb.getLogs(types.TyLogPreCreateToken, types.TokenStatusPreCreated)...)
	kv = append(kv, receipt.KV...)
	kv = append(kv, tokendb.GetKVSet(key)...)

	receipt = &types.Receipt{types.ExecOk, kv, logs}
	return receipt, nil
}

func (action *TokenAction) finishCreate(tokenFinish *types.TokenFinishCreate, approvers []string) (*types.Receipt, error) {
	token, err := getTokenFromDB(action.db, tokenFinish.GetSymbol(), tokenFinish.GetOwner())
	if err != nil  || token.Status != types.TokenStatusPreCreated {
		return nil, types.ErrTokenNotPrecreated
	}

	//确认交易发起者的身份，只有平台审核人才有权限执行该项操作
	approver_valid := false
	for _, approver := range approvers {
		if approver == action.fromaddr {
			approver_valid = true
			break;
		}
	}

	if !approver_valid {
		return nil, types.ErrTokenCreatedApprover
	}

	var receipt	*types.Receipt
    //将之前冻结的资金转账到fund合约账户中
	receiptForCoin, err := action.coinsAccount.ExecTransferFrozen(token.Creater, action.toaddr, action.execaddr, token.Price)
	if err != nil {
		tokenlog.Error("token finishcreate ", "addr", action.fromaddr, "execaddr", action.execaddr, "token", token.Symbol)
		return nil, err
	}

	//创建token类型的账户，同时需要创建的额度存入
	tokenAccount := account.NewTokenAccount(tokenFinish.GetSymbol(), action.db)
	receiptForToken, err := tokenAccount.GenesisInit(token.Owner, token.GetTotal())

	//更新token的状态为已经创建
	token.Status = types.TokenStatusCreated
	tokendb := &TokenDB{*token}
	key := tokenAddrKey(tokendb.token.Symbol, tokendb.token.Owner)
	tokendb.Save(action.db, key)

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	logs = append(logs, receipt.Logs...)
	logs = append(logs, receiptForCoin.Logs...)
	logs = append(logs, receiptForToken.Logs...)
	logs = append(logs, tokendb.getLogs(types.TyLogFinishCreateToken, types.TokenStatusCreated)...)
	kv = append(kv, receiptForCoin.KV...)
	kv = append(kv, receiptForToken.KV...)
	kv = append(kv, tokendb.GetKVSet(key)...)

	key = tokenKey(tokendb.token.Symbol)
	//因为该token已经被创建，需要保存一个全局的token，防止其他用户再次创建
	tokendb.Save(action.db, key)
	kv = append(kv, tokendb.GetKVSet(key)...)

	receipt = &types.Receipt{types.ExecOk, kv, logs}
	return receipt, nil
}

func (action *TokenAction) revokeCreate(tokenRevoke *types.TokenRevokeCreate) (*types.Receipt, error) {
	token, err := getTokenFromDB(action.db, tokenRevoke.GetSymbol(), tokenRevoke.GetOwner())
	if err != nil  || token.Status != types.TokenStatusCreated {
		return nil, types.ErrTokenNotPrecreated
	}

	//确认交易发起者的身份，token的发起人可以撤销该项token的创建
	//token的owner不允许撤销交易
	if action.fromaddr != token.Creater {
		return nil, types.ErrTokenRevokeCreater
	}

	//解锁之前冻结的资金
	receipt, err := action.coinsAccount.ExecActive(token.Owner, action.execaddr, token.Price)
	if err != nil {
		tokenlog.Error("token finishcreate ", "addr", action.fromaddr, "execaddr", action.execaddr, "token", token.Symbol)
		return nil, err
	}

	token.Status = types.TokenStatusCreateRevoked
	tokendb := &TokenDB{*token}
	key := tokenAddrKey(tokendb.token.Symbol, tokendb.token.Owner)
	tokendb.Save(action.db, key)

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	logs = append(logs, receipt.Logs...)
	logs = append(logs, tokendb.getLogs(types.TyLogRevokeCreateToken, types.TokenStatusCreateRevoked)...)
	kv = append(kv, receipt.KV...)
	kv = append(kv, tokendb.GetKVSet(key)...)

	receipt = &types.Receipt{types.ExecOk, kv, logs}
	return receipt, nil
}


func checkTokenExist(token string, db dbm.KVDB) bool {
	_, err := db.Get(tokenKey(token))
	return err == nil
}

func tokenKey(token string) (key []byte) {
	key = append(key, []byte("mavl-token-")...)
	key = append(key, []byte(token)...)
	return key
}

func tokenAddrKey(token string, owner string) (key []byte) {
	key = append(key, []byte("mavl-create-token-")...)
	key = append(key, []byte(owner)...)
	key = append(key, []byte("-")...)
	key = append(key, []byte(token)...)
	return key
}

func tokenStatusKey(token string, owner string, status int32) ([]byte) {
	return []byte(fmt.Sprint("mavl-create-token-%d:%s:%s:%s", status, token, owner))
}

func tokenStatusKeyPrefix(status int32) ([]byte) {
	return []byte(fmt.Sprint("mavl-create-token-%d:", status))
}