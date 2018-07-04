package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	//"github.com/a-fis/urhodling.com/models"
	"github.com/gin-gonic/gin"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

const (
	flashKeyInfo    = "flash_key_info"
	flashKeyError   = "flash_key_error"
	flashKeyWarning = "flash_key_warning"
	flashKeySuccess = "flash_key_success"
)

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func init() {}

// RenderTemplate ...
func RenderTemplate(c echo.Context, tmpl string, datatmp interface{}) error {
	data := map[string]interface{}{}

	list, ok := datatmp.(map[string]interface{})

	if ok {
		for n, v := range list {
			data[n] = v
		}
	}

	data["flash_error"] = GetFlashError(c)
	data["flash_warning"] = GetFlashWarning(c)
	data["flash_info"] = GetFlashInfo(c)
	data["flash_success"] = GetFlashSuccess(c)
	data["my_email"] = GetMyEmail(c)
	data["account"] = GetMyAccount(c)
	data["tradingSymbol"] = GetMarketSymble(c)
	data["marketCap"] = GetMarketCap(c)
	data["email_verified"] = GetEmailVerified(c)
	data["is_login"] = IsLogin(c)
	data["current_uri"] = c.Request().URL.Path
	if _, ok = data["page_title"]; !ok {
		data["page_title"] = data["site_title"]
	}

	csrf := c.Get("csrf")
	csrfToken, ok := csrf.(string)
	if ok && csrfToken != "" {
		data["csrf_token"] = csrfToken
	}

	localeStr, ok := c.Get("locale").(string)
	if ok && localeStr != "" {
		data["locale"] = localeStr
	} else {
		data["locale"] = "en"
	}

	displayCurrencyStr, ok := c.Get("display_currency").(string)
	if ok && localeStr != "" {
		data["displayCurrency"] = displayCurrencyStr
	} else {
		data["displayCurrency"] = "USD"
	}

	return c.Render(200, tmpl, data)
}

func getContextInt64WithDefault(c echo.Context, key string, val int64) int64 {
	val1 := c.Get(key)
	val2, ok := val1.(int64)
	if ok {
		return val2
	}

	return val
}

func getContextStringWithDefault(c echo.Context, key, dval string) string {
	val := c.Get(key)
	val2, ok := val.(string)
	if ok && val2 != "" {
		return val2
	}

	return dval
}

func setSessionString(c echo.Context, key, value string) error {
	session := session.Default(c)
	if value == "" {
		return nil
	}

	session.Set(key, value)
	session.Save()

	return nil
}

// Redirect ...
func Redirect(c echo.Context, url string) error {
	return c.Redirect(302, url)
}

// getSessionString ...
func getSessionString(c echo.Context, key string) string {
	session := session.Default(c)
	obj := session.Get(key)
	if val, ok := obj.(string); ok {
		return string(val)
	}

	return ""
}

// SetFlashError ...
func SetFlashError(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyError)
}

func setFlash(c echo.Context, msg, key string) error {
	if msg == "" {
		return nil
	}

	session := session.Default(c)
	session.Set(key, msg)
	err := session.Save()
	if err != nil {
		log.Print(err)
	}

	return err
}

func getFlash(c echo.Context, key string) string {

	sess := session.Default(c)
	obj := sess.Get(key)

	if msg, ok := obj.(string); ok {
		sess.Delete(key)
		sess.Save()

		return msg
	}

	return ""
}

// SetFlashSuccess ...
func SetFlashSuccess(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeySuccess)
}

// GetMyUserID ...
func GetMyUserID(c echo.Context) int64 {
	session := session.Default(c)
	if userID := session.Get("user_id"); userID != nil {
		val, ok := userID.(int64)
		if ok {
			return val
		}
	}

	return 0
}

// GetEmailVerified ...
func GetEmailVerified(c echo.Context) bool {
	session := session.Default(c)
	if emailVerified := session.Get("email_verified"); emailVerified != nil {
		val, ok := emailVerified.(bool)
		if ok {
			return val
		}
	}

	return false
}

// GetMyUserIDStr ...
func GetMyUserIDStr(c echo.Context) string {
	userID := GetMyUserID(c)
	if userID > 0 {
		return fmt.Sprintf("%d", userID)
	}

	return ""
}

// IsLogin ...
func IsLogin(c echo.Context) bool {
	isLogin := c.Get("is_login")
	isLoginBool, ok := isLogin.(bool)
	if ok && isLoginBool {
		return true
	}

	session := session.Default(c)
	flag := session.Get("is_login")
	if flag != nil {
		val, ok := flag.(int)
		if ok && val == 1 {
			return true
		}
	}

	return false
}

/*
// SetAuth ...
func SetAuth(c echo.Context, user models.User) {
	session := session.Default(c)
	session.Set("user_id", user.ID)
	session.Set("email", user.Email)
	session.Set("email_verified", user.Verified)
	session.Set("is_login", 1)
	session.Save()
}
*/

// SetFlashInfo ...
func SetFlashInfo(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyInfo)
}

// GetMarketSymble ...
func GetMarketSymble(c echo.Context) string {
	session := session.Default(c)
	if fname := session.Get("tradingSymbol"); fname != nil {
		val, ok := fname.(string)
		if ok {
			return val
		}
	}

	return ""
}

// GetMarketCap ...
func GetMarketCap(c echo.Context) string {
	session := session.Default(c)
	if fname := session.Get("marketCap"); fname != nil {
		val, ok := fname.(string)
		if ok {
			return val
		}
	}

	return ""
}

// GetMyAccount ...
func GetMyAccount(c echo.Context) string {
	session := session.Default(c)
	if fname := session.Get("account"); fname != nil {
		val, ok := fname.(string)
		if ok {
			return val
		}
	}

	return ""
}

// GetMyEmail ...
func GetMyEmail(c echo.Context) string {
	session := session.Default(c)
	if fname := session.Get("email"); fname != nil {
		val, ok := fname.(string)
		if ok {
			return val
		}
	}

	return ""
}

// GetFlashError ...
func GetFlashError(c echo.Context) string {
	return getFlash(c, flashKeyError)
}

// GetFlashSuccess ...
func GetFlashSuccess(c echo.Context) string {
	return getFlash(c, flashKeySuccess)
}

// GetFlashInfo ...
func GetFlashInfo(c echo.Context) string {
	return getFlash(c, flashKeyInfo)
}

// SetFlashWarning ...
func SetFlashWarning(c echo.Context, msg string) error {
	return setFlash(c, msg, flashKeyWarning)
}

// GetFlashWarning ...
func GetFlashWarning(c echo.Context) string {
	return getFlash(c, flashKeyWarning)
}

// ClearAuth ...
func ClearAuth(c echo.Context) {
	session := session.Default(c)
	session.Delete("is_login")
	session.Delete("email_verified")
	session.Delete("user_id")
	session.Delete("email")

	session.Save()
}

// getenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}

	return v
}

// OutputErrorJSON ...
func OutputErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "error",
		"message": msg,
	})
}

// OutputOKJSON ...
func OutputOKJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
	})
}

// OutputOKDataJSON ...
func OutputOKDataJSON(c *gin.Context, msg string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": msg,
		"data":    data,
	})
}
