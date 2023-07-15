package routes

const (
	BaseUrl = "/api"

	// auth
	AuthUrl             = BaseUrl + "/auth"
	SignUpUrl           = AuthUrl + "/sign-up"
	SignInUrl           = AuthUrl + "/sign-in"
	GoogleSignInUrl     = AuthUrl + "/google-sign-in"
	RenewAccessTokenUrl = AuthUrl + "/renew-access-token"

	// user
	UserBaseUrl     = BaseUrl + "/users"
	ListAllUsersUrl = UserBaseUrl

	// chat
	ChatBaseUrl        = BaseUrl + "/chats"
	RecentChatsUrl     = ChatBaseUrl
	CreateChatUrl      = ChatBaseUrl
	ListAllMessagesUrl = ChatBaseUrl + "/:chat_id/messages"
)
