package jwtutil

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	errormodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/error-model"
	jwtmodel "github.com/synzofficial/nsd-synz-sharelib-api/pkg/model/jwtmodel"
	"github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/jws"
	typeconvertutil "github.com/synzofficial/nsd-synz-sharelib-api/pkg/util/type-convert-util"
)

var (
	errIncorrectJWT = errors.New("JWT token is incorrect")
)

func ExtractJWT(authorization string) (*jwtmodel.CustomClaims, error) {
	now := time.Now()

	claimSet, err := jws.Decode(authorization)
	if err != nil {
		// TODO: logging
		return nil, err
	}

	iat := time.Unix(claimSet.Iat, 0)
	if now.Sub(iat) >= 15*time.Minute {
		return nil, errormodel.UnauthorizedError(err)
	}

	customClaim, ok := claimSet.PrivateClaims[jwtmodel.CUSTOM_CLAIMS_KEY].(map[string]interface{})
	if !ok {
		return nil, errors.New("unable to get custom claims from jwt")
	}
	jsonByte, err := json.Marshal(customClaim)
	if err != nil {
		// log.Errorf("error while marshal json byte: %+v", err)
		return nil, err
	}

	var customClaims jwtmodel.CustomClaims
	err = json.Unmarshal(jsonByte, &customClaims)
	if err != nil {
		// log.Errorf("error while unmashal custom claims: %+v", err)
		return nil, err
	}

	if !customClaims.UserType.IsValidate() {
		return nil, errors.New("JWT token user type is incorrect")
	}

	parts := strings.Split(authorization, ".")
	if len(parts) != 3 {
		// log.Errorf("JWT token is incorrect")
		return nil, errIncorrectJWT
	}

	// Decode jwt payload (second part)
	payloadJSON, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		// log.Errorf("error while decode jwt payload")
		return nil, errIncorrectJWT
	}

	var payloadData map[string]interface{}
	if err := json.Unmarshal(payloadJSON, &payloadData); err != nil {
		// log.Errorf("error parsing payload JSON: %v", err)
		return nil, errIncorrectJWT
	}

	// TODO: map some payload data to customClaims
	/*
		mobileNo, ok := payloadData["mobile_no"].(string)
		if ok {
			customClains.Phone = mobileNo
		}
	*/

	return &customClaims, nil
}

func GetJwt(ctx context.Context) (*jwtmodel.CustomClaims, bool) {
	ctx = typeconvertutil.ConvertGinContext(ctx)
	jwtModel, ok := ctx.Value(jwtmodel.JwtContextKey{}).(*jwtmodel.CustomClaims)
	return jwtModel, ok
}
