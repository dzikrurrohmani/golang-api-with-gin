package main

type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"user_password"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}