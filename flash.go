package flash

import (
	"encoding/base64"
	"net/http"
	"strings"
)

type (
	// DataValue flash data vakye
	DataValue map[string][]string
	// Data flash data
	Data struct {
		Name    string // flash name
		Data    DataValue
		Request *http.Request
	}
)

const (
	itemSeparator      = "\x25"
	keyValueSeparator  = "\x23"
	valueItemSeparator = "\x24"
)

// NewFlash -
func NewFlash(name string, r *http.Request) *Data {
	return &Data{
		Name:    name,
		Data:    make(DataValue),
		Request: r,
	}
}

// Set -
func (f *Data) Set(val DataValue) *Data {
	f.Data = val
	return f
}

// Save -
func (f *Data) Save(w http.ResponseWriter) {
	var val strings.Builder
	for k, v := range f.Data {
		val.WriteString(itemSeparator + k + keyValueSeparator + strings.Join(v, valueItemSeparator) + itemSeparator)
	}

	f.saveToCookie(w, val.String())
}

func (f *Data) Read(w http.ResponseWriter) DataValue {
	cookie, err := f.Request.Cookie(f.Name)
	if err != nil {
		return nil
	}

	v, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil
	}

	result := make(DataValue)
	items := strings.Split(string(v), itemSeparator)
	for _, item := range items {
		if len(item) > 0 {
			kv := strings.Split(item, keyValueSeparator)
			if (len(kv) == 2) && (kv[1] != "") {
				result[kv[0]] = strings.Split(kv[1], valueItemSeparator)
			}
		}
	}

	f.removeCookie(w)
	return result
}

func (f *Data) setCookie(w http.ResponseWriter, val string, maxAge int) {
	http.SetCookie(w, &http.Cookie{
		Name:     f.Name,
		Value:    val,
		MaxAge:   maxAge,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	})
}

func (f *Data) saveToCookie(w http.ResponseWriter, val string) {
	f.setCookie(w, base64.URLEncoding.EncodeToString([]byte(val)), 0)
}

func (f *Data) removeCookie(w http.ResponseWriter) {
	f.setCookie(w, "", -1)
}
