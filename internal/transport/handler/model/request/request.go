package request

type AuthRequest struct {
	// Имя пользователя для аутентификации.
	Username string `json:"username"`
	// Пароль для аутентификации.
	Password string `json:"password"`
}

type SendCoinRequest struct {
	// Имя пользователя, которому нужно отправить монеты.
	ToUser string `json:"toUser"`
	// Количество монет, которые необходимо отправить.
	Amount int32 `json:"amount"`
}

// type BuyMerchRequest struct {
// 	// Имя пользователя, которому нужно отправить монеты.
// 	ToUser string `json:"toUser"`
// 	// Количество монет, которые необходимо отправить.
// 	Amount int32 `json:"amount"`
// }
