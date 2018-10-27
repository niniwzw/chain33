package types

import (
	"errors"
)

var (
	ErrMethodReturnType        = errors.New("ErrMethodReturnType")
	ErrMethodNotFound          = errors.New("ErrMethodNotFound")
	ErrExecBlockNil            = errors.New("ErrExecBlockNil")
	ErrNotAllow                = errors.New("ErrNotAllow")
	ErrCanOnlyDelTopVersion    = errors.New("ErrCanOnlyDelTopVersion")
	ErrPrevVersion             = errors.New("ErrPrevVersion")
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
	ErrInvalidParam            = errors.New("ErrInvalidParam")
	ErrInvalidAddress          = errors.New("ErrInvalidAddress")
	ErrStateHashLost           = errors.New("ErrStateHashLost")
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
	//err for privacy
	ErrWrongKey     = errors.New("ErrWrongKey")
	ErrInvalidOrder = errors.New("ErrInvalidOrder")

	// err for blackwhite game
	ErrIncorrectStatus  = errors.New("ErrIncorrectStatus")
	ErrRepeatPlayerAddr = errors.New("ErrRepeatPlayerAddress")
	ErrNoTimeoutDone    = errors.New("ErrNoTimeoutDone")
	ErrNoExistAddr      = errors.New("ErrNoExistAddress")
	ErrNoLoopSeq        = errors.New("ErrBlackwhiteFinalloopLessThanSeq")

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
	ErrHeaderNotSet               = errors.New("ErrHeaderNotSet")
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
	ErrOpenTicketPubHash          = errors.New("ErrOpenTicketPubHash")

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
	ErrNoUTXORec4Token        = errors.New("ErrNoUTXORec4Token")
	ErrNoUTXORec4Amount       = errors.New("ErrNoUTXORec4Amount")
	ErrNotEnoughUTXOs         = errors.New("ErrNotEnoughUTXOs")
	ErrNoSuchPrivacyTX        = errors.New("ErrNoSuchPrivacyTX")
	ErrDoubleSpendOccur       = errors.New("ErrDoubleSpendOccur")
	ErrOutputIndex            = errors.New("ErrOutputIndex")
	ErrPubkeysOfUTXO          = errors.New("ErrPubkeysOfUTXO")
	ErrRecoverUTXO            = errors.New("ErrRecoverUTXO")
	ErrPeerInfoIsNil          = errors.New("ErrPeerInfoIsNil")
	//wallet
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
	ErrPubKeyLen            = errors.New("ErrPublicKeyLen")
	ErrPrivateKeyLen        = errors.New("ErrPrivateKeyLen")
	ErrSeedWord             = errors.New("ErrSeedWord")
	ErrNoPrivKeyOrAddr      = errors.New("ErrNoPrivKeyOrAddr")
	ErrNewWalletFromSeed    = errors.New("ErrNewWalletFromSeed")
	ErrNewKeyPair           = errors.New("ErrNewKeyPair")
	ErrPrivkeyToPub         = errors.New("ErrPrivkeyToPub")

	ErrOnlyTicketUnLocked    = errors.New("ErrOnlyTicketUnLocked")
	ErrNewCrypto             = errors.New("ErrNewCrypto")
	ErrFromHex               = errors.New("ErrFromHex")
	ErrPrivKeyFromBytes      = errors.New("ErrFromHex")
	ErrParentHash            = errors.New("ErrParentHash")
	ErrPrivacyNotEnabled     = errors.New("ErrPrivacyNotEnabled")
	ErrPrivacyTxFeeNotEnough = errors.New("ErrPrivacyTxFeeNotEnough")
	ErrRescanFlagScaning     = errors.New("ErrRescanFlagScaning")

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
	ErrTxGroupIndex               = errors.New("ErrTxGroupIndex")
	ErrTxGroupFormat              = errors.New("ErrTxGroupFormat")
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
	ErrNotAllowMemSetKey          = errors.New("ErrNotAllowMemSetKey")
	ErrDataBaseDamage             = errors.New("ErrDataBaseDamage")
	ErrIndex                      = errors.New("ErrIndex")
	// ring signature
	ErrGeFromBytesVartime = errors.New("ErrGeFromBytesVartime")
	ErrNotAllowPubkey     = errors.New("ErrNotAllowPubkey")

	ErrRelayBalanceNotEnough   = errors.New("ErrRelaySellBalanceNotEnough")
	ErrRelayOrderNotExist      = errors.New("ErrRelayOrderNotExist")
	ErrRelayOrderOnSell        = errors.New("ErrRelayOrderOnSell")
	ErrRelayOrderStatusErr     = errors.New("ErrRelayOrderStatusErr")
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
	ErrRelayBtcHeadNewBitsErr  = errors.New("ErrRelayBtcHeadNewBitsErr")

	//game err
	ErrGameCreateAmount = errors.New("You fill in more than the maximum number of games.")
	ErrGameCancleAddr   = errors.New("You don't have permission to cancel someone else's game.")
	ErrGameCloseAddr    = errors.New("The game time has not yet expired,You don't have permission to call yet.")
	ErrGameTimeOut      = errors.New("The game has expired.,You don't have permission to call.")
	ErrGameMatchStatus  = errors.New("can't join the game, the game has matched or finished!")
	ErrGameMatch        = errors.New("can't join the game, You can't match the game you created!")
	ErrGameCancleStatus = errors.New("can't cancle the game, the game has matched!")
	ErrGameCloseStatus  = errors.New("can't close the game again, the game has  finished!")

	//authority
	ErrValidateCertFailed  = errors.New("ErrValidateCertFailed")
	ErrGetHistoryCertData  = errors.New("ErrGetHistoryCertData")
	ErrUnknowAuthSignType  = errors.New("ErrUnknowAuthSignType")
	ErrInitializeAuthority = errors.New("ErrInitializeAuthority")
	//rpc
	ErrInvalidMainnetRpcAddr = errors.New("ErrInvalidMainnetRpcAddr")

	// executor.paracross
	ErrInvalidTitle         = errors.New("ErrInvalidTitle")
	ErrTitleNotExist        = errors.New("ErrTitleNotExist")
	ErrNodeNotForTheTitle   = errors.New("ErrNodeNotForTheTitle")
	ErrParaBlockHashNoMatch = errors.New("ErrParaBlockHashNoMatch")
	ErrTxGroupParaCount     = errors.New("ErrTxGroupParaCount")
	ErrParaMinerBaseIndex   = errors.New("ErrParaMinerBaseIndex")
	ErrParaMinerTxType      = errors.New("ErrParaMinerTxType")
	ErrParaEmptyMinerTx     = errors.New("ErrParaEmptyMinerTx")
	ErrParaMinerExecErr     = errors.New("ErrParaMinerExecErr")

	//lottery
	ErrLotteryStatus            = errors.New("ErrLotteryStatus")
	ErrLotteryDrawActionInvalid = errors.New("ErrLotteryDrawActionInvalid")
	ErrLotteryFundNotEnough     = errors.New("ErrLotteryFundNotEnough")
	ErrLotteryCreatorBuy        = errors.New("ErrLotteryCreatorBuy")
	ErrLotteryBuyAmount         = errors.New("ErrLotteryBuyAmount")
	ErrLotteryRepeatHash        = errors.New("ErrLotteryRepeatHash")
	ErrLotteryPurBlockLimit     = errors.New("ErrLotteryPurBlockLimit")
	ErrLotteryDrawBlockLimit    = errors.New("ErrLotteryDrawBlockLimit")
	ErrLotteryBuyNumber         = errors.New("ErrLotteryBuyNumber")
	ErrLotteryShowRepeated      = errors.New("ErrLotteryShowRepeated")
	ErrLotteryShowError         = errors.New("ErrLotteryShowError")
	ErrLotteryErrLuckyNum       = errors.New("ErrLotteryErrLuckyNum")
	ErrLotteryErrCloser         = errors.New("ErrLotteryErrCloser")
	ErrLotteryErrUnableClose    = errors.New("ErrLotteryErrUnableClose")
	ErrNodeNotExist             = errors.New("ErrNodeNotExist")
)
