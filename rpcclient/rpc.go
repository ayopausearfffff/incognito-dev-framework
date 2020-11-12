package rpcclient

import (
	"fmt"
	"incognito-dev-framework/account"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/common/base58"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/rpcserver/jsonresult"
	"github.com/incognitochain/incognito-chain/wallet"
)

type RPCClient struct {
	client ClientInterface
}

func NewRPCClient(client ClientInterface) *RPCClient {
	rpc := &RPCClient{
		client: client,
	}
	return rpc
}

func (r *RPCClient) API_SendTxPRV(privateKey string, receivers map[string]interface{}, fee float64, privacy float64) (*jsonresult.CreateTransactionResult, error) {
	result, err := r.client.CreateAndSendTransaction(privateKey, receivers, fee, privacy)
	return &result, err
}

func (r *RPCClient) API_SendTxCreateCustomToken(privateKey string, receiverPaymentAddress string, privacy float64, tokenName string, tokenSymbol string, tokenAmount uint64) (*jsonresult.CreateTransactionTokenResult, error) {
	result, err := r.client.CreateAndSendPrivacyCustomTokenTransaction(privateKey, nil, -1, privacy, map[string]interface{}{
		"Privacy":     true,
		"TokenID":     "",
		"TokenName":   tokenName,
		"TokenSymbol": tokenSymbol,
		"TokenFee":    float64(0),
		"TokenTxType": float64(0),
		"TokenAmount": float64(tokenAmount),
		"TokenReceivers": map[string]interface{}{
			receiverPaymentAddress: float64(tokenAmount),
		},
	}, "", privacy)
	return &result, err
}

func (r *RPCClient) API_SendTxCustomToken(privateKey string, tokenID string, receivers map[string]interface{}, fee float64, privacy float64) (*jsonresult.CreateTransactionTokenResult, error) {
	result, err := r.client.CreateAndSendPrivacyCustomTokenTransaction(privateKey, receivers, fee, privacy, map[string]interface{}{
		"Privacy":        true,
		"TokenID":        tokenID,
		"TokenName":      "",
		"TokenSymbol":    "",
		"TokenFee":       float64(0),
		"TokenTxType":    float64(1),
		"TokenAmount":    float64(0),
		"TokenReceivers": receivers,
	}, "", privacy)
	return &result, err
}

