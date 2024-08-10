/**
 * Generated by orval v6.31.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import type {
  AuthEmailBody,
  AuthEmailPasswordBody,
  AuthEmailVerifyBody,
  AuthPasswordBody,
  AuthPasswordCreateBody,
  AuthPasswordUpdateBody,
  AuthProviderListOKResponse,
  AuthSuccessOKResponse,
  OAuthProviderCallbackBody,
  PhoneRequestCodeBody,
  PhoneSubmitCodeBody,
  WebAuthnGetAssertionOKResponse,
  WebAuthnMakeAssertionBody,
  WebAuthnMakeCredentialBody,
  WebAuthnRequestCredentialOKResponse,
} from "../openapi-schema";
import { fetcher } from "../server";

/**
 * Retrieve a list of authentication providers. Storyden supports a few
ways to authenticate, from simple passwords to OAuth and WebAuthn. This
endpoint tells a client which auth capabilities are enabled.

 */
export type authProviderListResponse = {
  data: AuthProviderListOKResponse;
  status: number;
};

export const getAuthProviderListUrl = () => {
  return `/v1/auth`;
};

export const authProviderList = async (
  options?: RequestInit,
): Promise<authProviderListResponse> => {
  return fetcher<Promise<authProviderListResponse>>(getAuthProviderListUrl(), {
    ...options,
    method: "GET",
  });
};

/**
 * Register a new account with a username and password.
 */
export type authPasswordSignupResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthPasswordSignupUrl = () => {
  return `/v1/auth/password/signup`;
};

export const authPasswordSignup = async (
  authPasswordBody: AuthPasswordBody,
  options?: RequestInit,
): Promise<authPasswordSignupResponse> => {
  return fetcher<Promise<authPasswordSignupResponse>>(
    getAuthPasswordSignupUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(authPasswordBody),
    },
  );
};

/**
 * Sign in to an existing account with a username and password.
 */
export type authPasswordSigninResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthPasswordSigninUrl = () => {
  return `/v1/auth/password/signin`;
};

export const authPasswordSignin = async (
  authPasswordBody: AuthPasswordBody,
  options?: RequestInit,
): Promise<authPasswordSigninResponse> => {
  return fetcher<Promise<authPasswordSigninResponse>>(
    getAuthPasswordSigninUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(authPasswordBody),
    },
  );
};

/**
 * Given the requesting account does not have a password authentication,
add a password authentication method to it with the given password.

 */
export type authPasswordCreateResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthPasswordCreateUrl = () => {
  return `/v1/auth/password`;
};

export const authPasswordCreate = async (
  authPasswordCreateBody: AuthPasswordCreateBody,
  options?: RequestInit,
): Promise<authPasswordCreateResponse> => {
  return fetcher<Promise<authPasswordCreateResponse>>(
    getAuthPasswordCreateUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(authPasswordCreateBody),
    },
  );
};

/**
 * Given the requesting account has a password authentication, update the
password on file.

 */
export type authPasswordUpdateResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthPasswordUpdateUrl = () => {
  return `/v1/auth/password`;
};

export const authPasswordUpdate = async (
  authPasswordUpdateBody: AuthPasswordUpdateBody,
  options?: RequestInit,
): Promise<authPasswordUpdateResponse> => {
  return fetcher<Promise<authPasswordUpdateResponse>>(
    getAuthPasswordUpdateUrl(),
    {
      ...options,
      method: "PATCH",
      body: JSON.stringify(authPasswordUpdateBody),
    },
  );
};

/**
 * Register a new account with a email and password.
 */
export type authEmailPasswordSignupResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthEmailPasswordSignupUrl = () => {
  return `/v1/auth/email-password/signup`;
};

export const authEmailPasswordSignup = async (
  authEmailPasswordBody: AuthEmailPasswordBody,
  options?: RequestInit,
): Promise<authEmailPasswordSignupResponse> => {
  return fetcher<Promise<authEmailPasswordSignupResponse>>(
    getAuthEmailPasswordSignupUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(authEmailPasswordBody),
    },
  );
};

/**
 * Sign in to an existing account with a email and password.
 */
