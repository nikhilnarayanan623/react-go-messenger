export interface RecentlyChattedFriends {
    chat_id: number
    first_name: string
    user_name: string
    profile_picture: string
    last_message: string
    last_message_at: string
  }
  
  export interface Message {
    message_id: number
    is_current_user: boolean
    content: string
    created_at: string|number
  }
  