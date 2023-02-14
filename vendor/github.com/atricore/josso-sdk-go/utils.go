package cli

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

// Removes unsupported chars from name
func sanitizeName(name string) string {
	// Replace unsupported chars

	chars := []string{"]", "^", "\\\\", "[", "(", ")", "-", "/", "_"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	name = re.ReplaceAllString(name, "")

	return strings.ToLower(name)
}

func PtrString(s interface{}) *string {
	if _, ok := s.(string); ok {
		return api.PtrString(s.(string))
	}
	return nil
}

func PtrBool(s interface{}) *bool {
	if _, ok := s.(bool); ok {
		return api.PtrBool(s.(bool))
	}
	return nil
}

func StrDeref(p *string) string {
	result := "NIL"
	if p != nil {
		result = *p
	}
	return result
}

func BoolDeref(p *bool) bool {
	result := false
	if p != nil {
		result = *p
	}
	return result

}

func Int64Deref(p *int64) int64 {
	var result int64
	if p != nil {
		result = *p
	}
	return result

}

func Int32Deref(p *int32) int32 {
	var result int32
	if p != nil {
		result = *p
	}
	return result

}

func LocationToStr(l *api.LocationDTO) string {

	if l == nil {
		return ""
	}

	url := strings.ToLower(*l.Protocol) + "://" + *l.Host

	if *l.Port != 0 && *l.Port != 80 && *l.Port != 443 {
		url = fmt.Sprintf("%s:%d", url, *l.Port)
	}

	if *l.Context != "" {
		url += "/" + *l.Context
	}

	if *l.Uri != "" {
		url += "/" + *l.Uri
	}

	return url
}

func StrToLocation(v string) (*api.LocationDTO, error) {
	// Parse URL
	u, err := url.Parse(v)
	if err != nil {
		return nil, err
	}

	location := api.NewLocationDTO()
	location.Protocol = &u.Scheme

	// Strip port from host
	h := u.Hostname()
	location.Host = &h

	// Get Port
	location.Port, err = StrToPort(u.Port())

	// Default ports
	if *location.Port == 0 {
		switch u.Scheme {
		case "https":
			var p int32 = 443
			location.Port = &p
		default:
			var p int32 = 80
			location.Port = &p
		}

	}

	s := strings.SplitN(u.Path, "/", 3)

	if len(s) > 1 {
		location.Context = &s[1]
	} else {
		v = ""
		location.Context = &v
	}

	if len(s) > 2 {
		location.Uri = &s[2]
	} else {
		v = ""
		location.Uri = &v
	}

	return location, err
}

func StrToPort(v string) (*int32, error) {

	var port int32 = 0
	if v == "" {
		return &port, nil
	}
	i, err := strconv.Atoi(v)
	y := int32(i)
	return &y, err
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func buildErrorMsg(err string, valErrors []string) string {
	var msg string
	if len(valErrors) > 0 {
		msg = fmt.Sprintf("%s : %#v", err, valErrors)
	} else {
		msg = err
	}
	return msg
}
