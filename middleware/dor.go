package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

func DorAccessUser(s sessions.Session) (bool, TokenClaim, error) {

	tokenClaim := TokenClaim{}

	if s == nil {
		return false, tokenClaim, errors.New("session null")
	}

	claims := s.Get("claims")

	if claims != nil {

		mapstructure.Decode(claims, &tokenClaim)

		fmt.Println("check = ", tokenClaim)

		if tokenClaim.Client_Role == "user" || tokenClaim.Client_Role == "service" {
			return true, tokenClaim, nil
		} else {
			return false, tokenClaim, errors.New("user_biz tidak dapat mengakses")
		}

	} else {
		return false, tokenClaim, errors.New("session tidak ditemukan")
	}

}

func DorAccessUserBiz(s sessions.Session) (bool, TokenClaim, error) {

	tokenClaim := TokenClaim{}

	if s == nil {
		return false, tokenClaim, errors.New("session null")
	}

	claims := s.Get("claims")

	if claims != nil {

		mapstructure.Decode(claims, &tokenClaim)

		if tokenClaim.Client_Role == "user_biz" || tokenClaim.Client_Role == "service" {
			return true, tokenClaim, nil
		} else {
			return false, tokenClaim, errors.New("user tidak dapat mengakses")
		}

	} else {
		return false, tokenClaim, errors.New("session tidak ditemukan")
	}

}
func DorAccessService(s sessions.Session) (bool, TokenClaim, error) {

	tokenClaim := TokenClaim{}

	if s == nil {
		return false, tokenClaim, errors.New("session null")
	}

	claims := s.Get("claims")

	if claims != nil {

		mapstructure.Decode(claims, &tokenClaim)

		if tokenClaim.Client_Role == "service" {
			return true, tokenClaim, nil
		} else {
			return false, tokenClaim, errors.New("user & user_biz tidak dapat mengakses")
		}

	} else {
		return false, tokenClaim, errors.New("session tidak ditemukan")
	}

}
func DorAccessAll(s sessions.Session) (bool, TokenClaim, error) {

	tokenClaim := TokenClaim{}

	if s == nil {
		return false, tokenClaim, errors.New("session null")
	}

	claims := s.Get("claims")

	if claims != nil {

		mapstructure.Decode(claims, &tokenClaim)

		if tokenClaim.Client_Role == "service" || tokenClaim.Client_Role == "user" || tokenClaim.Client_Role == "user_biz" {
			return true, tokenClaim, nil
		} else {
			return false, tokenClaim, errors.New("unknow tidak dapat mengakses")
		}

	} else {
		return false, tokenClaim, errors.New("session tidak ditemukan")
	}

}
