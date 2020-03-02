package command

import(
	"github.com/alibaba/kt-connect/pkg/apiserver/util"
	"testing"
)

// func getTestClientset() *kubernetes.Clientset {
// 	//default
// 	userHome := util.HomeDir()
// 	configLocation := filepath.Join(userHome, ".kube", "config")
// 	clientSet, err := cluster.GetKubernetesClient(configLocation)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return clientSet
// }

// make sure that you have a real kubernetes cluster and you dare to create secret,or it will be fail
func TestCreateSecretFromFile(t *testing.T) {
	util.GetKubernetesClient()
}
