package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"code.aliyun.com/chain33/chain33/account"
	"code.aliyun.com/chain33/chain33/common"
	"code.aliyun.com/chain33/chain33/common/crypto"
	dbm "code.aliyun.com/chain33/chain33/common/db"
	"code.aliyun.com/chain33/chain33/queue"
	"code.aliyun.com/chain33/chain33/types"
	"github.com/golang/protobuf/proto"
	log "github.com/inconshreveable/log15"
)

var (
	minFee            int64 = types.MinFee
	maxTxNumPerBlock  int64 = types.MaxTxsPerBlock
	MaxTxHashsPerTime int64 = 100
	walletlog               = log.New("module", "wallet")
	SignType          int   = 1 //1；secp256k1，2：ed25519，3：sm2
	accountdb               = account.NewCoinsAccount()
)

type Wallet struct {
	qclient        queue.Client
	q              *queue.Queue
	mtx            sync.Mutex
	timeout        *time.Timer
	minertimeout   *time.Timer
	isclosed       int32
	isWalletLocked bool
	isTicketLocked bool

	autoMinerFlag int32
	Password      string
	FeeAmount     int64
	EncryptFlag   int64
	miningTicket  *time.Ticker
	wg            *sync.WaitGroup
	walletStore   *WalletStore
	random        *rand.Rand
	done          chan struct{}
}

func SetLogLevel(level string) {
	common.SetLogLevel(level)
}

func DisableLog() {
	walletlog.SetHandler(log.DiscardHandler())
	storelog.SetHandler(log.DiscardHandler())
}

func New(cfg *types.Wallet) *Wallet {
	//walletStore
	walletStoreDB := dbm.NewDB("wallet", cfg.Driver, cfg.DbPath, 16)
	walletStore := NewWalletStore(walletStoreDB)
	minFee = cfg.MinFee
	if "secp256k1" == cfg.SignType {
		SignType = 1
	} else if "ed25519" == cfg.SignType {
		SignType = 2
	}
	wallet := &Wallet{
		walletStore:    walletStore,
		isWalletLocked: true,
		isTicketLocked: true,
		autoMinerFlag:  0,
		wg:             &sync.WaitGroup{},
		FeeAmount:      walletStore.GetFeeAmount(),
		EncryptFlag:    walletStore.GetEncryptionFlag(),
		miningTicket:   time.NewTicker(2 * time.Minute),
		done:           make(chan struct{}),
	}
	value := walletStore.db.Get([]byte("WalletAutoMiner"))
	if value != nil && string(value) == "1" {
		wallet.autoMinerFlag = 1
	}
	wallet.random = rand.New(rand.NewSource(time.Now().UnixNano()))
	return wallet
}

func (wallet *Wallet) setAutoMining(flag int32) {
	atomic.StoreInt32(&wallet.autoMinerFlag, flag)
}

func (wallet *Wallet) isAutoMining() bool {
	return atomic.LoadInt32(&wallet.autoMinerFlag) == 1
}

func (wallet *Wallet) Close() {
	//等待所有的子线程退出
	//set close flag to isclosed == 1
	atomic.StoreInt32(&wallet.isclosed, 1)
	wallet.miningTicket.Stop()
	close(wallet.done)
	wallet.qclient.Close()
	wallet.wg.Wait()
	//关闭数据库
	wallet.walletStore.db.Close()
	walletlog.Info("wallet module closed")
}

func (wallet *Wallet) IsWalletLocked() bool {
	return wallet.isWalletLocked
}

func (wallet *Wallet) SetQueue(q *queue.Queue) {
	wallet.qclient = q.NewClient()
	wallet.qclient.Sub("wallet")
	wallet.q = q
	wallet.wg.Add(2)
	go wallet.ProcRecvMsg()
	go wallet.autoMining()
}

//检查周期 --> 10分
//开启挖矿：
//1. 自动把成熟的ticket关闭
//2. 查找超过1万余额的账户，自动购买ticket
//3. 查找mineraddress 和他对应的 账户的余额（不在1中），余额超过1万的自动购买ticket 挖矿
//
//停止挖矿：
//1. 自动把成熟的ticket关闭
//2. 查找ticket 可取的余额
//3. 取出ticket 里面的钱
func (wallet *Wallet) autoMining() {
	defer wallet.wg.Done()
	for {
		select {
		case <-wallet.miningTicket.C:
			if !wallet.IsCaughtUp() {
				break
			}
			walletlog.Info("BEG miningTicket")
			if wallet.isAutoMining() {
				n1, err := wallet.closeTicket()
				if err != nil {
					walletlog.Error("closeTicket", "err", err)
				}
				hashes1, n2, err := wallet.buyTicket()
				if err != nil {
					walletlog.Error("buyTicket", "err", err)
				}
				hashes2, n3, err := wallet.buyMinerAddrTicket()
				if err != nil {
					walletlog.Error("buyMinerAddrTicket", "err", err)
				}
				hashes := append(hashes1, hashes2...)
				if len(hashes) > 0 {
					wallet.waitTxs(hashes)
				}
				if n1+n2+n3 > 0 {
					wallet.flushTicket()
				}
			} else {
				n1, err := wallet.closeTicket()
				if err != nil {
					walletlog.Error("closeTicket", "err", err)
				}
				hashes, err := wallet.withdrawFromTicket()
				if err != nil {
					walletlog.Error("withdrawFromTicket", "err", err)
				}
				if len(hashes) > 0 {
					wallet.waitTxs(hashes)
				}
				if n1 > 0 {
					wallet.flushTicket()
				}
			}
			walletlog.Info("END miningTicket")
		case <-wallet.done:
			return
		}
	}
}

func (wallet *Wallet) buyTicket() ([][]byte, int, error) {
	privs, err := wallet.getAllPrivKeys()
	if err != nil {
		walletlog.Error("buyTicket.getAllPrivKeys", "err", err)
		return nil, 0, err
	}
	count := 0
	var hashes [][]byte
	for _, priv := range privs {
		hash, n, err := wallet.buyTicketOne(priv)
		if err != nil {
			walletlog.Error("buyTicketOne", "err", err)
			continue
		}
		count += n
		if hash != nil {
			hashes = append(hashes, hash)
		}
	}
	return hashes, count, nil
}

func (wallet *Wallet) buyMinerAddrTicket() ([][]byte, int, error) {
	privs, err := wallet.getAllPrivKeys()
	if err != nil {
		walletlog.Error("buyMinerAddrTicket.getAllPrivKeys", "err", err)
		return nil, 0, err
	}
	count := 0
	var hashes [][]byte
	for _, priv := range privs {
		hashlist, n, err := wallet.buyMinerAddrTicketOne(priv)
		if err != nil {
			if err != types.ErrNotFound {
				walletlog.Error("buyMinerAddrTicketOne", "err", err)
			}
			continue
		}
		count += n
		if hashlist != nil {
			hashes = append(hashes, hashlist...)
		}
	}
	return hashes, count, nil
}

