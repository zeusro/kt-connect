package util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// ParseK8sYaml parse text to runtime.Object
func ParseK8sYaml(contentBytes []byte) ([]runtime.Object, error) {
	fileAsString := string(contentBytes[:])
	sepYamlfiles := strings.Split(fileAsString, "---")
	retVal := make([]runtime.Object, 0, len(sepYamlfiles))
	for _, f := range sepYamlfiles {
		if f == "\n" || f == "" {
			// ignore empty cases
			continue
		}
		obj, err := StringToObject(f)
		if err != nil {
			return retVal, err
			// if error is  acceptable ,we can use continue
			// continue
		} else {
			retVal = append(retVal, obj)
		}
	}
	return retVal, nil
}

// StringToObject convert string to runtime.Object
func StringToObject(str string) (runtime.Object, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, groupVersionKind, err := decode([]byte(str), nil, nil)
	if err != nil {
		msg := fmt.Sprintf("Error while decoding YAML object. Err was: %s", err)
		return nil, errors.New(msg)
	}
	acceptedK8sTypes := regexp.MustCompile(`(Role|ClusterRole|RoleBinding|ClusterRoleBinding|ServiceAccount|Secret)`)
	if !acceptedK8sTypes.MatchString(groupVersionKind.Kind) {
		errMsg := fmt.Sprintf("unsupport resource. type: %s", groupVersionKind.Kind)
		return nil, errors.New(errMsg)
		// log.Printf("The custom-roles configMap contained K8s object types which are not supported! Skipping object with type: %s", groupVersionKind.Kind)
	}
	return obj, nil
}

// ObjectToSecret convert runtime.Object to secret
func ObjectToSecret(object runtime.Object) (*v1.Secret, error) {
	secret, ok := object.(*v1.Secret)
	var err error
	if !ok {
		errMsg := fmt.Sprintf("'object' is not a secret,it is %s", object.GetObjectKind())
		err = errors.New(errMsg)
	}
	return secret, err
}
