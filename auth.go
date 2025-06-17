package vrchat

import (
	"fmt"
	"net/http"
)

// Authenticate authenticates the client with the VRChat API
func (c *Client) Authenticate(username, password, totp string, userAgent string) error {
	c.client.SetHeaders(map[string]string{
		"User-Agent":   userAgent,
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	resp, err := c.client.R().
		SetBasicAuth(username, password).
		SetBody(map[string]string{
			"code": totp,
		}).
		Post("/auth/twofactorauth/totp/verify")
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("failed to authenticate: %s", resp.String())
	}

	cookies := resp.Cookies()
	c.client.SetCookies(cookies)

	for _, cookie := range cookies {
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