func (wallet *Wallet) withdrawFromTicket() (hashes [][]byte, err error) {
	privs, err := wallet.getAllPrivKeys()
	if err != nil {
		walletlog.Error("withdrawFromTicket.getAllPrivKeys", "err", err)
		return nil, err
	}
	for _, priv := range privs {
		hash, err := wallet.withdrawFromTicketOne(priv)
		if err != nil {
			walletlog.Error("withdrawFromTicketOne", "err", err)
			continue
		}
		if hash != nil {
			hashes = append(hashes, hash)
		}
	}
	return hashes, nil
}

func (wallet *Wallet) closeTicket() (int, error) {
	return wallet.closeAllTickets()
}

func (wallet *Wallet) forceCloseTicket() ([][]byte, error) {
	return wallet.forceCloseAllTicket()
}

func (wallet *Wallet) flushTicket() {
	walletlog.Info("wallet FLUSH TICKET")
	hashList := wallet.qclient.NewMessage("consensus", types.EventFlushTicket, nil)
	wallet.qclient.Send(hashList, false)
}

func (wallet *Wallet) ProcRecvMsg() {
	defer wallet.wg.Done()
	for msg := range wallet.qclient.Recv() {
		walletlog.Debug("wallet recv", "msg", msg)
		msgtype := msg.Ty
		switch msgtype {

		case types.EventWalletGetAccountList:
			WalletAccounts, err := wallet.ProcGetAccountList()
			if err != nil {
				walletlog.Error("ProcGetAccountList", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccountList, err))
			} else {
				walletlog.Debug("process WalletAccounts OK")
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccountList, WalletAccounts))
			}

		case types.EventWalletAutoMiner:
			flag := msg.GetData().(*types.MinerFlag).Flag
			if flag == 1 {
				wallet.walletStore.db.Set([]byte("WalletAutoMiner"), []byte("1"))
			} else {
				wallet.walletStore.db.Set([]byte("WalletAutoMiner"), []byte("0"))
			}
			wallet.setAutoMining(flag)
			wallet.flushTicket()
			msg.ReplyErr("WalletSetAutoMiner", nil)

		case types.EventWalletGetTickets:
			tickets, privs, err := wallet.GetTickets(1)
			if err != nil {
				walletlog.Error("GetTickets", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("consensus", types.EventWalletTickets, err))
			} else {
				tks := &types.ReplyWalletTickets{tickets, privs}
				walletlog.Debug("process GetTickets OK")
				msg.Reply(wallet.qclient.NewMessage("consensus", types.EventWalletTickets, tks))
			}

		case types.EventNewAccount:
			NewAccount := msg.Data.(*types.ReqNewAccount)
			WalletAccount, err := wallet.ProcCreatNewAccount(NewAccount)
			if err != nil {
				walletlog.Error("ProcCreatNewAccount", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, WalletAccount))
			}

		case types.EventWalletTransactionList:
			WalletTxList := msg.Data.(*types.ReqWalletTransactionList)
			TransactionDetails, err := wallet.ProcWalletTxList(WalletTxList)
			if err != nil {
				walletlog.Error("ProcWalletTxList", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventTransactionDetails, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventTransactionDetails, TransactionDetails))
			}

		case types.EventWalletImportprivkey:
			ImportPrivKey := msg.Data.(*types.ReqWalletImportPrivKey)
			WalletAccount, err := wallet.ProcImportPrivKey(ImportPrivKey)
			if err != nil {
				walletlog.Error("ProcImportPrivKey", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, WalletAccount))
			}
			wallet.flushTicket()

		case types.EventWalletSendToAddress:
			SendToAddress := msg.Data.(*types.ReqWalletSendToAddress)
			ReplyHash, err := wallet.ProcSendToAddress(SendToAddress)
			if err != nil {
				walletlog.Error("ProcSendToAddress", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, ReplyHash))
			}

		case types.EventWalletSetFee:
			WalletSetFee := msg.Data.(*types.ReqWalletSetFee)

			var reply types.Reply
			reply.IsOk = true
			err := wallet.ProcWalletSetFee(WalletSetFee)
			if err != nil {
				walletlog.Error("ProcWalletSetFee", "err", err.Error())
				reply.IsOk = false
				reply.Msg = []byte(err.Error())
			}
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReply, &reply))

		case types.EventWalletSetLabel:
			WalletSetLabel := msg.Data.(*types.ReqWalletSetLabel)
			WalletAccount, err := wallet.ProcWalletSetLabel(WalletSetLabel)

			if err != nil {
				walletlog.Error("ProcWalletSetLabel", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventWalletAccount, WalletAccount))
			}

		case types.EventWalletMergeBalance:
			MergeBalance := msg.Data.(*types.ReqWalletMergeBalance)
			ReplyHashes, err := wallet.ProcMergeBalance(MergeBalance)
			if err != nil {
				walletlog.Error("ProcMergeBalance", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, ReplyHashes))
			}

		case types.EventWalletSetPasswd:
			SetPasswd := msg.Data.(*types.ReqWalletSetPasswd)

			var reply types.Reply
			reply.IsOk = true
			err := wallet.ProcWalletSetPasswd(SetPasswd)
			if err != nil {
				walletlog.Error("ProcWalletSetPasswd", "err", err.Error())
				reply.IsOk = false
				reply.Msg = []byte(err.Error())
			}
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReply, &reply))

		case types.EventWalletLock:
			var reply types.Reply
			reply.IsOk = true
			err := wallet.ProcWalletLock()
			if err != nil {
				walletlog.Error("ProcWalletLock", "err", err.Error())
				reply.IsOk = false
				reply.Msg = []byte(err.Error())
			}
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReply, &reply))

		case types.EventWalletUnLock:
			WalletUnLock := msg.Data.(*types.WalletUnLock)
			var reply types.Reply
			reply.IsOk = true
			err := wallet.ProcWalletUnLock(WalletUnLock)
			if err != nil {
				walletlog.Error("ProcWalletUnLock", "err", err.Error())
				reply.IsOk = false
				reply.Msg = []byte(err.Error())
			}
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReply, &reply))
			wallet.flushTicket()

		case types.EventAddBlock:
			block := msg.Data.(*types.BlockDetail)
			wallet.ProcWalletAddBlock(block)
			walletlog.Debug("wallet add block --->", "height", block.Block.GetHeight())

		case types.EventDelBlock:
			block := msg.Data.(*types.BlockDetail)
			wallet.ProcWalletDelBlock(block)
			walletlog.Debug("wallet del block --->", "height", block.Block.GetHeight())

		//seed
		case types.EventGenSeed:
			genSeedLang := msg.Data.(*types.GenSeedLang)
			replySeed, err := wallet.genSeed(genSeedLang.Lang)
			if err != nil {
				walletlog.Error("genSeed", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyGenSeed, err))
			} else {
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyGenSeed, replySeed))
			}

		case types.EventGetSeed:
			Pw := msg.Data.(*types.GetSeedByPw)
			seed, err := wallet.getSeed(Pw.Passwd)
			if err != nil {
				walletlog.Error("getSeed", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyGetSeed, err))
			} else {
				var replySeed types.ReplySeed
				replySeed.Seed = seed
				walletlog.Error("EventGetSeed", "seed", seed)
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyGetSeed, &replySeed))
			}

		case types.EventSaveSeed:
			saveseed := msg.Data.(*types.SaveSeedByPw)
			var reply types.Reply
			reply.IsOk = true
			ok, err := wallet.saveSeed(saveseed.Passwd, saveseed.Seed)
			if !ok {
				walletlog.Error("saveSeed", "err", err.Error())
				reply.IsOk = false
				reply.Msg = []byte(err.Error())
			}
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReply, &reply))

		case types.EventGetWalletStatus:
			s := wallet.GetWalletStatus()
			msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyWalletStatus, s))

		case types.EventDumpPrivkey:
			addr := msg.Data.(*types.ReqStr)
			privkey, err := wallet.ProcDumpPrivkey(addr.Reqstr)
			if err != nil {
				walletlog.Error("ProcDumpPrivkey", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyPrivkey, err))
			} else {
				var replyStr types.ReplyStr
				replyStr.Replystr = privkey
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyPrivkey, &replyStr))
			}

		case types.EventCloseTickets:
			var reply types.Reply
			reply.IsOk = true
			hashes, err := wallet.forceCloseTicket()
			if err != nil {
				walletlog.Error("closeTicket", "err", err.Error())
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, err))
			} else {
				var replyHashes types.TxHashList
				replyHashes.Hashes = hashes
				msg.Reply(wallet.qclient.NewMessage("rpc", types.EventReplyHashes, &replyHashes))
			}

		default:
			walletlog.Info("ProcRecvMsg unknow msg", "msgtype", msgtype)
		}
		walletlog.Debug("end process")
	}
}

