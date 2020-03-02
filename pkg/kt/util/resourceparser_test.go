package util

import (
	"testing"
)

func TestStringToObject(t *testing.T) {
	str := prepareSecretStr()
	obj, err := StringToObject(str)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(obj.GetObjectKind())
}

func prepareSecretStr() string {
	str := `apiVersion: v1
kind: Secret
metadata:
  labels:
    app: kt-connect
  name: kt-connect-shadow
type: Opaque`
	return str
}

func TestObjectToSecret(t *testing.T) {
	str := prepareSecretStr()
	obj, err := StringToObject(str)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	_, err = ObjectToSecret(obj)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