func (r *RPCClient) API_SendTxWithWithdrawalReqV2(privateKey string, receivers map[string]interface{}, fee float64, privacy float64, reqInfo map[string]interface{}) (*jsonresult.CreateTransactionResult, error) {
	result, err := r.client.CreateAndSendTxWithWithdrawalReqV2(privateKey, receivers, fee, privacy, reqInfo)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPDEFeeWithdrawalReq(privateKey string, receivers map[string]interface{}, fee float64, privacy float64, reqInfo map[string]interface{}) (*jsonresult.CreateTransactionResult, error) {
	result, err := r.client.CreateAndSendTxWithPDEFeeWithdrawalReq(privateKey, receivers, fee, privacy, reqInfo)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPTokenTradeReq(privateKey string, receivers map[string]interface{}, fee float64, privacy float64, reqInfo map[string]interface{}) (*jsonresult.CreateTransactionTokenResult, error) {
	result, err := r.client.CreateAndSendTxWithPTokenTradeReq(privateKey, receivers, fee, privacy, reqInfo, "", 0)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPTokenCrossPoolTradeReq(acount account.Account, tokenID string, buyTokenID string, amount string) (*jsonresult.CreateTransactionTokenResult, error) {
	burnAddr, err := r.client.GetBurningAddress(float64(0))
	if err != nil {
		return nil, err
	}
	reqInfo := map[string]interface{}{
		"Privacy":     true,
		"TokenID":     tokenID,
		"TokenTxType": float64(1),
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount,
		"TokenReceivers": map[string]interface{}{
			burnAddr: amount,
		},
		"TokenFee":            "0",
		"TokenIDToBuyStr":     buyTokenID,
		"TokenIDToSellStr":    tokenID,
		"SellAmount":          amount,
		"MinAcceptableAmount": "1",
		"TradingFee":          "1",
		"TraderAddressStr":    acount.PaymentAddress,
	}
	result, err := r.client.CreateAndSendTxWithPTokenCrossPoolTradeReq(acount.PrivateKey, map[string]interface{}{burnAddr: "1"}, -1, 0, reqInfo, "", 0)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPRVTradeReq(privateKey string, receivers map[string]interface{}, fee float64, privacy float64, reqInfo map[string]interface{}) (*jsonresult.CreateTransactionResult, error) {
	result, err := r.client.CreateAndSendTxWithPRVTradeReq(privateKey, receivers, fee, privacy, reqInfo)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPRVCrossPoolTradeReq(account account.Account, buyTokenID string, amount string) (*jsonresult.CreateTransactionResult, error) {
	burnAddr, err := r.client.GetBurningAddress(float64(0))
	if err != nil {
		return nil, err
	}
	reqInfo := map[string]interface{}{
		"TokenIDToBuyStr":     buyTokenID,
		"TokenIDToSellStr":    "0000000000000000000000000000000000000000000000000000000000000004",
		"SellAmount":          amount,
		"MinAcceptableAmount": "1",
		"TradingFee":          "0",
		"TraderAddressStr":    account.PaymentAddress,
	}
	result, err := r.client.CreateAndSendTxWithPRVCrossPoolTradeReq(account.PrivateKey, map[string]interface{}{
		burnAddr: amount,
	}, -1, -1, reqInfo)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPTokenContributionV2(account account.Account, tokenID string, tokenAmount string, pairID string) (*jsonresult.CreateTransactionTokenResult, error) {
	burnAddr, err := r.client.GetBurningAddress(float64(0))
	if err != nil {
		return nil, err
	}
	reqInfo := map[string]interface{}{
		"Privacy":     true,
		"TokenID":     tokenID,
		"TokenTxType": float64(1),
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": tokenAmount,
		"TokenReceivers": map[string]interface{}{
			burnAddr: tokenAmount,
		},
		"TokenFee":              "0",
		"PDEContributionPairID": pairID,
		"ContributorAddressStr": account.PaymentAddress,
		"ContributedAmount":     tokenAmount,
		"TokenIDStr":            tokenID,
	}
	result, err := r.client.CreateAndSendTxWithPTokenContributionV2(account.PrivateKey, nil, -1, 0, reqInfo, "", 0)
	return &result, err
}
func (r *RPCClient) API_SendTxWithPRVContributionV2(account account.Account, prvAmount string, pairID string) (*jsonresult.CreateTransactionResult, error) {
	burnAddr, err := r.client.GetBurningAddress(float64(0))
	if err != nil {
		return nil, err
	}
	reqInfo := map[string]interface{}{
		"PDEContributionPairID": pairID,
		"ContributorAddressStr": account.PaymentAddress,
		"ContributedAmount":     prvAmount,
		"TokenIDStr":            "0000000000000000000000000000000000000000000000000000000000000004",
	}
	result, err := r.client.CreateAndSendTxWithPRVContributionV2(account.PrivateKey, map[string]interface{}{burnAddr: prvAmount}, -1, 0, reqInfo)
	return &result, err
}
func (r *RPCClient) API_GetPDEState(beaconHeight float64) (jsonresult.CurrentPDEState, error) {
	result, err := r.client.GetPDEState(map[string]interface{}{"BeaconHeight": beaconHeight})
	return result, err
}

func (r *RPCClient) API_GetBeaconBestState() (jsonresult.GetBeaconBestState, error) {
	result, err := r.client.GetBeaconBestState()
	return result, err
}

func (r *RPCClient) API_GetShardBestState(sid int) (jsonresult.GetShardBestState, error) {
	result, err := r.client.GetShardBestState(sid)
	return result, err
}

func (r *RPCClient) API_GetTransactionHash(h string) (jsonresult.TransactionDetail, error) {
	result, err := r.client.GetTransactionByHash(h)
	return *result, err
}

func (r *RPCClient) API_GetPrivacyCustomToken(h string) (*jsonresult.GetCustomToken, error) {
	result, err := r.client.GetPrivacyCustomToken(h)
	return result, err
}

func (r *RPCClient) API_GetBalance(acc account.Account) (map[string]uint64, error) {
	tokenList := make(map[string]uint64)
	prv, _ := r.client.GetBalanceByPrivateKey(acc.PrivateKey)
	tokenList["PRV"] = prv

	tokenBL, _ := r.client.GetListPrivacyCustomTokenBalance(acc.PrivateKey)
	for _, token := range tokenBL.ListCustomTokenBalance {
		tokenList[token.Name] = token.Amount

	}
	return tokenList, nil
}

const (
	stakeShardAmount   int = 1750000000000
	stakeBeaceonAmount int = stakeShardAmount * 3
)

type StakingTxParam struct {
	CommitteeKey *incognitokey.CommitteePublicKey
	BurnAddr     string
	StakerPrk    string
	MinerPrk     string
	RewardAddr   string
	StakeShard   bool
	AutoRestake  bool
	Name         string
}

type StopStakingParam struct {
	BurnAddr  string
	StakerPrk string
	MinerPrk  string
}

func (r *RPCClient) API_SendTxStaking(stakeMeta StakingTxParam) (*jsonresult.CreateTransactionResult, error) {
	stakeAmount := 0
	stakingType := 0
	if stakeMeta.StakeShard {
		stakeAmount = stakeShardAmount
		stakingType = 63
	} else {
		stakeAmount = stakeBeaceonAmount
		stakingType = 64
	}

	if stakeMeta.RewardAddr == "" {
		wl, err := wallet.Base58CheckDeserialize(stakeMeta.StakerPrk)
		if err != nil {
			return nil, err
		}
		stakeMeta.RewardAddr = wl.Base58CheckSerialize(wallet.PaymentAddressType)
	}

	if stakeMeta.MinerPrk == "" {
		stakeMeta.MinerPrk = stakeMeta.StakerPrk
	}
	wl, err := wallet.Base58CheckDeserialize(stakeMeta.MinerPrk)
	if err != nil {
		return nil, err
	}
	privateSeedBytes := common.HashB(common.HashB(wl.KeySet.PrivateKey))
	privateSeed := base58.Base58Check{}.Encode(privateSeedBytes, common.Base58Version)
	minerPayment := wl.Base58CheckSerialize(wallet.PaymentAddressType)

	candidateWallet, err := wallet.Base58CheckDeserialize(minerPayment)
	if err != nil || candidateWallet == nil {
		fmt.Println(stakeMeta.MinerPrk, wl.KeySet.PaymentAddress, minerPayment)
		fmt.Println(err, candidateWallet)
		panic(0)
	}
	burnAddr := stakeMeta.BurnAddr
	txResp, err := r.client.CreateAndSendStakingTransaction(stakeMeta.StakerPrk, map[string]interface{}{burnAddr: float64(stakeAmount)}, 1, 0, map[string]interface{}{
		"StakingType":                  float64(stakingType),
		"CandidatePaymentAddress":      minerPayment,
		"PrivateSeed":                  privateSeed,
		"RewardReceiverPaymentAddress": stakeMeta.RewardAddr,
		"AutoReStaking":                stakeMeta.AutoRestake,
	})

	if err != nil {
		return nil, err
	}
	return &txResp, nil
}

func (r *RPCClient) API_SendTxStopAutoStake(stopStakeMeta StopStakingParam) (*jsonresult.CreateTransactionResult, error) {
	if stopStakeMeta.MinerPrk == "" {
		stopStakeMeta.MinerPrk = stopStakeMeta.StakerPrk
	}
	wl, err := wallet.Base58CheckDeserialize(stopStakeMeta.MinerPrk)
	if err != nil {
		return nil, err
	}
	privateSeedBytes := common.HashB(common.HashB(wl.KeySet.PrivateKey))
	privateSeed := base58.Base58Check{}.Encode(privateSeedBytes, common.Base58Version)
	minerPayment := wl.Base58CheckSerialize(wallet.PaymentAddressType)

	burnAddr := stopStakeMeta.BurnAddr

	txResp, err := r.client.CreateAndSendStopAutoStakingTransaction(stopStakeMeta.StakerPrk, map[string]interface{}{burnAddr: float64(0)}, 1, 0, map[string]interface{}{
		"StopAutoStakingType":     float64(127),
		"CandidatePaymentAddress": minerPayment,
		"PrivateSeed":             privateSeed,
	})
	if err != nil {
		return nil, err
	}
	return &txResp, nil
}

func (r *RPCClient) API_GetRewardAmount(paymentAddress string) (map[string]float64, error) {
	result := make(map[string]float64)
	rs, err := r.client.GetRewardAmount(paymentAddress)
	if err != nil {
		return nil, err
	}
	for token, amount := range rs {
		result[token] = float64(amount) / 1e9
	}
	return result, nil
}

func (r *RPCClient) API_SendTxWithdrawReward(privateKey string, paymentAddress string) (*jsonresult.CreateTransactionResult, error) {

	txResp, err := r.client.WithdrawReward(privateKey, nil, 0, 0, map[string]interface{}{
		"PaymentAddress": paymentAddress, "TokenID": "0000000000000000000000000000000000000000000000000000000000000004", "Version": 0,
	})
	if err != nil {
		return nil, err
	}
	return &txResp, nil
}
