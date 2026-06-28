package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"strconv"
	"time"
)

type JWTConfig struct {
	privateKeyAccess *rsa.PrivateKey
	publicKeyAccess  *rsa.PublicKey
	ttlAccess        time.Duration

	privateKeyRefresh *rsa.PrivateKey
	publicKeyRefresh  *rsa.PublicKey
	ttlRefresh        time.Duration

	privateKeyRegister []byte
	ttlRegister        time.Duration
}

func GetJWTConfig(publicData []byte, privateData []byte) *JWTConfig {

	var jwtConfig JWTConfig
	//TTL
	ttlAccessToken, err := strconv.Atoi(os.Getenv("TOKEN_ACCESS_TTL"))
	if err != nil {
		log.Fatal("JWT config: cannot parse ttl access token: " + err.Error())
	}
	jwtConfig.ttlAccess = time.Second * time.Duration(ttlAccessToken)

	ttlRefreshtoken, err := strconv.Atoi(os.Getenv("TOKEN_REFRESH_TTL"))
	if err != nil {
		log.Fatal("JWT config: cannot parse ttl refresh token: " + err.Error())
	}
	jwtConfig.ttlRefresh = time.Second * time.Duration(ttlRefreshtoken)
	
	ttlRegistertoken, err := strconv.Atoi(os.Getenv("TOKEN_REGISTER_TTL"))
	if err != nil {
		log.Fatal("JWT config: cannot parse ttl register token: " + err.Error())
	}
	jwtConfig.ttlRegister = time.Second * time.Duration(ttlRegistertoken)

	for {
		block, rest := pem.Decode(privateData)
		if block == nil {
			break
		}
		switch block.Type {
		case "REFRESH PRIVATE KEY":
			key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				log.Fatal(err.Error())
			}
			jwtConfig.privateKeyRefresh = key.(*rsa.PrivateKey)
		case "ACCESS PRIVATE KEY":
			key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				log.Fatal(err.Error())
			}
			jwtConfig.privateKeyAccess = key.(*rsa.PrivateKey)
		case "REGISTER PRIVATE KEY":
			jwtConfig.privateKeyRegister = block.Bytes
		}
		privateData = rest
	}

	for {
		block, rest := pem.Decode(publicData)
		if block == nil {
			break
		}
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			log.Fatal(err.Error())
		}
		switch block.Type {
		case "REFRESH PUBLIC KEY":
			jwtConfig.publicKeyRefresh = key.(*rsa.PublicKey)
		case "ACCESS PUBLIC KEY":
			jwtConfig.publicKeyAccess = key.(*rsa.PublicKey)
		}
		publicData = rest
	}

	return &jwtConfig
}

func (j *JWTConfig) PrivateKeyAccess() *rsa.PrivateKey {
	return j.privateKeyAccess
}

func (j *JWTConfig) PublicKeyAccess() *rsa.PublicKey {
	return j.publicKeyAccess
}

func (j *JWTConfig) PrivateKeyRefresh() *rsa.PrivateKey {
	return j.privateKeyRefresh
}

func (j *JWTConfig) PublicKeyRefresh() *rsa.PublicKey {
	return j.publicKeyRefresh
}

func (j *JWTConfig) PrivateKeyRegister() []byte {
	return j.privateKeyRegister
}

func (j *JWTConfig) TtlAccess() time.Duration {
	return j.ttlAccess
}

func (j *JWTConfig) TtlRefresh() time.Duration {
	return j.ttlRefresh
}

func (j *JWTConfig) TtlRegister() time.Duration {
	return j.ttlRegister
}