//output:
//type WalletAccounts struct {
//	Wallets []*WalletAccount
//type WalletAccount struct {
//	Acc   *Account
//	Label string
//获取钱包的地址列表
func (wallet *Wallet) ProcGetAccountList() (*types.WalletAccounts, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	//通过Account前缀查找获取钱包中的所有账户信息
	WalletAccStores, err := wallet.walletStore.GetAccountByPrefix("Account")
	if err != nil || len(WalletAccStores) == 0 {
		walletlog.Info("ProcGetAccountList", "GetAccountByPrefix:err", err)
		return nil, err
	}

	addrs := make([]string, len(WalletAccStores))
	for index, AccStore := range WalletAccStores {
		if len(AccStore.Addr) != 0 {
			addrs[index] = AccStore.Addr
		}
		//walletlog.Debug("ProcGetAccountList", "all AccStore", AccStore.String())
	}
	//获取所有地址对应的账户详细信息从account模块
	accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
	if err != nil || len(accounts) == 0 {
		walletlog.Error("ProcGetAccountList", "LoadAccounts:err", err)
		return nil, err
	}

	//异常打印信息
	if len(WalletAccStores) != len(accounts) {
		walletlog.Error("ProcGetAccountList err!", "AccStores)", len(WalletAccStores), "accounts", len(accounts))
	}

	var WalletAccounts types.WalletAccounts
	WalletAccounts.Wallets = make([]*types.WalletAccount, len(WalletAccStores))

	for index, Account := range accounts {
		var WalletAccount types.WalletAccount
		//此账户还没有参与交易所在account模块没有记录
		if len(Account.Addr) == 0 {
			Account.Addr = addrs[index]
		}
		WalletAccount.Acc = Account
		WalletAccount.Label = WalletAccStores[index].GetLabel()
		WalletAccounts.Wallets[index] = &WalletAccount

		//walletlog.Info("ProcGetAccountList", "LoadAccounts:account", Account.String())
	}
	return &WalletAccounts, nil
}

//input:
//type ReqNewAccount struct {
//	Label string
//output:
//type WalletAccount struct {
//	Acc   *Account
//	Label string
//type Account struct {
//	Currency int32
//	Balance  int64
//	Frozen   int64
//	Addr     string
//创建一个新的账户
func (wallet *Wallet) ProcCreatNewAccount(Label *types.ReqNewAccount) (*types.WalletAccount, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	ok, err := wallet.CheckWalletStatus()
	if !ok {
		return nil, err
	}

	if Label == nil || len(Label.GetLabel()) == 0 {
		walletlog.Error("ProcCreatNewAccount Label is nil")
		return nil, types.ErrInputPara
	}

	//首先校验label是否已被使用
	WalletAccStores, err := wallet.walletStore.GetAccountByLabel(Label.GetLabel())
	if WalletAccStores != nil {
		walletlog.Error("ProcCreatNewAccount Label is exist in wallet!")
		return nil, types.ErrLabelHasUsed
	}

	var Account types.Account
	var walletAccount types.WalletAccount
	var WalletAccStore types.WalletAccountStore

	//生成一个pubkey然后换算成对应的addr
	cr, err := crypto.New(types.GetSignatureTypeName(SignType))
	if err != nil {
		walletlog.Error("ProcCreatNewAccount", "err", err)
		return nil, err
	}
	//通过seed获取私钥, 首先通过钱包密码解锁seed然后通过seed生成私钥
	seed, err := wallet.getSeed(wallet.Password)
	if err != nil {
		walletlog.Error("ProcCreatNewAccount", "getSeed err", err)
		return nil, err
	}
	privkeyhex, err := GetPrivkeyBySeed(wallet.walletStore.db, seed)
	if err != nil {
		walletlog.Error("ProcCreatNewAccount", "GetPrivkeyBySeed err", err)
		return nil, err
	}
	privkeybyte, err := common.FromHex(privkeyhex)
	if err != nil || len(privkeybyte) == 0 {
		walletlog.Error("ProcCreatNewAccount", "FromHex err", err)
		return nil, err
	}
	priv, err := cr.PrivKeyFromBytes(privkeybyte)
	if err != nil {
		walletlog.Error("ProcCreatNewAccount", "PrivKeyFromBytes err", err)
		return nil, err
	}
	addr := account.PubKeyToAddress(priv.PubKey().Bytes())
	Account.Addr = addr.String()
	Account.Currency = 0
	Account.Balance = 0
	Account.Frozen = 0

	walletAccount.Acc = &Account
	walletAccount.Label = Label.GetLabel()

	//使用钱包的password对私钥加密 aes cbc
	Encrypted := CBCEncrypterPrivkey([]byte(wallet.Password), priv.Bytes())
	WalletAccStore.Privkey = common.ToHex(Encrypted)
	WalletAccStore.Label = Label.GetLabel()
	WalletAccStore.Addr = addr.String()

	//存储账户信息到wallet数据库中
	err = wallet.walletStore.SetWalletAccount(false, Account.Addr, &WalletAccStore)
	if err != nil {
		return nil, err
	}

	//获取地址对应的账户信息从account模块
	addrs := make([]string, 1)
	addrs[0] = addr.String()
	accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
	if err != nil {
		walletlog.Error("ProcCreatNewAccount", "LoadAccounts err", err)
		return nil, err
	}
	// 本账户是首次创建
	if len(accounts[0].Addr) == 0 {
		accounts[0].Addr = addr.String()
	}
	walletAccount.Acc = accounts[0]

	//从blockchain模块同步Account.Addr对应的所有交易详细信息
	wallet.wg.Add(1)
	go wallet.ReqTxDetailByAddr(addr.String())

	return &walletAccount, nil
}

