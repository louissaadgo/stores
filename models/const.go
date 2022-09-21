package models

const (
	TypeErrorResponse   = "error"
	TypeSuccessResponse = "success"
	TypeBannedResponse  = "banned"
	TypeAuthResponse    = "auth"
	TypeAdmin           = "admin"
	TypeMerchant        = "merchant"
	TypeUser            = "user"
	TypeAllAdmins       = "all_admins"
	TypeAllMerchants    = "all_merchants"
	TypeCurrentAdmin    = "current_admin"
	TypeAllCoupons      = "all_coupons"
	TypeAllAttributes   = "all_attributes"
	TypeAllCategories   = "all_categories"
	TypeAllCurrencies   = "all_currencies"
	TypeAllWallets      = "all_wallets"
	TypeUserCart        = "user_cart"
	PlatformMobile      = "mobile"
	PlatformWeb         = "web"
	SignTypeNative      = "native"
	SignTypeParty       = "party"

	MerchantStatusActive   = "active"
	MerchantStatusInactive = "inactive"
	MerchantStatusBanned   = "banned"

	UserStatusActive       = "active"
	UserStatusBanned       = "banned"
	ItemStatusActive       = "active"
	CouponTypeFixed        = "fixed"
	CouponTypePercentage   = "percentage"
	OrderStatusProccessing = "proccessing"
	PaymentWallet          = "wallet"
	PaymentCOD             = "cash on delivery"
	ItemPaid               = "paid"
	ItemUnpaid             = "unpaid"
)
