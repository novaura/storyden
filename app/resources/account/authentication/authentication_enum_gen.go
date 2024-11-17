// Code generated by enumerator. DO NOT EDIT.

package authentication

import (
	"database/sql/driver"
	"fmt"
)

type Mode struct {
	v modeEnum
}

var (
	ModeHandle = Mode{modeHandle}
	ModeEmail  = Mode{modeEmail}
	ModePhone  = Mode{modePhone}
)

func (r Mode) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	case 'v':
		switch r {
		case ModeHandle:
			fmt.Fprint(f, "Username (default)")
		case ModeEmail:
			fmt.Fprint(f, "Email address")
		case ModePhone:
			fmt.Fprint(f, "Phone number")
		default:
			fmt.Fprint(f, "")
		}
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r Mode) String() string {
	return string(r.v)
}
func (r Mode) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *Mode) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewMode(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r Mode) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *Mode) Scan(__iNpUt__ any) error {
	s, err := NewMode(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewMode(__iNpUt__ string) (Mode, error) {
	switch __iNpUt__ {
	case string(modeHandle):
		return ModeHandle, nil
	case string(modeEmail):
		return ModeEmail, nil
	case string(modePhone):
		return ModePhone, nil
	default:
		return Mode{}, fmt.Errorf("invalid value for type 'Mode': '%s'", __iNpUt__)
	}
}

type Service struct {
	v serviceEnum
}

var (
	ServiceUsernamePassword = Service{serviceUsernamePassword}
	ServiceEmailPassword    = Service{serviceEmailPassword}
	ServiceEmailVerify      = Service{serviceEmailVerify}
	ServicePhoneVerify      = Service{servicePhoneVerify}
	ServiceWebAuthn         = Service{serviceWebAuthn}
	ServiceOAuthGoogle      = Service{serviceOAuthGoogle}
	ServiceOAuthGitHub      = Service{serviceOAuthGitHub}
	ServiceOAuthLinkedin    = Service{serviceOAuthLinkedin}
)

func (r Service) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	case 'v':
		switch r {
		case ServiceUsernamePassword:
			fmt.Fprint(f, "User/email + password")
		case ServiceEmailPassword:
			fmt.Fprint(f, "Email + password")
		case ServiceEmailVerify:
			fmt.Fprint(f, "Email + verification code")
		case ServicePhoneVerify:
			fmt.Fprint(f, "Phone number + verification code")
		case ServiceWebAuthn:
			fmt.Fprint(f, "WebAuthn/Passkey")
		case ServiceOAuthGoogle:
			fmt.Fprint(f, "Google")
		case ServiceOAuthGitHub:
			fmt.Fprint(f, "GitHub")
		case ServiceOAuthLinkedin:
			fmt.Fprint(f, "LinkedIn")
		default:
			fmt.Fprint(f, "")
		}
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r Service) String() string {
	return string(r.v)
}
func (r Service) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *Service) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewService(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r Service) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *Service) Scan(__iNpUt__ any) error {
	s, err := NewService(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewService(__iNpUt__ string) (Service, error) {
	switch __iNpUt__ {
	case string(serviceUsernamePassword):
		return ServiceUsernamePassword, nil
	case string(serviceEmailPassword):
		return ServiceEmailPassword, nil
	case string(serviceEmailVerify):
		return ServiceEmailVerify, nil
	case string(servicePhoneVerify):
		return ServicePhoneVerify, nil
	case string(serviceWebAuthn):
		return ServiceWebAuthn, nil
	case string(serviceOAuthGoogle):
		return ServiceOAuthGoogle, nil
	case string(serviceOAuthGitHub):
		return ServiceOAuthGitHub, nil
	case string(serviceOAuthLinkedin):
		return ServiceOAuthLinkedin, nil
	default:
		return Service{}, fmt.Errorf("invalid value for type 'Service': '%s'", __iNpUt__)
	}
}

type TokenType struct {
	v tokenTypeEnum
}

var (
	TokenTypeNone     = TokenType{tokenTypeNone}
	TokenTypePassword = TokenType{tokenTypePassword}
	TokenTypeWebAuthn = TokenType{tokenTypeWebAuthn}
	TokenTypeOAuth    = TokenType{tokenTypeOAuth}
)

func (r TokenType) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	case 'v':
		switch r {
		case TokenTypeNone:
			fmt.Fprint(f, "Authenticated by other means")
		case TokenTypePassword:
			fmt.Fprint(f, "argon2 hashed password")
		case TokenTypeWebAuthn:
			fmt.Fprint(f, "WebAuthn token")
		case TokenTypeOAuth:
			fmt.Fprint(f, "OAuth2 token")
		default:
			fmt.Fprint(f, "")
		}
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r TokenType) String() string {
	return string(r.v)
}
func (r TokenType) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *TokenType) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewTokenType(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r TokenType) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *TokenType) Scan(__iNpUt__ any) error {
	s, err := NewTokenType(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewTokenType(__iNpUt__ string) (TokenType, error) {
	switch __iNpUt__ {
	case string(tokenTypeNone):
		return TokenTypeNone, nil
	case string(tokenTypePassword):
		return TokenTypePassword, nil
	case string(tokenTypeWebAuthn):
		return TokenTypeWebAuthn, nil
	case string(tokenTypeOAuth):
		return TokenTypeOAuth, nil
	default:
		return TokenType{}, fmt.Errorf("invalid value for type 'TokenType': '%s'", __iNpUt__)
	}
}