//input:
//type ReqWalletTransactionList struct {
//	FromTx []byte
//	Count  int32
//output:
//type WalletTxDetails struct {
//	TxDetails []*WalletTxDetail
//type WalletTxDetail struct {
//	Tx      *Transaction
//	Receipt *ReceiptData
//	Height  int64
//	Index   int64
//获取所有钱包的交易记录
func (wallet *Wallet) ProcWalletTxList(TxList *types.ReqWalletTransactionList) (*types.WalletTxDetails, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	if TxList == nil {
		walletlog.Error("ProcWalletTxList TxList is nil!")
		return nil, types.ErrInputPara
	}
	if TxList.GetDirection() != 0 && TxList.GetDirection() != 1 {
		walletlog.Error("ProcWalletTxList Direction err!")
		return nil, types.ErrInputPara
	}
	WalletTxDetails, err := wallet.walletStore.GetTxDetailByIter(TxList)
	if err != nil {
		walletlog.Error("ProcWalletTxList", "GetTxDetailByIter err", err)
		return nil, err
	}
	return WalletTxDetails, nil
}

//input:
//type ReqWalletImportPrivKey struct {
//	Privkey string
//	Label   string
//output:
//type WalletAccount struct {
//	Acc   *Account
//	Label string
//导入私钥，并且同时会导入交易
func (wallet *Wallet) ProcImportPrivKey(PrivKey *types.ReqWalletImportPrivKey) (*types.WalletAccount, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	ok, err := wallet.CheckWalletStatus()
	if !ok {
		return nil, err
	}

	if PrivKey == nil || len(PrivKey.GetLabel()) == 0 || len(PrivKey.GetPrivkey()) == 0 {
		walletlog.Error("ProcImportPrivKey input parameter is nil!")
		return nil, types.ErrInputPara
	}

	//校验label是否已经被使用
	Account, err := wallet.walletStore.GetAccountByLabel(PrivKey.GetLabel())
	if Account != nil {
		walletlog.Error("ProcImportPrivKey Label is exist in wallet!")
		return nil, types.ErrLabelHasUsed
	}

	//通过privkey生成一个pubkey然后换算成对应的addr
	cr, err := crypto.New(types.GetSignatureTypeName(SignType))
	if err != nil {
		walletlog.Error("ProcImportPrivKey", "err", err)
		return nil, err
	}
	privkeybyte, err := common.FromHex(PrivKey.Privkey)
	if err != nil || len(privkeybyte) == 0 {
		walletlog.Error("ProcImportPrivKey", "FromHex err", err)
		return nil, err
	}
	priv, err := cr.PrivKeyFromBytes(privkeybyte)
	if err != nil {
		walletlog.Error("ProcImportPrivKey", "PrivKeyFromBytes err", err)
		return nil, err
	}
	addr := account.PubKeyToAddress(priv.PubKey().Bytes())

	//对私钥加密
	Encryptered := CBCEncrypterPrivkey([]byte(wallet.Password), privkeybyte)
	Encrypteredstr := common.ToHex(Encryptered)
	//校验PrivKey对应的addr是否已经存在钱包中
	Account, err = wallet.walletStore.GetAccountByAddr(addr.String())
	if Account != nil {
		if Account.Privkey == Encrypteredstr {
			walletlog.Error("ProcImportPrivKey Privkey is exist in wallet!")
			return nil, types.ErrPrivkeyExist
		} else {
			walletlog.Error("ProcImportPrivKey!", "Account.Privkey", Account.Privkey, "input Privkey", PrivKey.Privkey)
			return nil, types.ErrPrivkey
		}
	}

	var walletaccount types.WalletAccount
	var WalletAccStore types.WalletAccountStore
	WalletAccStore.Privkey = Encrypteredstr //存储加密后的私钥
	WalletAccStore.Label = PrivKey.GetLabel()
	WalletAccStore.Addr = addr.String()
	//存储Addr:label+privkey+addr到数据库
	err = wallet.walletStore.SetWalletAccount(false, addr.String(), &WalletAccStore)
	if err != nil {
		walletlog.Error("ProcImportPrivKey", "SetWalletAccount err", err)
		return nil, err
	}

	//获取地址对应的账户信息从account模块
	addrs := make([]string, 1)
	addrs[0] = addr.String()
	accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
	if err != nil {
		walletlog.Error("ProcImportPrivKey", "LoadAccounts err", err)
		return nil, err
	}
	// 本账户是首次创建
	if len(accounts[0].Addr) == 0 {
		accounts[0].Addr = addr.String()
	}
	walletaccount.Acc = accounts[0]
	walletaccount.Label = PrivKey.Label

	//从blockchain模块同步Account.Addr对应的所有交易详细信息
	wallet.wg.Add(1)
	go wallet.ReqTxDetailByAddr(addr.String())

	return &walletaccount, nil
}

//input:
//type ReqWalletSendToAddress struct {
//	From   string
//	To     string
//	Amount int64
//	Note   string
//output:
//type ReplyHash struct {
//	Hashe []byte
//发送一笔交易给对方地址，返回交易hash
func (wallet *Wallet) ProcSendToAddress(SendToAddress *types.ReqWalletSendToAddress) (*types.ReplyHash, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	if SendToAddress == nil {
		walletlog.Error("ProcSendToAddress input para is nil")
		return nil, types.ErrInputPara
	}
	if len(SendToAddress.From) == 0 || len(SendToAddress.To) == 0 {
		walletlog.Error("ProcSendToAddress input para From or To is nil!")
		return nil, types.ErrInputPara
	}

	ok, err := wallet.IsTransfer(SendToAddress.GetTo())
	if !ok {
		return nil, err
	}

	//获取from账户的余额从account模块，校验余额是否充足
	addrs := make([]string, 1)
	addrs[0] = SendToAddress.GetFrom()

	accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
	if err != nil || len(accounts) == 0 {
		walletlog.Error("ProcMergeBalance", "LoadAccounts err", err)
		return nil, err
	}
	Balance := accounts[0].Balance
	amount := SendToAddress.GetAmount()

	if Balance < amount+wallet.FeeAmount {
		return nil, types.ErrInsufficientBalance
	}

	addrto := SendToAddress.GetTo()
	note := SendToAddress.GetNote()
	priv, err := wallet.getPrivKeyByAddr(addrs[0])
	if err != nil {
		return nil, err
	}
	return wallet.sendToAddress(priv, addrto, amount, note)
}

func (wallet *Wallet) getPrivKeyByAddr(addr string) (crypto.PrivKey, error) {
	//获取指定地址在钱包里的账户信息
	Accountstor, err := wallet.walletStore.GetAccountByAddr(addr)
	if err != nil {
		walletlog.Error("ProcSendToAddress", "GetAccountByAddr err:", err)
		return nil, err
	}

	//通过password解密存储的私钥
	prikeybyte, err := common.FromHex(Accountstor.GetPrivkey())
	if err != nil || len(prikeybyte) == 0 {
		walletlog.Error("ProcSendToAddress", "FromHex err", err)
		return nil, err
	}

	privkey := CBCDecrypterPrivkey([]byte(wallet.Password), prikeybyte)
	//通过privkey生成一个pubkey然后换算成对应的addr
	cr, err := crypto.New(types.GetSignatureTypeName(SignType))
	if err != nil {
		walletlog.Error("ProcSendToAddress", "err", err)
		return nil, err
	}
	priv, err := cr.PrivKeyFromBytes(privkey)
	if err != nil {
		walletlog.Error("ProcSendToAddress", "PrivKeyFromBytes err", err)
		return nil, err
	}
	return priv, nil
}

