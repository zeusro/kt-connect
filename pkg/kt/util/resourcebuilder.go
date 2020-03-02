package util

import (
	"encoding/base64"

	v1 "k8s.io/api/core/v1"
)

// SecretBuilder Help you create k8s secret
type SecretBuilder struct {
	data            map[string][]byte
	secretPrototype *v1.Secret
}

// NewSecretBuilder NEW SecretBuilder
func NewSecretBuilder() *SecretBuilder {
	builder := SecretBuilder{
		secretPrototype: &v1.Secret{},
	}
	return &builder
}

func (builder *SecretBuilder) LoadYAMLTemplate(content string) {
	//TODO
}

// AddData add key value data
func (builder *SecretBuilder) AddData(datas map[string]string) {
	builder.data = make(map[string][]byte, 0)
	for k, v := range datas {
		var data []byte
		base64.StdEncoding.Encode(data, []byte(v))
		builder.data[k] = data
	}
}

// Build Build Secret
func (builder *SecretBuilder) Build() (secret *v1.Secret, err error) {
	return builder.secretPrototype, err
}
