package vrchat

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

// cookiesHeader builds a Cookie header value from a slice of cookies,
// bypassing resty/http.CookieJar domain matching (which drops empty-domain cookies).
func cookiesHeader(cookies []*http.Cookie) string {
	parts := make([]string, 0, len(cookies))
	for _, ck := range cookies {
		parts = append(parts, ck.Name+"="+ck.Value)
	}
	return strings.Join(parts, "; ")
}

// Authenticate authenticates the client with the VRChat API
func (c *Client) Authenticate(username, password, totpSecret string, userAgent string) error {
	c.client.SetHeaders(map[string]string{
		"User-Agent":   userAgent,
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})

	// Step 1: GET /auth/user with Basic Auth to establish session
	fmt.Printf("Sending request to /auth/user\n")
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)

	authResp, err := c.client.R().
		SetBasicAuth(username, password).
		Get("/auth/user")
	if err != nil {
		return err
	}
	fmt.Printf("Response: %+v\n", authResp)
	if authResp.StatusCode() != 200 {
		return fmt.Errorf("failed to authenticate: %s", authResp.String())
	}

	// Collect cookies from auth response.
	// VRChat returns auth cookies with empty Domain; resty's cookie jar drops these
	// during domain matching, so we pass them as a raw Cookie header instead.
	allCookies := authResp.Cookies()
	fmt.Printf("Cookies from /auth/user: %d\n", len(allCookies))
	for _, ck := range allCookies {
		fmt.Printf("  Cookie: %s=%s Domain=%s\n", ck.Name, ck.Value, ck.Domain)
	}
	c.client.SetHeader("Cookie", cookiesHeader(allCookies))

	// Step 2: POST TOTP code to verify 2FA.
	// Reset rate limiter so the POST fires immediately — code expires in 30s,
	// but the rate limiter would otherwise wait up to 60s and expire it.
	if totpSecret != "" {
		c.ResetRateLimit()
		code, err := totp.GenerateCode(totpSecret, time.Now())
		if err != nil {
			return fmt.Errorf("failed to generate TOTP code: %w", err)
		}

		fmt.Printf("Sending request to /auth/twofactorauth/totp/verify\n")
		fmt.Printf("TOTP Code: %s\n", code)

		verifyResp, err := c.client.R().
			SetBody(map[string]string{"code": code}).
			Post("/auth/twofactorauth/totp/verify")
		if err != nil {
			return err
		}
		fmt.Printf("Response: %+v\n", verifyResp)
		if verifyResp.StatusCode() != 200 {
			return fmt.Errorf("failed to authenticate: %s", verifyResp.String())
		}
		allCookies = append(allCookies, verifyResp.Cookies()...)
		c.client.SetHeader("Cookie", cookiesHeader(allCookies))
	}

	for _, cookie := range allCookies {
		fmt.Printf("Cookie: %s=%s; Domain=%s; Path=%s; Expires=%s; Secure=%t; HttpOnly=%t\n",
			cookie.Name, cookie.Value, cookie.Domain, cookie.Path, cookie.Expires, cookie.Secure, cookie.HttpOnly)
	}

	return nil
}

func (c *Client) AuthenticateWithCookies(cookies []*http.Cookie, userAgent string) error {
	// Same empty-domain issue: set Cookie header directly instead of using SetCookies.
	c.client.SetHeader("Cookie", cookiesHeader(cookies))
	c.client.SetHeader("User-Agent", userAgent)
	return nil
}