//type ReqWalletSetFee struct {
//	Amount int64
//设置钱包默认的手续费
func (wallet *Wallet) ProcWalletSetFee(WalletSetFee *types.ReqWalletSetFee) error {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	if WalletSetFee.Amount < minFee {
		walletlog.Error("ProcWalletSetFee err!", "Amount", WalletSetFee.Amount, "MinFee", minFee)
		return types.ErrInputPara
	}
	err := wallet.walletStore.SetFeeAmount(WalletSetFee.Amount)
	if err == nil {
		walletlog.Debug("ProcWalletSetFee success!")
		wallet.FeeAmount = WalletSetFee.Amount
	}
	return err
}

//input:
//type ReqWalletSetLabel struct {
//	Addr  string
//	Label string
//output:
//type WalletAccount struct {
//	Acc   *Account
//	Label string
//设置某个账户的标签
func (wallet *Wallet) ProcWalletSetLabel(SetLabel *types.ReqWalletSetLabel) (*types.WalletAccount, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	if SetLabel == nil || len(SetLabel.Addr) == 0 || len(SetLabel.Label) == 0 {
		walletlog.Error("ProcWalletSetLabel input parameter is nil!")
		return nil, types.ErrInputPara
	}
	//校验label是否已经被使用
	Account, err := wallet.walletStore.GetAccountByLabel(SetLabel.GetLabel())
	if Account != nil {
		walletlog.Error("ProcWalletSetLabel Label is exist in wallet!")
		return nil, types.ErrLabelHasUsed
	}
	//获取地址对应的账户信息从钱包中,然后修改label
	Account, err = wallet.walletStore.GetAccountByAddr(SetLabel.Addr)
	if err == nil && Account != nil {
		oldLabel := Account.Label
		Account.Label = SetLabel.GetLabel()
		err := wallet.walletStore.SetWalletAccount(true, SetLabel.Addr, Account)
		if err == nil {
			//新的label设置成功之后需要删除旧的label在db的数据
			wallet.walletStore.DelAccountByLabel(oldLabel)

			//获取地址对应的账户详细信息从account模块
			addrs := make([]string, 1)
			addrs[0] = SetLabel.Addr
			accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
			if err != nil || len(accounts) == 0 {
				walletlog.Error("ProcWalletSetLabel", "LoadAccounts err", err)
				return nil, err
			}
			var walletAccount types.WalletAccount
			walletAccount.Acc = accounts[0]
			walletAccount.Label = SetLabel.GetLabel()
			return &walletAccount, err
		}
	}
	return nil, err
}

//input:
//type ReqWalletMergeBalance struct {
//	To string
//output:
//type ReplyHashes struct {
//	Hashes [][]byte
//合并所有的balance 到一个地址
func (wallet *Wallet) ProcMergeBalance(MergeBalance *types.ReqWalletMergeBalance) (*types.ReplyHashes, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	ok, err := wallet.CheckWalletStatus()
	if !ok {
		return nil, err
	}

	if len(MergeBalance.GetTo()) == 0 {
		walletlog.Error("ProcMergeBalance input para is nil!")
		return nil, types.ErrInputPara
	}

	//获取钱包上的所有账户信息
	WalletAccStores, err := wallet.walletStore.GetAccountByPrefix("Account")
	if err != nil || len(WalletAccStores) == 0 {
		walletlog.Error("ProcMergeBalance", "GetAccountByPrefix err", err)
		return nil, err
	}

	addrs := make([]string, len(WalletAccStores))
	for index, AccStore := range WalletAccStores {
		if len(AccStore.Addr) != 0 {
			addrs[index] = AccStore.Addr
		}
	}
	//获取所有地址对应的账户信息从account模块
	accounts, err := accountdb.LoadAccounts(wallet.qclient, addrs)
	if err != nil || len(accounts) == 0 {
		walletlog.Error("ProcMergeBalance", "LoadAccounts err", err)
		return nil, err
	}

	//异常信息记录
	if len(WalletAccStores) != len(accounts) {
		walletlog.Error("ProcMergeBalance", "AccStores", len(WalletAccStores), "accounts", len(accounts))
	}
	//通过privkey生成一个pubkey然后换算成对应的addr
	cr, err := crypto.New(types.GetSignatureTypeName(SignType))
	if err != nil {
		walletlog.Error("ProcMergeBalance", "err", err)
		return nil, err
	}

	addrto := MergeBalance.GetTo()
	note := "MergeBalance"

	var ReplyHashes types.ReplyHashes
	//ReplyHashes.Hashes = make([][]byte, len(accounts))

	for index, Account := range accounts {
		Privkey := WalletAccStores[index].Privkey
		//解密存储的私钥
		prikeybyte, err := common.FromHex(Privkey)
		if err != nil || len(prikeybyte) == 0 {
			walletlog.Error("ProcMergeBalance", "FromHex err", err)
			return nil, err
		}

		privkey := CBCDecrypterPrivkey([]byte(wallet.Password), prikeybyte)
		priv, err := cr.PrivKeyFromBytes(privkey)
		if err != nil {
			walletlog.Error("ProcMergeBalance", "PrivKeyFromBytes err", err, "index", index)
			//ReplyHashes.Hashes[index] = common.Hash{}.Bytes()
			continue
		}
		//过滤掉to地址
		if Account.Addr == addrto {
			continue
		}
		//获取账户的余额，过滤掉余额不足的地址
		amount := Account.GetBalance()
		if amount < wallet.FeeAmount {
			continue
		}
		amount = amount - wallet.FeeAmount
		v := &types.CoinsAction_Transfer{&types.CoinsTransfer{Amount: amount, Note: note}}
		transfer := &types.CoinsAction{Value: v, Ty: types.CoinsActionTransfer}
		//初始化随机数
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		tx := &types.Transaction{Execer: []byte("coins"), Payload: types.Encode(transfer), Fee: wallet.FeeAmount, To: addrto, Nonce: r.Int63()}
		tx.Sign(int32(SignType), priv)

		//发送交易信息给mempool模块
		msg := wallet.qclient.NewMessage("mempool", types.EventTx, tx)
		wallet.qclient.Send(msg, true)
		_, err = wallet.qclient.Wait(msg)
		if err != nil {
			walletlog.Error("ProcMergeBalance", "Send tx err", err, "index", index)
			//ReplyHashes.Hashes[index] = common.Hash{}.Bytes()
			continue
		}

		ReplyHashes.Hashes = append(ReplyHashes.Hashes, tx.Hash())
	}
	return &ReplyHashes, nil
}

