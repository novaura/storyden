package webauthn

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/app/resources/account"
)

// TODO: Move to actual accounts model.
type temporary struct{ handle string }

func (t *temporary) WebAuthnID() []byte                         { return []byte(t.handle) }
func (t *temporary) WebAuthnName() string                       { return t.handle }
func (t *temporary) WebAuthnDisplayName() string                { return t.handle }
func (t *temporary) WebAuthnIcon() string                       { return "" }
func (t *temporary) WebAuthnCredentials() []webauthn.Credential { return nil }

func (p *Provider) BeginRegistration(ctx context.Context, handle string) (*protocol.CredentialCreation, *webauthn.SessionData, error) {
	t := temporary{handle}

	// TODO: Check if handle already exists
	// if it exists, maybe we can short circuit the flow and switch to login?

	credentialOptions, sessionData, err := p.wa.BeginRegistration(&t,
		webauthn.WithAuthenticatorSelection(
			protocol.AuthenticatorSelection{
				// AuthenticatorAttachment: protocol.AuthenticatorAttachment(authType),
				// RequireResidentKey:      residentKeyRequirement,
				// UserVerification:        protocol.UserVerificationRequirement(userVer),
			}),
		// webauthn.WithConveyancePreference(protocol.ConveyancePreference(attType)),
	)
	if err != nil {
		return nil, nil, fault.Wrap(err, fmsg.With("failed to start registration"))
	}

	return credentialOptions, sessionData, nil
}

func (p *Provider) FinishRegistration(ctx context.Context, handle string, session webauthn.SessionData, parsedResponse *protocol.ParsedCredentialCreationData) (*webauthn.Credential, account.AccountID, error) {
	t := temporary{handle}

	credential, err := p.wa.CreateCredential(&t, session, parsedResponse)
	if err != nil {
		return nil, account.AccountID(xid.NilID()), fault.Wrap(err, fctx.With(ctx))
	}

	// TODO: Create a new account and then create an authentication item using
	// the credential obtained above.
	accountID := account.AccountID(xid.NilID())

	return credential, accountID, nil
}

func (p *Provider) BeginLogin(ctx context.Context, handle string) (*webauthn.SessionData, error) {
	t := temporary{handle}

	_, sd, err := p.wa.BeginLogin(&t)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return sd, nil
}

func (p *Provider) FinishLogin(ctx context.Context, handle string, session webauthn.SessionData, parsedResponse *protocol.ParsedCredentialAssertionData) (*webauthn.Credential, error) {
	t := temporary{handle}

	cred, err := p.wa.ValidateLogin(&t, session, parsedResponse)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx))
	}

	return cred, nil
}