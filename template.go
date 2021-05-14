package main

// ClientTemplate ..
const ClientTemplate = `{
	"redirect_uris": [
	  "{{ .CallbackURL }}"
	],
	"client_name": "{{ .ClientName }}",
	"contacts": [
	  "client@iam.test"
	],
	"token_endpoint_auth_method": "client_secret_basic",
	"scope": "address phone openid email profile offline_access",
	"grant_types": [
	  "refresh_token",
	  "authorization_code"
	],
	"response_types": [
	  "code"
	]
  }`