//input:
//type ReqWalletSetPasswd struct {
//	Oldpass string
//	Newpass string
//设置或者修改密码
func (wallet *Wallet) ProcWalletSetPasswd(Passwd *types.ReqWalletSetPasswd) error {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	isok, err := wallet.CheckWalletStatus()
	if !isok && err == types.ErrSaveSeedFirst {
		return err
	}
	//保存钱包的锁状态，需要暂时的解锁，函数退出时再恢复回去
	tempislock := wallet.isWalletLocked
	wallet.isWalletLocked = false
	defer func() {
		wallet.isWalletLocked = tempislock
	}()

	// 钱包已经加密需要验证oldpass的正确性
	if len(wallet.Password) == 0 && wallet.EncryptFlag == 1 {
		isok := wallet.walletStore.VerifyPasswordHash(Passwd.Oldpass)
		if !isok {
			walletlog.Error("ProcWalletSetPasswd Verify Oldpasswd fail!")
			return types.ErrVerifyOldpasswdFail
		}
	}

	if len(wallet.Password) != 0 && Passwd.Oldpass != wallet.Password {
		walletlog.Error("ProcWalletSetPasswd Oldpass err!")
		return types.ErrVerifyOldpasswdFail
	}

	//使用新的密码生成passwdhash用于下次密码的验证
	err = wallet.walletStore.SetPasswordHash(Passwd.Newpass)
	if err != nil {
		walletlog.Error("ProcWalletSetPasswd", "SetPasswordHash err", err)
		return err
	}
	//设置钱包加密标志位
	err = wallet.walletStore.SetEncryptionFlag()
	if err != nil {
		walletlog.Error("ProcWalletSetPasswd", "SetEncryptionFlag err", err)
		return err
	}
	//使用old密码解密seed然后用新的钱包密码重新加密seed
	seed, err := wallet.getSeed(Passwd.Oldpass)
	if err != nil {
		walletlog.Error("ProcWalletSetPasswd", "getSeed err", err)
		return err
	}
	ok, err := SaveSeed(wallet.walletStore.db, seed, Passwd.Newpass)
	if !ok {
		walletlog.Error("ProcWalletSetPasswd", "SaveSeed err", err)
		return err
	}

	//对所有存储的私钥重新使用新的密码加密,通过Account前缀查找获取钱包中的所有账户信息
	WalletAccStores, err := wallet.walletStore.GetAccountByPrefix("Account")
	if err != nil || len(WalletAccStores) == 0 {
		walletlog.Error("ProcWalletSetPasswd", "GetAccountByPrefix:err", err)
	}
	if WalletAccStores != nil {
		for _, AccStore := range WalletAccStores {
			//使用old Password解密存储的私钥
			storekey, err := common.FromHex(AccStore.GetPrivkey())
			if err != nil || len(storekey) == 0 {
				walletlog.Info("ProcWalletSetPasswd", "addr", AccStore.Addr, "FromHex err", err)
				continue
			}
			Decrypter := CBCDecrypterPrivkey([]byte(Passwd.Oldpass), storekey)

			//使用新的密码重新加密私钥
			Encrypter := CBCEncrypterPrivkey([]byte(Passwd.Newpass), Decrypter)
			AccStore.Privkey = common.ToHex(Encrypter)
			err = wallet.walletStore.SetWalletAccount(true, AccStore.Addr, AccStore)
			if err != nil {
				walletlog.Info("ProcWalletSetPasswd", "addr", AccStore.Addr, "SetWalletAccount err", err)
			}
		}
	}
	wallet.Password = Passwd.Newpass
	return nil
}

//锁定钱包
func (wallet *Wallet) ProcWalletLock() error {
	//判断钱包是否已保存seed
	has, _ := HasSeed(wallet.walletStore.db)
	if !has {
		return types.ErrSaveSeedFirst
	}

	wallet.isWalletLocked = true
	wallet.isTicketLocked = true
	return nil
}

//input:
//type WalletUnLock struct {
//	Passwd  string
//	Timeout int64
//解锁钱包Timeout时间，超时后继续锁住
func (wallet *Wallet) ProcWalletUnLock(WalletUnLock *types.WalletUnLock) error {
	//判断钱包是否已保存seed
	has, _ := HasSeed(wallet.walletStore.db)
	if !has {
		return types.ErrSaveSeedFirst
	}
	// 钱包已经加密需要验证passwd的正确性
	if len(wallet.Password) == 0 && wallet.EncryptFlag == 1 {
		isok := wallet.walletStore.VerifyPasswordHash(WalletUnLock.Passwd)
		if !isok {
			walletlog.Error("ProcWalletUnLock Verify Oldpasswd fail!")
			return types.ErrVerifyOldpasswdFail
		}
	}

	//内存中已经记录password时的校验
	if len(wallet.Password) != 0 && WalletUnLock.Passwd != wallet.Password {
		return types.ErrInputPassword
	}
	//本钱包没有设置密码加密过,只需要解锁不需要记录解锁密码
	wallet.Password = WalletUnLock.Passwd

	walletlog.Error("ProcWalletUnLock !", "WalletOrTicket", WalletUnLock.WalletOrTicket)

	//只解锁挖矿的转账
	if WalletUnLock.WalletOrTicket {
		wallet.isTicketLocked = false
	} else {
		wallet.isWalletLocked = false
	}
	if WalletUnLock.Timeout != 0 {
		wallet.resetTimeout(WalletUnLock.WalletOrTicket, WalletUnLock.Timeout)
	}
	return nil

}

//解锁超时处理，需要区分整个钱包的解锁或者只挖矿的解锁
func (wallet *Wallet) resetTimeout(IsTicket bool, Timeout int64) {
	//只挖矿的解锁超时
	if IsTicket {
		if wallet.minertimeout == nil {
			wallet.minertimeout = time.AfterFunc(time.Second*time.Duration(Timeout), func() {
				wallet.isTicketLocked = true
			})
		} else {
			wallet.minertimeout.Reset(time.Second * time.Duration(Timeout))
		}
	} else { //整个钱包的解锁超时
		if wallet.timeout == nil {
			wallet.timeout = time.AfterFunc(time.Second*time.Duration(Timeout), func() {
				wallet.isWalletLocked = true
			})
		} else {
			wallet.timeout.Reset(time.Second * time.Duration(Timeout))
		}
	}
}

