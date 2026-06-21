package vrchat

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
)

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

	// Collect cookies from auth response
	allCookies := authResp.Cookies()
	c.client.SetCookies(allCookies)

	// Step 2: POST TOTP code to verify 2FA
	// Generate code here (after rate-limit wait) so it's within the 30s window
	if totpSecret != "" {
		code, err := totp.GenerateCode(totpSecret, time.Now())
		if err != nil {
			return fmt.Errorf("failed to generate TOTP code: %w", err)
		}

		fmt.Printf("Sending request to /auth/twofactorauth/totp/verify\n")
		fmt.Printf("TOTP Code: %s\n", code)

		verifyResp, err := c.client.R().
			SetBody(map[string]string{
				"code": code,
			}).
			Post("/auth/twofactorauth/totp/verify")
		if err != nil {
			return err
		}
		fmt.Printf("Response: %+v\n", verifyResp)
		if verifyResp.StatusCode() != 200 {
			return fmt.Errorf("failed to authenticate: %s", verifyResp.String())
		}
		// Merge cookies from TOTP verify response
		allCookies = append(allCookies, verifyResp.Cookies()...)
		c.client.SetCookies(allCookies)
	}

	for _, cookie := range allCookies {
		fmt.Printf("Cookie: %s=%s; Domain=%s; Path=%s; Expires=%s; Secure=%t; HttpOnly=%t\n",
			cookie.Name, cookie.Value, cookie.Domain, cookie.Path, cookie.Expires, cookie.Secure, cookie.HttpOnly)
	}

	return nil
}

func (c *Client) AuthenticateWithCookies(cookies []*http.Cookie, userAgent string) error {
	c.client.SetCookies(cookies)
	c.client.SetHeader("User-Agent", userAgent)
	return nil
}