export type authEmailPasswordSigninResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthEmailPasswordSigninUrl = () => {
  return `/v1/auth/email-password/signin`;
};

export const authEmailPasswordSignin = async (
  authEmailPasswordBody: AuthEmailPasswordBody,
  options?: RequestInit,
): Promise<authEmailPasswordSigninResponse> => {
  return fetcher<Promise<authEmailPasswordSigninResponse>>(
    getAuthEmailPasswordSigninUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(authEmailPasswordBody),
    },
  );
};

/**
 * Register a new account with an email and optional password. The password
requirement is dependent on how the instance is configured for account
authentication with email addresses (password vs magic link.)

When the email address has not been registered, this endpoint will send
a verification email however it will also return a session cookie to
facilitate pre-verification usage of the platform. If the email address
already exists, no session cookie will be returned in order to prevent
arbitrary account control by a malicious actor. In this case, the email
will be sent again with the same OTP for the case where the user has
cleared their cookies or switched device but hasn't yet verified due to
missing the email or a delivery failure. In this sense, the endpoint can
act as a "resend verification email" operation as well as registration.

In the first case, a 200 response is provided with the session cookie,
in the second case, a 422 response is provided without a session cookie.

Given that this is an unauthenticated endpoint that triggers an email to
be sent to any public address, it MUST be heavily rate limited.

 */
export type authEmailSignupResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthEmailSignupUrl = () => {
  return `/v1/auth/email/signup`;
};

export const authEmailSignup = async (
  authEmailBody: AuthEmailBody,
  options?: RequestInit,
): Promise<authEmailSignupResponse> => {
  return fetcher<Promise<authEmailSignupResponse>>(getAuthEmailSignupUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(authEmailBody),
  });
};

/**
 * Sign in to an existing account with an email and optional password. The
behaviour of this endpoint depends on how the instance is configured. If
email+password is the preferred method, a cookie is returned on success
but if magic links are preferred, the endpoint will start the code flow.

 */
export type authEmailSigninResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthEmailSigninUrl = () => {
  return `/v1/auth/email/signin`;
};

export const authEmailSignin = async (
  authEmailBody: AuthEmailBody,
  options?: RequestInit,
): Promise<authEmailSigninResponse> => {
  return fetcher<Promise<authEmailSigninResponse>>(getAuthEmailSigninUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(authEmailBody),
  });
};

/**
 * Verify an email address using a token that was emailed to one of the
account's email addresses either set via sign up or added later.

 */
export type authEmailVerifyResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getAuthEmailVerifyUrl = () => {
  return `/v1/auth/email/verify`;
};

export const authEmailVerify = async (
  authEmailVerifyBody: AuthEmailVerifyBody,
  options?: RequestInit,
): Promise<authEmailVerifyResponse> => {
  return fetcher<Promise<authEmailVerifyResponse>>(getAuthEmailVerifyUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(authEmailVerifyBody),
  });
};

/**
 * OAuth2 callback.
 */
export type oAuthProviderCallbackResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getOAuthProviderCallbackUrl = (oauthProvider: string) => {
  return `/v1/auth/oauth/${oauthProvider}/callback`;
};

export const oAuthProviderCallback = async (
  oauthProvider: string,
  oAuthProviderCallbackBody: OAuthProviderCallbackBody,
  options?: RequestInit,
): Promise<oAuthProviderCallbackResponse> => {
  return fetcher<Promise<oAuthProviderCallbackResponse>>(
    getOAuthProviderCallbackUrl(oauthProvider),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(oAuthProviderCallbackBody),
    },
  );
};

/**
 * Start the WebAuthn registration process by requesting a credential.

 */
export type webAuthnRequestCredentialResponse = {
  data: WebAuthnRequestCredentialOKResponse;
  status: number;
};

export const getWebAuthnRequestCredentialUrl = (accountHandle: string) => {
  return `/v1/auth/webauthn/make/${accountHandle}`;
};

