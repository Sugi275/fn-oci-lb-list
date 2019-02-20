package ocilib

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/common"
)

//GetEnvConfigProvider 環境変数を使用したConfigrationProviderを取得する
func GetEnvConfigProvider() common.ConfigurationProvider {
	return envConfigProvider{}
}

type envConfigProvider struct {
}

func (p envConfigProvider) PrivateRSAKey() (key *rsa.PrivateKey, err error) {
	envKey := "OCI_PrivateRSAKey"
	envKeyPassphrase := "OCI_PrivateRSAKey_passphrase"

	var privateKey string
	var privateKeyPassphrase string
	var ok bool

	if privateKey, ok = os.LookupEnv(envKey); !ok {
		err = fmt.Errorf("can not read PrivateKey from environment variable %s", envKey)
		return nil, err
	}

	if privateKeyPassphrase, ok = os.LookupEnv(envKeyPassphrase); !ok {
		err = fmt.Errorf("can not read PrivateKeyPassphrase from environment variable %s", envKey)
		return nil, err
	}

	key, err = common.PrivateKeyFromBytes([]byte(privateKey), &privateKeyPassphrase)
	return key, nil
}

func (p envConfigProvider) KeyID() (keyID string, err error) {
	ocid, err := p.TenancyOCID()
	if err != nil {
		return "", err
	}

	userocid, err := p.UserOCID()
	if err != nil {
		return "", err
	}

	fingerprint, err := p.KeyFingerprint()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", ocid, userocid, fingerprint), nil
}

func (p envConfigProvider) TenancyOCID() (value string, err error) {
	envKey := "OCI_TenancyOCID"
	var ok bool
	if value, ok = os.LookupEnv(envKey); !ok {
		err = fmt.Errorf("can not read Tenancy from environment variable %s", envKey)
		return "", err
	}

	return value, nil
}

func (p envConfigProvider) UserOCID() (value string, err error) {
	envKey := "OCI_UserOCID"
	var ok bool
	if value, ok = os.LookupEnv(envKey); !ok {
		err = fmt.Errorf("can not read user id from environment variable %s", envKey)
		return "", err
	}

	return value, nil
}

func (p envConfigProvider) KeyFingerprint() (value string, err error) {
	envKey := "OCI_KeyFingerprint"
	var ok bool
	if value, ok = os.LookupEnv(envKey); !ok {
		err = fmt.Errorf("can not read fingerprint from environment variable %s", envKey)
		return "", err
	}

	return value, nil
}

func (p envConfigProvider) Region() (value string, err error) {
	envKey := "OCI_Region"
	var ok bool
	if value, ok = os.LookupEnv(envKey); !ok {
		err = fmt.Errorf("can not read region from environment variable %s", envKey)
		return "", err
	}

	return value, nil
}