//wallet模块收到blockchain广播的addblock消息，需要解析钱包相关的tx并存储到db中
func (wallet *Wallet) ProcWalletAddBlock(block *types.BlockDetail) {
	if block == nil {
		walletlog.Error("ProcWalletAddBlock input para is nil!")
		return
	}
	//walletlog.Error("ProcWalletAddBlock", "height", block.GetBlock().GetHeight())
	txlen := len(block.Block.GetTxs())
	newbatch := wallet.walletStore.NewBatch(true)
	needflush := false
	for index := 0; index < txlen; index++ {
		blockheight := block.Block.Height*maxTxNumPerBlock + int64(index)
		heightstr := fmt.Sprintf("%018d", blockheight)

		var txdetail types.WalletTxDetail
		txdetail.Tx = block.Block.Txs[index]
		txdetail.Height = block.Block.Height
		txdetail.Index = int64(index)
		txdetail.Receipt = block.Receipts[index]
		txdetail.Blocktime = block.Block.BlockTime

		//获取Amount
		amount, err := txdetail.Tx.Amount()
		if err != nil {
			continue
		}
		txdetail.ActionName = txdetail.Tx.ActionName()
		txdetail.Amount = amount

		//获取from地址
		pubkey := block.Block.Txs[index].Signature.GetPubkey()
		addr := account.PubKeyToAddress(pubkey)
		txdetail.Fromaddr = addr.String()

		txdetailbyte, err := proto.Marshal(&txdetail)
		if err != nil {
			storelog.Error("ProcWalletAddBlock Marshal txdetail err", "Height", block.Block.Height, "index", index)
			continue
		}

		//from addr
		fromaddress := addr.String()
		if len(fromaddress) != 0 && wallet.AddrInWallet(fromaddress) {
			newbatch.Set([]byte(calcTxKey(heightstr)), txdetailbyte)
			walletlog.Debug("ProcWalletAddBlock", "fromaddress", fromaddress, "heightstr", heightstr)
			continue
		}

		//toaddr
		toaddr := block.Block.Txs[index].GetTo()
		if len(toaddr) != 0 && wallet.AddrInWallet(toaddr) {
			newbatch.Set([]byte(calcTxKey(heightstr)), txdetailbyte)
			walletlog.Debug("ProcWalletAddBlock", "toaddr", toaddr, "heightstr", heightstr)
		}

		if "ticket" == string(block.Block.Txs[index].Execer) {
			tx := block.Block.Txs[index]
			receipt := block.Receipts[index]
			if wallet.needFlushTicket(tx, receipt) {
				needflush = true
			}
		}
	}
	newbatch.Write()
	if needflush {
		//wallet.flushTicket()
	}
}

func (wallet *Wallet) needFlushTicket(tx *types.Transaction, receipt *types.ReceiptData) bool {
	if receipt.Ty != types.ExecOk || string(tx.Execer) != "ticket" {
		return false
	}
	pubkey := tx.Signature.GetPubkey()
	addr := account.PubKeyToAddress(pubkey)
	if wallet.AddrInWallet(addr.String()) {
		return true
	}
	return false
}

//wallet模块收到blockchain广播的delblock消息，需要解析钱包相关的tx并存db中删除
func (wallet *Wallet) ProcWalletDelBlock(block *types.BlockDetail) {
	if block == nil {
		walletlog.Error("ProcWalletDelBlock input para is nil!")
		return
	}
	//walletlog.Error("ProcWalletDelBlock", "height", block.GetBlock().GetHeight())

	txlen := len(block.Block.GetTxs())
	newbatch := wallet.walletStore.NewBatch(true)
	needflush := false
	for index := 0; index < txlen; index++ {
		blockheight := block.Block.Height*maxTxNumPerBlock + int64(index)
		heightstr := fmt.Sprintf("%018d", blockheight)
		if "ticket" == string(block.Block.Txs[index].Execer) {
			tx := block.Block.Txs[index]
			receipt := block.Receipts[index]
			if wallet.needFlushTicket(tx, receipt) {
				needflush = true
			}
		}
		//获取from地址
		pubkey := block.Block.Txs[index].Signature.GetPubkey()
		addr := account.PubKeyToAddress(pubkey)
		fromaddress := addr.String()
		if len(fromaddress) != 0 && wallet.AddrInWallet(fromaddress) {
			newbatch.Delete([]byte(calcTxKey(heightstr)))
			//walletlog.Error("ProcWalletAddBlock", "fromaddress", fromaddress, "heightstr", heightstr)
			continue
		}
		//toaddr
		toaddr := block.Block.Txs[index].GetTo()
		if len(toaddr) != 0 && wallet.AddrInWallet(toaddr) {
			newbatch.Delete([]byte(calcTxKey(heightstr)))
			//walletlog.Error("ProcWalletAddBlock", "toaddr", toaddr, "heightstr", heightstr)
		}
	}
	newbatch.Write()
	if needflush {
		wallet.flushTicket()
	}
}

//地址对应的账户是否属于本钱包
func (wallet *Wallet) AddrInWallet(addr string) bool {
	if len(addr) == 0 {
		return false
	}
	account, err := wallet.walletStore.GetAccountByAddr(addr)
	if err == nil && account != nil {
		return true
	}
	return false
}
func (wallet *Wallet) GetTxDetailByHashs(ReqHashes *types.ReqHashes) {
	//通过txhashs获取对应的txdetail
	msg := wallet.qclient.NewMessage("blockchain", types.EventGetTransactionByHash, ReqHashes)
	wallet.qclient.Send(msg, true)
	resp, err := wallet.qclient.Wait(msg)
	if err != nil {
		walletlog.Error("GetTxDetailByHashs EventGetTransactionByHash", "err", err)
		return
	}
	TxDetails := resp.GetData().(*types.TransactionDetails)
	if TxDetails == nil {
		walletlog.Info("GetTxDetailByHashs TransactionDetails is nil")
		return
	}

	//批量存储地址对应的所有交易的详细信息到wallet db中
	newbatch := wallet.walletStore.NewBatch(true)
	for _, txdetal := range TxDetails.Txs {
		height := txdetal.GetHeight()
		txindex := txdetal.GetIndex()

		blockheight := height*maxTxNumPerBlock + int64(txindex)
		heightstr := fmt.Sprintf("%018d", blockheight)
		var txdetail types.WalletTxDetail
		txdetail.Tx = txdetal.GetTx()
		txdetail.Height = txdetal.GetHeight()
		txdetail.Index = txdetal.GetIndex()
		txdetail.Receipt = txdetal.GetReceipt()
		txdetail.Blocktime = txdetal.GetBlocktime()
		txdetail.Amount = txdetal.GetAmount()
		txdetail.Fromaddr = txdetal.GetFromaddr()
		txdetail.ActionName = txdetal.GetTx().ActionName()

		txdetailbyte, err := proto.Marshal(&txdetail)
		if err != nil {
			storelog.Error("GetTxDetailByHashs Marshal txdetail err", "Height", height, "index", txindex)
			return
		}
		newbatch.Set([]byte(calcTxKey(heightstr)), txdetailbyte)
		walletlog.Debug("GetTxDetailByHashs", "heightstr", heightstr, "txdetail", txdetail.String())
	}
	newbatch.Write()
}