export const webAuthnRequestCredential = async (
  accountHandle: string,
  options?: RequestInit,
): Promise<webAuthnRequestCredentialResponse> => {
  return fetcher<Promise<webAuthnRequestCredentialResponse>>(
    getWebAuthnRequestCredentialUrl(accountHandle),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Complete WebAuthn registration by creating a new credential.
 */
export type webAuthnMakeCredentialResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getWebAuthnMakeCredentialUrl = () => {
  return `/v1/auth/webauthn/make`;
};

export const webAuthnMakeCredential = async (
  webAuthnMakeCredentialBody: WebAuthnMakeCredentialBody,
  options?: RequestInit,
): Promise<webAuthnMakeCredentialResponse> => {
  return fetcher<Promise<webAuthnMakeCredentialResponse>>(
    getWebAuthnMakeCredentialUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(webAuthnMakeCredentialBody),
    },
  );
};

/**
 * Start the WebAuthn assertion for an existing account.
 */
export type webAuthnGetAssertionResponse = {
  data: WebAuthnGetAssertionOKResponse;
  status: number;
};

export const getWebAuthnGetAssertionUrl = (accountHandle: string) => {
  return `/v1/auth/webauthn/assert/${accountHandle}`;
};

export const webAuthnGetAssertion = async (
  accountHandle: string,
  options?: RequestInit,
): Promise<webAuthnGetAssertionResponse> => {
  return fetcher<Promise<webAuthnGetAssertionResponse>>(
    getWebAuthnGetAssertionUrl(accountHandle),
    {
      ...options,
      method: "GET",
    },
  );
};

/**
 * Complete the credential assertion and sign in to an account.
 */
export type webAuthnMakeAssertionResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getWebAuthnMakeAssertionUrl = () => {
  return `/v1/auth/webauthn/assert`;
};

export const webAuthnMakeAssertion = async (
  webAuthnMakeAssertionBody: WebAuthnMakeAssertionBody,
  options?: RequestInit,
): Promise<webAuthnMakeAssertionResponse> => {
  return fetcher<Promise<webAuthnMakeAssertionResponse>>(
    getWebAuthnMakeAssertionUrl(),
    {
      ...options,
      method: "POST",
      body: JSON.stringify(webAuthnMakeAssertionBody),
    },
  );
};

/**
 * Start the authentication flow with a phone number. The handler will send
a one-time code to the provided phone number which must then be sent to
the other phone endpoint to verify the number and validate the account.

 */
export type phoneRequestCodeResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getPhoneRequestCodeUrl = () => {
  return `/v1/auth/phone`;
};

export const phoneRequestCode = async (
  phoneRequestCodeBody: PhoneRequestCodeBody,
  options?: RequestInit,
): Promise<phoneRequestCodeResponse> => {
  return fetcher<Promise<phoneRequestCodeResponse>>(getPhoneRequestCodeUrl(), {
    ...options,
    method: "POST",
    body: JSON.stringify(phoneRequestCodeBody),
  });
};

/**
 * Complete the phone number authentication flow by submitting the one-time
code that was sent to the user's phone.

 */
export type phoneSubmitCodeResponse = {
  data: AuthSuccessOKResponse;
  status: number;
};

export const getPhoneSubmitCodeUrl = (accountHandle: string) => {
  return `/v1/auth/phone/${accountHandle}`;
};

export const phoneSubmitCode = async (
  accountHandle: string,
  phoneSubmitCodeBody: PhoneSubmitCodeBody,
  options?: RequestInit,
): Promise<phoneSubmitCodeResponse> => {
  return fetcher<Promise<phoneSubmitCodeResponse>>(
    getPhoneSubmitCodeUrl(accountHandle),
    {
      ...options,
      method: "PUT",
      body: JSON.stringify(phoneSubmitCodeBody),
    },
  );
};

/**
 * Remove cookies from requesting client.
 */
export type authProviderLogoutResponse = {
  data: void;
  status: number;
};

export const getAuthProviderLogoutUrl = () => {
  return `/v1/auth/logout`;
};

export const authProviderLogout = async (
  options?: RequestInit,
): Promise<authProviderLogoutResponse> => {
  return fetcher<Promise<authProviderLogoutResponse>>(
    getAuthProviderLogoutUrl(),
    {
      ...options,
      method: "GET",
    },
  );
};