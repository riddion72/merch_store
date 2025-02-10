package response

type AuthResponse struct {
	// JWT-токен для доступа к защищенным ресурсам.
	Token string `json:"token,omitempty"`
}

type InfoResponse struct {
	// Количество доступных монет.
	Coins int32 `json:"coins,omitempty"`

	Inventory []InfoResponseInventory `json:"inventory,omitempty"`

	CoinHistory *InfoResponseCoinHistory `json:"coinHistory,omitempty"`
}

type InfoResponseInventory struct {
	// Тип предмета.
	Type_ string `json:"type,omitempty"`
	// Количество предметов.
	Quantity int32 `json:"quantity,omitempty"`
}

type InfoResponseCoinHistory struct {
	Received []InfoResponseCoinHistoryReceived `json:"received,omitempty"`

	Sent []InfoResponseCoinHistorySent `json:"sent,omitempty"`
}

type InfoResponseCoinHistorySent struct {
	// Имя пользователя, которому отправлены монеты.
	ToUser string `json:"toUser,omitempty"`
	// Количество отправленных монет.
	Amount int32 `json:"amount,omitempty"`
}

type InfoResponseCoinHistoryReceived struct {
	// Имя пользователя, который отправил монеты.
	FromUser string `json:"fromUser,omitempty"`
	// Количество полученных монет.
	Amount int32 `json:"amount,omitempty"`
}

type ErrorResponse struct {
	// Сообщение об ошибке, описывающее проблему.
	Errors string `json:"errors,omitempty"`
}
