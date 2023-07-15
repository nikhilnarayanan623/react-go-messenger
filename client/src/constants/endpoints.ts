export const API_BASE_URL = "http://localhost:8080/api"

const END_POINTS = {
    SIGN_UP:`${API_BASE_URL}/auth/signup`,
    SIGN_IN:`${API_BASE_URL}/auth/login`,
    GET_ALL_USERS:`${API_BASE_URL}/users/all `
}

export default END_POINTS