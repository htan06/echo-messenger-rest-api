package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"log"
	"os"
	"strconv"
	"time"
)

type JWTConfig struct {
	privateKeyAccess  *rsa.PrivateKey
	publicKeyAccess   *rsa.PublicKey
	privateKeyRefresh *rsa.PrivateKey
	publicKeyRefresh  *rsa.PublicKey
	ttlAccess         time.Duration
	ttlRefresh        time.Duration
}

func GetJWTConfig() *JWTConfig {
	// PRIVATE ACCESS KEY
	privateKeyAccessByte, err := base64.StdEncoding.DecodeString(os.Getenv("TOKEN_ACCESS_PRIVATE_KEY"))
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}
	rsaAccessPrivateKey, err := x509.ParsePKCS8PrivateKey(privateKeyAccessByte)
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}

	// PUBLIC ACCESS KEY
	publicKeyAccessByte, err := base64.StdEncoding.DecodeString(os.Getenv("TOKEN_ACCESS_PUBLIC_KEY"))
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}
	rsaAccessPublicKey, err := x509.ParsePKIXPublicKey(publicKeyAccessByte)
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}
	
	// PRIVATE REFRESH KEY
	privateKeyRefreshByte, err := base64.StdEncoding.DecodeString(os.Getenv("TOKEN_REFRESH_PRIVATE_KEY"))
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}
	rsaRefreshPrivateKey, err := x509.ParsePKCS8PrivateKey(privateKeyRefreshByte)
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}

	// PUBLIC REFRESH KEY
	publicKeyRefreshByte, err := base64.StdEncoding.DecodeString(os.Getenv("TOKEN_REFRESH_PUBLIC_KEY"))
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}
	rsaRefreshPublicKey, err := x509.ParsePKIXPublicKey(publicKeyRefreshByte)
	if err != nil {
		log.Fatal("JWT config: " + err.Error())
	}

	//TTL
	ttlAccessToken, err := strconv.Atoi(os.Getenv("TOKEN_ACCESS_TTL"))
	if err != nil {
		log.Fatal("JWT config: cannot parse ttl access token: " + err.Error())
	}
	ttlRefreshtoken, err := strconv.Atoi(os.Getenv("TOKEN_REFRESH_TTL"))
	if err != nil {
		log.Fatal("JWT config: cannot parse ttl refresh token: " + err.Error())
	}

	return &JWTConfig{
		privateKeyAccess:  rsaAccessPrivateKey.(*rsa.PrivateKey),
		publicKeyAccess:   rsaAccessPublicKey.(*rsa.PublicKey),
		privateKeyRefresh: rsaRefreshPrivateKey.(*rsa.PrivateKey),
		publicKeyRefresh:  rsaRefreshPublicKey.(*rsa.PublicKey),
		ttlAccess:         time.Second * time.Duration(ttlAccessToken),
		ttlRefresh:        time.Second * time.Duration(ttlRefreshtoken),
	}
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

func (j *JWTConfig) TtlAccess() time.Duration {
	return j.ttlAccess
}

func (j *JWTConfig) TtlRefresh() time.Duration {
	return j.ttlRefresh
}