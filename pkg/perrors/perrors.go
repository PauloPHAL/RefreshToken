package perrors

import "errors"

var (
	// Erros de Usuário/Entidade
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")

	// Erros de Autenticação e Token
	ErrInvalidPassword      = errors.New("invalid password")
	ErrInvalidToken         = errors.New("invalid token")
	ErrTokenExpired         = errors.New("token expired")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
	ErrRefreshTokenRevoked  = errors.New("refresh token revoked")

	// Erros de Validação
	ErrNameRequired       = errors.New("name is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrInvalidEmail       = errors.New("email format is invalid")
	ErrPasswordRequired   = errors.New("password is required")
	ErrInvalidCredentials = errors.New("invalid credentials")

	// Validação de Nome
	ErrNameTooShort = errors.New("name must be at least 2 characters long")
	ErrNameTooLong  = errors.New("name is too long (max 100 characters)")

	// Validação de Email
	ErrEmailTooLong = errors.New("email is too long (max 254 characters)")

	// Validação de Senha
	ErrPasswordTooShort = errors.New("password must be at least 6 characters long")
	ErrPasswordTooLong  = errors.New("password is too long (max 128 characters)")

	// Erros de Token/JWT
	ErrTokenGeneration     = errors.New("failed to generate token")
	ErrInvalidTokenClaims  = errors.New("invalid token claims")
	ErrInvalidTokenSubject = errors.New("invalid subject in token")
	ErrUnexpectedSigning   = errors.New("unexpected signing method")
)