//从blockchain模块同步addr参与的所有交易详细信息
func (wallet *Wallet) ReqTxDetailByAddr(addr string) {
	defer wallet.wg.Done()
	if len(addr) == 0 {
		walletlog.Error("ReqTxInfosByAddr input addr is nil!")
		return
	}
	var txInfo types.ReplyTxInfo

	i := 0
	for {
		//首先从blockchain模块获取地址对应的所有交易hashs列表,从最新的交易开始获取
		var ReqAddr types.ReqAddr
		ReqAddr.Addr = addr
		ReqAddr.Flag = 0
		ReqAddr.Direction = 0
		ReqAddr.Count = int32(MaxTxHashsPerTime)
		if i == 0 {
			ReqAddr.Height = -1
			ReqAddr.Index = 0
		} else {
			ReqAddr.Height = txInfo.GetHeight()
			ReqAddr.Index = txInfo.GetIndex()
		}
		i++
		msg := wallet.qclient.NewMessage("blockchain", types.EventGetTransactionByAddr, &ReqAddr)
		wallet.qclient.Send(msg, true)
		resp, err := wallet.qclient.Wait(msg)
		if err != nil {
			walletlog.Error("ReqTxInfosByAddr EventGetTransactionByAddr", "err", err, "addr", addr)
			return
		}

		ReplyTxInfos := resp.GetData().(*types.ReplyTxInfos)
		if ReplyTxInfos == nil {
			walletlog.Info("ReqTxInfosByAddr ReplyTxInfos is nil")
			return
		}
		txcount := len(ReplyTxInfos.TxInfos)

		var ReqHashes types.ReqHashes
		ReqHashes.Hashes = make([][]byte, len(ReplyTxInfos.TxInfos))
		for index, ReplyTxInfo := range ReplyTxInfos.TxInfos {
			ReqHashes.Hashes[index] = ReplyTxInfo.GetHash()
			txInfo.Hash = ReplyTxInfo.GetHash()
			txInfo.Height = ReplyTxInfo.GetHeight()
			txInfo.Index = ReplyTxInfo.GetIndex()
		}
		wallet.GetTxDetailByHashs(&ReqHashes)
		if txcount < int(MaxTxHashsPerTime) {
			return
		}
	}
}

//使用钱包的password对私钥进行aes cbc加密,返回加密后的privkey
func CBCEncrypterPrivkey(password []byte, privkey []byte) []byte {
	key := make([]byte, 32)
	Encrypted := make([]byte, len(privkey))
	if len(password) > 32 {
		key = password[0:32]
	} else {
		copy(key, password)
	}

	block, _ := aes.NewCipher(key)
	iv := key[:block.BlockSize()]
	//walletlog.Info("CBCEncrypterPrivkey", "password", string(key), "Privkey", common.ToHex(privkey))

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(Encrypted, privkey)

	//walletlog.Info("CBCEncrypterPrivkey", "Encrypted", common.ToHex(Encrypted))
	return Encrypted
}

//使用钱包的password对私钥进行aes cbc解密,返回解密后的privkey
func CBCDecrypterPrivkey(password []byte, privkey []byte) []byte {
	key := make([]byte, 32)
	if len(password) > 32 {
		key = password[0:32]
	} else {
		copy(key, password)
	}

	block, _ := aes.NewCipher(key)
	iv := key[:block.BlockSize()]
	decryptered := make([]byte, len(privkey))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(decryptered, privkey)
	//walletlog.Info("CBCDecrypterPrivkey", "password", string(key), "Encrypted", common.ToHex(privkey), "decryptered", common.ToHex(decryptered))
	return decryptered
}

//生成一个随机的seed种子, 目前支持英文单词和简体中文
func (wallet *Wallet) genSeed(lang int32) (*types.ReplySeed, error) {
	seed, err := CreateSeed("", lang)
	if err != nil {
		walletlog.Error("genSeed", "CreateSeed err", err)
		return nil, err
	}
	var ReplySeed types.ReplySeed
	ReplySeed.Seed = seed
	return &ReplySeed, nil
}

//获取seed种子, 通过钱包密码
func (wallet *Wallet) getSeed(password string) (string, error) {
	ok, err := wallet.CheckWalletStatus()
	if !ok {
		return "", err
	}

	seed, err := GetSeed(wallet.walletStore.db, password)
	if err != nil {
		walletlog.Error("getSeed", "GetSeed err", err)
		return "", err
	}
	return seed, nil
}

//保存seed种子到数据库中, 并通过钱包密码加密, 钱包起来首先要设置seed
func (wallet *Wallet) saveSeed(password string, seed string) (bool, error) {

	//首先需要判断钱包是否已经设置seed，如果已经设置提示不需要再设置，一个钱包只能保存一个seed
	exit, err := HasSeed(wallet.walletStore.db)
	if exit {
		return false, types.ErrSeedExist
	}
	//入参数校验，seed必须是15个单词或者汉字
	if len(password) == 0 || len(seed) == 0 {
		return false, types.ErrInputPara
	}

	seedarry := strings.Fields(seed)
	if len(seedarry) != SeedLong {
		return false, types.ErrSeedWordNum
	}
	var newseed string
	for index, seedstr := range seedarry {
		//walletlog.Error("saveSeed", "seedstr", seedstr)
		if index != SeedLong-1 {
			newseed += seedstr + " "
		} else {
			newseed += seedstr
		}
	}

	ok, err := SaveSeed(wallet.walletStore.db, newseed, password)
	//seed保存成功需要更新钱包密码
	if ok {
		var ReqWalletSetPasswd types.ReqWalletSetPasswd
		ReqWalletSetPasswd.Oldpass = password
		ReqWalletSetPasswd.Newpass = password
		Err := wallet.ProcWalletSetPasswd(&ReqWalletSetPasswd)
		if Err != nil {
			walletlog.Error("saveSeed", "ProcWalletSetPasswd err", err)
		}
	}
	return ok, err
}

//钱包状态检测函数,解锁状态，seed是否已保存
func (wallet *Wallet) CheckWalletStatus() (bool, error) {
	// 钱包锁定，ticket已经解锁，返回只解锁了ticket的错误
	if wallet.IsWalletLocked() && !wallet.isTicketLocked {
		return false, types.ErrOnlyTicketUnLocked
	} else if wallet.IsWalletLocked() {
		return false, types.ErrWalletIsLocked
	}

	//判断钱包是否已保存seed
	has, _ := HasSeed(wallet.walletStore.db)
	if !has {
		return false, types.ErrSaveSeedFirst
	}
	return true, nil
}

func (wallet *Wallet) GetWalletStatus() *types.WalletStatus {
	s := &types.WalletStatus{}
	s.IsWalletLock = wallet.IsWalletLocked()
	s.IsHasSeed, _ = HasSeed(wallet.walletStore.db)
	s.IsAutoMining = wallet.isAutoMining()
	s.IsTicketLock = wallet.isTicketLocked
	return s
}

//获取地址对应的私钥
func (wallet *Wallet) ProcDumpPrivkey(addr string) (string, error) {
	wallet.mtx.Lock()
	defer wallet.mtx.Unlock()

	ok, err := wallet.CheckWalletStatus()
	if !ok {
		return "", err
	}
	if len(addr) == 0 {
		walletlog.Error("ProcDumpPrivkey input para is nil!")
		return "", types.ErrInputPara
	}

	priv, err := wallet.getPrivKeyByAddr(addr)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(common.ToHex(priv.Bytes())), nil
}

//检测钱包是否允许转账到指定地址，判断钱包锁和是否有seed以及挖矿锁
func (wallet *Wallet) IsTransfer(addr string) (bool, error) {

	ok, err := wallet.CheckWalletStatus()
	//钱包已经解锁或者错误是ErrSaveSeedFirst直接返回
	if ok || err == types.ErrSaveSeedFirst {
		return ok, err
	}
	//钱包已经锁定，挖矿锁已经解锁,需要判断addr是否是挖矿合约地址
	if wallet.isTicketLocked == false {
		if addr == account.ExecAddress("ticket").String() {
			return true, nil
		}
	}
	return ok, err
}
