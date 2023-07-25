package routes

const (
	BaseURL = "/api"

	// Authentication
	AuthURL             = BaseURL + "/auth"
	SignUpURL           = AuthURL + "/sign-up"
	SignInURL           = AuthURL + "/sign-in"
	GoogleSignInURL     = AuthURL + "/google-sign-in"
	RenewAccessTokenURL = AuthURL + "/renew-access-token"

	// Websockets
	SocketURL = BaseURL + "/ws"

	// Users
	UsersURL = BaseURL + "/users"

	// Friends
	UserFriendsURL        = UsersURL + "friends"
	UserFriendsRequestUrl = UserFriendsURL + "/request"

	// Chats
	ChatsURL       = BaseURL + "/chats"
	ChatsMessgeURL = ChatsURL + "/:chat_id/messages"
)
