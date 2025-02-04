package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rob-bender/invest-market-backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api-v1")
	{
		api.GET("/user/exchangeRates", h.exchangeRates)                          // получение обменных курсов
		api.POST("/user/checkAuth", h.checkAuth)                                 // проверка на регистрацию пользователя
		api.POST("/user/registration", h.registrationUser)                       // Регистрация пользователя
		api.POST("/user/getUserLanguage", h.getUserLanguage)                     // Получить выбранный язык пользователя
		api.POST("/user/updateLanguage", h.updateLanguage)                       // Изменение языка пользователя
		api.POST("/user/getUserCurrency", h.getUserCurrency)                     // Получить выбранную валюту пользователя
		api.POST("/user/updateCurrency", h.updateCurrency)                       // Изменить валюту пользователя
		api.POST("/user/checkIsCaptcha", h.checkIsCaptcha)                       // проверить на значение капчи
		api.POST("/user/checkIsTerms", h.checkIsTerms)                           // проверка на пользовательское соглашение
		api.POST("/user/agreeTerms", h.agreeTerms)                               // согласие пользователя с пользовательским соглашением
		api.POST("/user/getUserProfile", h.getUserProfile)                       // получение профиля пользователя
		api.POST("/user/getUserMinPrice", h.getUserMinPrice)                     // получение минималки пользователя
		api.POST("/user/getAdminByUser", h.getAdminByUser)                       // получить закрепленного за пользователем администратора
		api.POST("/user/getUserBalance", h.getUserBalance)                       // узнать баланс пользователя
		api.POST("/user/checkUserToken", h.checkUserToken)                       // проверить есть ли у пользователя данный токен
		api.POST("/user/buyUserToken", h.buyUserToken)                           // покупка токена пользователем
		api.POST("/user/sellUserToken", h.sellUserToken)                         // продажа токена пользователем
		api.POST("/user/getUserPaymentEvent", h.getUserPaymentEvent)             // получить событие продажи по uid
		api.POST("/user/getUserNft", h.getUserNft)                               // получить nft пользователя
		api.POST("/user/createWithDrawEvent", h.createWithDrawEvent)             // создание события вывода пользователем
		api.POST("/user/getWithDrawEvent", h.getWithDrawEvent)                   // получить событие вывода пользователем
		api.POST("/user/updateTradingAsset", h.updateTradingAsset)               // обновить значения выбранного актива для трейдинга
		api.POST("/user/updateTradingInvestPrice", h.updateTradingInvestPrice)   // обновить значение инвестирования пользователя
		api.POST("/user/updateTradingWaitTime", h.updateTradingWaitTime)         // обновить значение время ожидания при трейдинге
		api.POST("/user/updateTradingMovePrice", h.updateTradingMovePrice)       // обновить значение куда пойдет цена при трейдинге
		api.POST("/user/getUserTradingParams", h.getUserTradingParams)           // получить данные для трейдинга
		api.POST("/user/getUserTotalEarn", h.getUserTotalEarn)                   // получить количество всего заработано
		api.POST("/user/changeUserTotalEarn", h.changeUserTotalEarn)             // изменить количество всего заработано
		api.POST("/user/getUserCurrentLuck", h.getUserCurrentLuck)               // получить удачу пользователя выбранная администратором
		api.GET("/payment/getAll", h.getAllPayments)                             // получение всех платёжек
		api.POST("/payment/createPayment", h.createPayment)                      // создание платёжки
		api.GET("/collection/getAll", h.getAllCollections)                       // получение всех коллекций
		api.POST("/collection/createCollection", h.createCollection)             // создание коллекции
		api.POST("/collection/createToken", h.createToken)                       // создание токена для коллекции
		api.POST("/collection/getAllTokensCollection", h.getAllTokensCollection) // получение всех токенов коллекции
		api.POST("/collection/getToken", h.getToken)                             // получение токена
		api.POST("/admin/checkIsAdmin", h.checkIsAdmin)                          //  проверка на администратора
		api.POST("/admin/createReferral", h.createReferral)                      // создание рефералки
		api.POST("/admin/checkUserReferral", h.checkUserReferral)                // получить количество пользователей по рефералке администратора
		api.POST("/admin/getUsersReferral", h.getUsersReferral)                  // получение всех пользователей по рефералке администратора
		api.POST("/admin/getUserReferral", h.getUserReferral)                    // получение пользователя по рефералке администратора
		api.POST("/admin/getUserProfile", h.adminGetUserProfile)                 // получение профиля пользователя
		api.POST("/admin/checkIsPremium", h.checkIsPremium)                      // проверка пользователя на премиум
		api.POST("/admin/updatePremium", h.updatePremium)                        // изменить премиум
		api.POST("/admin/checkIsVerified", h.checkIsVerification)                // проверка на верификацию пользователя
		api.POST("/admin/updateVerification", h.updateVerification)              // изменить верификацию
		api.POST("/admin/adminUpdateMinimPrice", h.adminUpdateMinimPrice)        // изменить минималку у администратора для новых пользователей, которые перейдут по ссылке
		api.POST("/admin/adminAddBalance", h.adminAddBalance)                    // пополнить баланс пользователя
		api.POST("/admin/adminChangeMinUser", h.adminChangeMinUser)              // изменить минималку для пользователя
		api.POST("/admin/adminChangeBalance", h.adminChangeBalance)              // изменение баланса пользователя
		api.POST("/admin/checkIsBlockUser", h.checkIsBlockUser)                  // проверка на блокировку пользователя
		api.POST("/admin/checkIsVisibleName", h.checkIsVisibleName)              // проверка на видимость имени администратора
		api.POST("/admin/changeVisibleName", h.adminChangeVisibleName)           // показать имя администратора при отправке количества заработка или скрыть его
		api.POST("/admin/blockUser", h.adminBlockUser)                           // проверка на блокировку пользователя
		api.POST("/admin/adminBuyTokenUser", h.adminBuyTokenUser)                // покупка администратором токена пользователя
		api.POST("/depot/createDepot", h.createDepot)                            // создание депа
		api.POST("/admin/adminWithdrawApprove", h.adminWithdrawApprove)          // согласие администратора с выводом денег пользователя
		api.POST("/admin/adminWithdrawRefuse", h.adminWithdrawRefuse)            // отказ администратора с выводом денег пользователя
		api.POST("/admin/adminChangeLuckUser", h.adminChangeLuckUser)            // изменить удачу пользователя
	}

	return router
}
