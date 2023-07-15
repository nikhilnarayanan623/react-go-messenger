package routes

const (
	BaseUrl = "/api"
	// auth side
	AuthUrl             = BaseUrl + "/auth"
	SignUpUrl           = AuthUrl + "/signup"
	SignInUrl           = AuthUrl + "/login"
	GoogleSignInUrl     = AuthUrl + "google-auth"
	RenewAccessTokenUrl = AuthUrl + "/renew-access-token"

	UserBaseUrl     = BaseUrl + "/users"
	ListAllUsersUrl = "/all"

	ChatBaseUrl    = BaseUrl + "/chats"
	RecentChatsUrl = ChatBaseUrl + ""
)
