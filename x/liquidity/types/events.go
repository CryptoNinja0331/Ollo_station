package types

// Event types for the liquidity module.
const (
	EventTypeCreatePair       = "create_pair"
	EventTypeCreatePool       = "create_pool"
	EventTypeCreateRangedPool = "create_ranged_pool"
	EventTypeDeposit          = "deposit"
	EventTypeWithdraw         = "withdraw"
	EventTypeLimitOrder       = "limit_order"
	EventTypeMarketOrder      = "market_order"
	EventTypeMMOrder          = "mm_order"
	EventTypeCancelOrder      = "cancel_order"
	EventTypeCancelAllOrders  = "cancel_all_orders"
	EventTypeCancelMMOrder    = "cancel_mm_order"
	EventTypeDepositResult    = "deposit_result"
	EventTypeWithdrawalResult = "withdrawal_result"
	EventTypeOrderResult      = "order_result"
	EventTypeUserOrderMatched = "user_order_matched"
	EventTypePoolOrderMatched = "pool_order_matched"

	AttributeKeyCreator            = "creator"
	AttributeKeyDepositor          = "depositor"
	AttributeKeyWithdrawer         = "withdrawer"
	AttributeKeyOrderer            = "orderer"
	AttributeKeyBaseCoinDenom      = "base_coin_denom"
	AttributeKeyQuoteCoinDenom     = "quote_coin_denom"
	AttributeKeyDepositCoins       = "deposit_coins"
	AttributeKeyAcceptedCoins      = "accepted_coins"
	AttributeKeyMintedPoolCoin     = "minted_pool_coin"
	AttributeKeyPoolCoin           = "pool_coin"
	AttributeKeyWithdrawnCoins     = "withdrawn_coins"
	AttributeKeyRefundedCoins      = "refunded_coins"
	AttributeKeyReserveAddress     = "reserve_address"
	AttributeKeyEscrowAddress      = "escrow_address"
	AttributeKeyRequestId          = "request_id"
	AttributeKeyPoolId             = "pool_id"
	AttributeKeyPairId             = "pair_id"
	AttributeKeyBatchId            = "batch_id"
	AttributeKeyOrderId            = "order_id"
	AttributeKeyOrderIds           = "order_ids"
	AttributeKeyOrderDirection     = "order_direction"
	AttributeKeyOfferCoin          = "offer_coin"
	AttributeKeyDemandCoinDenom    = "demand_coin_denom"
	AttributeKeyPrice              = "price"
	AttributeKeyAmount             = "amount"
	AttributeKeyOpenAmount         = "open_amount"
	AttributeKeyExpireAt           = "expire_at"
	AttributeKeyRemainingOfferCoin = "remaining_offer_coin"
	AttributeKeyReceivedCoin       = "received_coin"
	AttributeKeyPairIds            = "pair_ids"
	AttributeKeyCanceledOrderIds   = "canceled_order_ids"
	AttributeKeyStatus             = "status"
	AttributeKeyMatchedAmount      = "matched_amount"
	AttributeKeyPaidCoin           = "paid_coin"
)