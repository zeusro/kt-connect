package ssh

import (
	"fmt"
	"testing"

	"github.com/alibaba/kt-connect/pkg/kt/util"
	"github.com/stretchr/testify/assert"
)

func TestGeneraKeyPair(t *testing.T) {
	privateKeyPem, publicKeyPem, err := GeneraKeyPair(2014)
	t.Logf("privateKeyPem:\n%s\npublicKeyPem:%s\n", privateKeyPem, publicKeyPem)
	assert.True(t, len(privateKeyPem) > 0)
	assert.True(t, len(publicKeyPem) > 0)
	assert.Nil(t, err)
}

func BenchmarkGeneraKeyPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneraKeyPair(2014)
	}
}

func TestGetPrivateKeyFromFile(t *testing.T) {
	userHome := util.HomeDir()
	filePath := fmt.Sprintf("%s/.ssh/id_rsa", userHome)
	privateKeyPem, err := GetPrivateKeyFromFile(filePath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Nil(t, err)
	// >
	assert.Greater(t, len(privateKeyPem), 0)
}
func TestGetPrivateKeyFromFile_NotExistFile(t *testing.T) {
	filePath := "~/.ssh/id_rsa"
	_, err := GetPrivateKeyFromFile(filePath)
	assert.NotNil(t, err)
}

func TestGetPublicKeyFromFile(t *testing.T) {
	userHome := util.HomeDir()
	filePath := fmt.Sprintf("%s/.ssh/id_rsa.pub", userHome)
	publicKeyPem, err := GetPublicKeyFromFile(filePath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.Nil(t, err)
	// >
	assert.Greater(t, len(publicKeyPem), 0)
}

func TestGetPublicKeyFromFile_NotExistFile(t *testing.T) {
	filePath := "~/.ssh/id_rsa.pub"
	_, err := GetPublicKeyFromFile(filePath)
	assert.NotNil(t, err)
}
