package auth

import (
	"errors"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type ITokenManager interface {
	CreateToken(userID string, duration time.Duration) (string, error)
	GetUserID(token string) (string, error)
}

type TokenManager struct {
	signAlg    jwa.SignatureAlgorithm
	encryptAlg jwa.KeyEncryptionAlgorithm
	privateKey jwk.RSAPrivateKey
	issuer     string
}

func NewTokenManager(issuer string, keyPath string) (*TokenManager, error) {
	if issuer == "" || keyPath == "" {
		return nil, errors.New("[ERROR] invalid parameter")
	}
	src, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, errors.New("[ERROR] failed to read private-key-file")
	}
	key, err := jwk.ParseKey(src, jwk.WithPEM(true))
	if err != nil {
		return nil, errors.New("[ERROR] failed to parse pem file")
	}
	var jwkPrivateKey jwk.RSAPrivateKey
	var ok bool
	if jwkPrivateKey, ok = key.(jwk.RSAPrivateKey); !ok {
		return nil, errors.New("[ERROR] failed to parse jwk-private-key from rsa-private-key")
	}
	return &TokenManager{
		signAlg:    jwa.RS256(),
		encryptAlg: jwa.RSA_OAEP(),
		issuer:     issuer,
		privateKey: jwkPrivateKey,
	}, nil
}

func (m *TokenManager) CreateToken(userID string, duration time.Duration) (string, error) {
	if userID == "" {
		return "", errors.New("[ERROR] invalid token parameter")
	}
	publicKey, err := m.privateKey.PublicKey()
	if err != nil {
		return "", err
	}
	now := time.Now().UTC()
	token, err := jwt.NewBuilder().Issuer(m.issuer).IssuedAt(now).Subject(userID).Expiration(now.Add(duration)).Build()
	if err != nil {
		return "", err
	}
	signed, err := jwt.Sign(token, jwt.WithKey(m.signAlg, m.privateKey))
	if err != nil {
		return "", err
	}
	encrypted, err := jwe.Encrypt(signed, jwe.WithKey(m.encryptAlg, publicKey))
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func (m *TokenManager) GetUserID(token string) (string, error) {
	publicKey, err := m.privateKey.PublicKey()
	if err != nil {
		return "", err
	}
	decrypted, err := jwe.Decrypt([]byte(token), jwe.WithKey(m.encryptAlg, m.privateKey))
	if err != nil {
		return "", err
	}
	verifyed, err := jwt.Parse(decrypted, jwt.WithKey(m.signAlg, publicKey))
	if err != nil {
		return "", errors.New("[ERROR] failed to verify token")
	}
	subject, _ := verifyed.Subject()
	return subject, nil
}
