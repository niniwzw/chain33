package types

import (
	"errors"
)

var (
	ErrNoExecerInMavlKey       = errors.New("ErrNoExecerInMavlKey")
	ErrMavlKeyNotStartWithMavl = errors.New("ErrMavlKeyNotStartWithMavl")
	ErrNotFound                = errors.New("ErrNotFound")
	ErrBlockExec               = errors.New("ErrBlockExec")
	ErrCheckStateHash          = errors.New("ErrCheckStateHash")
	ErrCheckTxHash             = errors.New("ErrCheckTxHash")
	ErrReRunGenesis            = errors.New("ErrReRunGenesis")
	ErrActionNotSupport        = errors.New("ErrActionNotSupport")
	ErrQueryNotSupport         = errors.New("ErrQueryNotSupport")
	ErrChannelFull             = errors.New("ErrChannelFull")
	ErrAmount                  = errors.New("ErrAmount")
	ErrNoTicket                = errors.New("ErrNoTicket")
	ErrMinerIsStared           = errors.New("ErrMinerIsStared")
	ErrMinerNotStared          = errors.New("ErrMinerNotStared")
	ErrMinerNotClosed          = errors.New("ErrMinerNotClosed")
	ErrTicketCount             = errors.New("ErrTicketCount")
	ErrHashlockAmount          = errors.New("ErrHashlockAmount")
	ErrHashlockHash            = errors.New("ErrHashlockHash")
	ErrHashlockStatus          = errors.New("ErrHashlockStatus")
	ErrNoPeer                  = errors.New("ErrNoPeer")
	ErrExecNameNotMatch        = errors.New("ErrExecNameNotMatch")
	ErrChannelClosed           = errors.New("ErrChannelClosed")
	ErrNotMinered              = errors.New("ErrNotMinered")
	ErrTime                    = errors.New("ErrTime")
	ErrFromAddr                = errors.New("ErrFromAddr")
	ErrBlockHeight             = errors.New("ErrBlockHeight")
	ErrCoinBaseExecer          = errors.New("ErrCoinBaseExecer")
	ErrCoinBaseTxType          = errors.New("ErrCoinBaseTxType")
	ErrCoinBaseExecErr         = errors.New("ErrCoinBaseExecErr")
	ErrCoinBaseTarget          = errors.New("ErrCoinBaseTarget")
	ErrCoinbaseReward          = errors.New("ErrCoinbaseReward")
	ErrNotAllowDeposit         = errors.New("ErrNotAllowDeposit")
	ErrCoinBaseIndex           = errors.New("ErrCoinBaseIndex")
	ErrCoinBaseTicketStatus    = errors.New("ErrCoinBaseTicketStatus")
	ErrBlockNotFound           = errors.New("ErrBlockNotFound")
	ErrHashlockReturnAddrss    = errors.New("ErrHashlockReturnAddrss")
	ErrHashlockTime            = errors.New("ErrHashlockTime")
	ErrHashlockReapeathash     = errors.New("ErrHashlockReapeathash")
	ErrHashlockSendAddress     = errors.New("ErrHashlockSendAddress")
	ErrRetrieveRepeatAddress   = errors.New("ErrRetrieveRepeatAddress")
	ErrRetrieveDefaultAddress  = errors.New("ErrRetrieveDefaultAddress")
	ErrRetrievePeriodLimit     = errors.New("ErrRetrievePeriodLimit")
	ErrRetrieveAmountLimit     = errors.New("ErrRetrieveAmountLimit")
	ErrRetrieveTimeweightLimit = errors.New("ErrRetrieveTimeweightLimit")
	ErrRetrievePrepareAddress  = errors.New("ErrRetrievePrepareAddress")
	ErrRetrievePerformAddress  = errors.New("ErrRetrievePerformAddress")
	ErrRetrieveCancelAddress   = errors.New("ErrRetrieveCancelAddress")
	ErrRetrieveStatus          = errors.New("ErrRetrieveStatus")
	ErrRetrieveRelateLimit     = errors.New("ErrRetrieveRelateLimit")
	ErrRetrieveRelation        = errors.New("ErrRetrieveRelation")
	ErrRetrieveNoBalance       = errors.New("ErrRetrieveNoBalance")
	ErrLogType                 = errors.New("ErrLogType")
	ErrInvalidParam            = errors.New("ErrInvalidParameters")
	ErrInvalidAddress          = errors.New("ErrInvalidAddress")
	//err for token
	ErrTokenNameLen         = errors.New("ErrTokenNameLength")
	ErrTokenSymbolLen       = errors.New("ErrTokenSymbolLength")
	ErrTokenTotalOverflow   = errors.New("ErrTokenTotalOverflow")
	ErrTokenSymbolUpper     = errors.New("ErrTokenSymbolUpper")
	ErrTokenIntroLen        = errors.New("ErrTokenIntroductionLen")
	ErrTokenExist           = errors.New("ErrTokenSymbolExistAlready")
	ErrTokenNotPrecreated   = errors.New("ErrTokenNotPrecreated")
	ErrTokenCreatedApprover = errors.New("ErrTokenCreatedApprover")
	ErrTokenRevoker         = errors.New("ErrTokenRevoker")
	ErrTokenCanotRevoked    = errors.New("ErrTokenCanotRevokedWithWrongStatus")
	ErrTokenOwner           = errors.New("ErrTokenSymbolOwnerNotMatch")
	ErrTokenHavePrecreated  = errors.New("ErrOwnerHaveTokenPrecreateYet")
	ErrTokenBlacklist       = errors.New("ErrTokenBlacklist")
	ErrTokenNotExist        = errors.New("ErrTokenSymbolNotExist")
	//err for trade
	ErrTSellBalanceNotEnough   = errors.New("ErrTradeSellBalanceNotEnough")
	ErrTSellOrderNotExist      = errors.New("ErrTradeSellOrderNotExist")
	ErrTSellOrderNotStart      = errors.New("ErrTradeSellOrderNotStart")
	ErrTSellOrderNotEnough     = errors.New("ErrTradeSellOrderNotEnough")
	ErrTSellOrderSoldout       = errors.New("ErrTradeSellOrderSoldout")
	ErrTSellOrderRevoked       = errors.New("ErrTradeSellOrderRevoked")
	ErrTSellOrderExpired       = errors.New("ErrTradeSellOrderExpired")
	ErrTSellOrderRevoke        = errors.New("ErrTradeSellOrderRevokeNotAllowed")
	ErrTSellNoSuchOrder        = errors.New("ErrTradeSellNoSuchOrder")
	ErrTBuyOrderNotExist       = errors.New("ErrTradeBuyOrderNotExist")
	ErrTBuyOrderNotEnough      = errors.New("ErrTradeBuyOrderNotEnough")
	ErrTBuyOrderSoldout        = errors.New("ErrTradeBuyOrderSoldout")
	ErrTBuyOrderRevoked        = errors.New("ErrTradeBuyOrderRevoked")
	ErrTBuyOrderRevoke         = errors.New("ErrTradeBuyOrderRevokeNotAllowed")
	ErrTCntLessThanMinBoardlot = errors.New("ErrTradeCountLessThanMinBoardlot")

	ErrStartBigThanEnd            = errors.New("ErrStartBigThanEnd")
	ErrToAddrNotSameToExecAddr    = errors.New("ErrToAddrNotSameToExecAddr")
	ErrTypeAsset                  = errors.New("ErrTypeAsset")
	ErrEmpty                      = errors.New("ErrEmpty")
	ErrSendSameToRecv             = errors.New("ErrSendSameToRecv")
	ErrExecNameNotAllow           = errors.New("ErrExecNameNotAllow")
	ErrLocalDBPerfix              = errors.New("ErrLocalDBPerfix")
	ErrTimeout                    = errors.New("ErrTimeout")
	ErrBlockHeaderDifficulty      = errors.New("ErrBlockHeaderDifficulty")
	ErrNoTx                       = errors.New("ErrNoTx")
	ErrTxExist                    = errors.New("ErrTxExist")
	ErrManyTx                     = errors.New("ErrManyTx")
	ErrDupTx                      = errors.New("ErrDupTx")
	ErrMemFull                    = errors.New("ErrMemFull")
	ErrNoBalance                  = errors.New("ErrNoBalance")
	ErrBalanceLessThanTenTimesFee = errors.New("ErrBalanceLessThanTenTimesFee")
	ErrTxExpire                   = errors.New("ErrTxExpire")
	ErrSign                       = errors.New("ErrSign")
	ErrFeeTooLow                  = errors.New("ErrFeeTooLow")
	ErrEmptyTx                    = errors.New("ErrEmptyTx")
	ErrTxFeeTooLow                = errors.New("ErrTxFeeTooLow")
	ErrTxMsgSizeTooBig            = errors.New("ErrTxMsgSizeTooBig")
	ErrTicketClosed               = errors.New("ErrTicketClosed")
	ErrEmptyMinerTx               = errors.New("ErrEmptyMinerTx")
	ErrMinerNotPermit             = errors.New("ErrMinerNotPermit")
	ErrMinerAddr                  = errors.New("ErrMinerAddr")
	ErrModify                     = errors.New("ErrModify")
	ErrFutureBlock                = errors.New("ErrFutureBlock")
	ErrHashNotFound               = errors.New("ErrHashNotFound")
	ErrTxDup                      = errors.New("ErrTxDup")
	ErrNotSync                    = errors.New("ErrNotSync")
	ErrSize                       = errors.New("ErrSize")
	ErrMinerTx                    = errors.New("ErrMinerTx")

	// BlockChain Error Types
	ErrHashNotExist           = errors.New("ErrHashNotExist")
	ErrHeightNotExist         = errors.New("ErrHeightNotExist")
	ErrTxNotExist             = errors.New("ErrTxNotExist")
	ErrAddrNotExist           = errors.New("ErrAddrNotExist")
	ErrStartHeight            = errors.New("ErrStartHeight")
	ErrEndLessThanStartHeight = errors.New("ErrEndLessThanStartHeight")
	ErrClientNotBindQueue     = errors.New("ErrClientNotBindQueue")
	ErrContinueBack           = errors.New("ErrContinueBack")
	ErrUnmarshal              = errors.New("ErrUnmarshal")
	ErrMarshal                = errors.New("ErrMarshal")
	ErrBlockExist             = errors.New("ErrBlockExist")
	ErrParentBlockNoExist     = errors.New("ErrParentBlockNoExist")
	ErrBlockHeightNoMatch     = errors.New("ErrBlockHeightNoEqual")
	ErrParentTdNoExist        = errors.New("ErrParentTdNoExist")
	ErrBlockHashNoMatch       = errors.New("ErrBlockHashNoMatch")
	ErrIsClosed               = errors.New("ErrIsClosed")
	ErrDecode                 = errors.New("ErrDecode")
	ErrNotRollBack            = errors.New("ErrNotRollBack")
	//wallet
	ErrInputPara            = errors.New("ErrInputPara")
	ErrWalletIsLocked       = errors.New("ErrWalletIsLocked")
	ErrSaveSeedFirst        = errors.New("ErrSaveSeedFirst")
	ErrUnLockFirst          = errors.New("ErrUnLockFirst")
	ErrLabelHasUsed         = errors.New("ErrLabelHasUsed")
	ErrPrivkeyExist         = errors.New("ErrPrivkeyExist")
	ErrPrivkey              = errors.New("ErrPrivkey")
	ErrInsufficientBalance  = errors.New("ErrInsufficientBalance")
	ErrInsufficientTokenBal = errors.New("ErrInsufficientTokenBalance")
	ErrInsuffSellOrder      = errors.New("ErrInsufficientSellOrder2buy")
	ErrVerifyOldpasswdFail  = errors.New("ErrVerifyOldpasswdFail")
	ErrInputPassword        = errors.New("ErrInputPassword")
	ErrSeedlang             = errors.New("ErrSeedlang")
	ErrSeedNotExist         = errors.New("ErrSeedNotExist")
	ErrSubPubKeyVerifyFail  = errors.New("ErrSubPubKeyVerifyFail")
	ErrLabelNotExist        = errors.New("ErrLabelNotExist")
	ErrAccountNotExist      = errors.New("ErrAccountNotExist")
	ErrSeedExist            = errors.New("ErrSeedExist")
	ErrNotSupport           = errors.New("ErrNotSupport")
	ErrSeedWordNum          = errors.New("ErrSeedWordNum")
	ErrSeedWord             = errors.New("ErrSeedWord")
	ErrNoPrivKeyOrAddr      = errors.New("ErrNoPrivKeyOrAddr")
	ErrNewWalletFromSeed    = errors.New("ErrNewWalletFromSeed")
	ErrNewKeyPair           = errors.New("ErrNewKeyPair")
	ErrPrivkeyToPub         = errors.New("ErrPrivkeyToPub")

	ErrOnlyTicketUnLocked = errors.New("ErrOnlyTicketUnLocked")
	ErrNewCrypto          = errors.New("ErrNewCrypto")
	ErrFromHex            = errors.New("ErrFromHex")
	ErrPrivKeyFromBytes   = errors.New("ErrFromHex")
	ErrParentHash         = errors.New("ErrParentHash")
	// manage
	ErrNoPrivilege    = errors.New("ErrNoPrivilege")
	ErrBadConfigKey   = errors.New("ErrBadConfigKey")
	ErrBadConfigOp    = errors.New("ErrBadConfigOp")
	ErrBadConfigValue = errors.New("ErrBadConfigValue")
	//p2p
	ErrPing       = errors.New("ErrPingSignature")
	ErrVersion    = errors.New("ErrVersionNoSupport")
	ErrStreamPing = errors.New("ErrStreamPing")
	ErrPeerStop   = errors.New("ErrPeerStop")

	ErrBlockSize                  = errors.New("ErrBlockSize")
	ErrTxGroupCountLessThanTwo    = errors.New("ErrTxGroupCountLessThanTwo")
	ErrTxGroupHeader              = errors.New("ErrTxGroupHeader")
	ErrTxGroupNext                = errors.New("ErrTxGroupNext")
	ErrTxGroupCountBigThanMaxSize = errors.New("ErrTxGroupCountBigThanMaxSize")
	ErrTxGroupEmpty               = errors.New("ErrTxGroupEmpty")
	ErrTxGroupCount               = errors.New("ErrTxGroupCount")
	ErrTxGroupFeeNotZero          = errors.New("ErrTxGroupFeeNotZero")
	ErrNomalTx                    = errors.New("ErrNomalTx")
	ErrUnknowDriver               = errors.New("ErrUnknowDriver")
	ErrSymbolNameNotAllow         = errors.New("ErrSymbolNameNotAllow")
	ErrTxGroupNotSupport          = errors.New("ErrTxGroupNotSupport")
	ErrNotAllowKey                = errors.New("ErrNotAllowKey")
	ErrDataBaseDamage             = errors.New("ErrDataBaseDamage")
	ErrIndex                      = errors.New("ErrIndex")
	ErrNotAllowPubkey             = errors.New("ErrNotAllowPubkey")

	ErrRelayBalanceNotEnough   = errors.New("ErrRelaySellBalanceNotEnough")
	ErrRelayOrderNotExist      = errors.New("ErrRelayOrderNotExist")
	ErrRelayOrderOnSell        = errors.New("ErrRelayOrderOnSell")
	ErrRelayOrderParamErr      = errors.New("ErrRelayOrderParamErr")
	ErrRelayOrderSoldout       = errors.New("ErrRelayOrderSoldout")
	ErrRelayOrderRevoked       = errors.New("ErrRelayOrderRevoked")
	ErrRelayOrderConfirming    = errors.New("ErrRelayOrderConfirming")
	ErrRelayOrderFinished      = errors.New("ErrRelayOrderFinished")
	ErrRelayReturnAddr         = errors.New("ErrRelayReturnAddr")
	ErrRelayVerify             = errors.New("ErrRelayVerify")
	ErrRelayVerifyAddrNotFound = errors.New("ErrRelayVerifyAddrNotFound")
	ErrRelayWaitBlocksErr      = errors.New("ErrRelayWaitBlocks")
	ErrRelayCoinTxHashUsed     = errors.New("ErrRelayCoinTxHashUsed")
	ErrRelayBtcTxTimeErr       = errors.New("ErrRelayBtcTxTimeErr")
	ErrRelayBtcHeadSequenceErr = errors.New("ErrRelayBtcHeadSequenceErr")
	ErrRelayBtcHeadHashErr     = errors.New("ErrRelayBtcHeadHashErr")
	ErrRelayBtcHeadBitsErr     = errors.New("ErrRelayBtcHeadBitsErr")

	//authority
	ErrValidateCertFailed = errors.New("validate cert failed")
)
