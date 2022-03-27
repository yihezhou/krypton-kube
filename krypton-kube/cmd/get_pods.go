package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var podCmd = &cobra.Command{
	Use:	"get_pod",
	Short:	"get_pod subcommand to get kubernetes pods.",
	Long:	"get_pod subcommand to get kubernetes pods. return the pods information.",
	Run: 	func (cmd *cobra.Command, args []string) {
		//fmt.Println("here to use go-client library to get pods information.")
		clientset, err := GetK8sClient()
		if err != nil {
			log.Fatal(err)
		} else {
			//fmt.Println("clientset init OK", clientset)
			pods, err := clientset.CoreV1().Pods(nameSpace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Fatal(err)
			}
			var result []map[string]interface{}
			for _, node := range pods.Items {
				res := make(map[string]interface{})
				res["namespace"] = node.Namespace
				res["name"] = node.Name
				res["status"] = node.Status.Phase
				res["restarts"] = node.Status.ContainerStatuses[0].RestartCount
				res["node"] = node.Status.HostIP
				res["create_at"] = node.CreationTimestamp
				result = append(result, res)
			}
			if jsonStdout, err := json.Marshal(result); err != nil {
				panic(err)
			} else {
				fmt.Print(string(jsonStdout))
			}
		}
	},
}

func init() {
	podCmd.Flags().StringVar(&authType, "auth_type", "kubeconfig", "auth_type: kubeconfig, password, token")
	podCmd.Flags().StringVar(&bearerToken, "bearer_token", "", "bearer token for k8s api server")
	podCmd.Flags().StringVar(&clusterUrl, "cluster_url", "", "k8s api server url")
	podCmd.Flags().StringVar(&userName, "username", "", "username for k8s api server")
	podCmd.Flags().StringVar(&passWord, "password", "", "password for k8s api server")
	podCmd.Flags().StringVar(&podName, "pod_name", "", "pod name for query")
	podCmd.Flags().StringVar(&nameSpace, "namespace", "", "k8s namespace")
	podCmd.Flags().StringVar(&kubeConfig, "kube_config", "", "kube config save path")

	rootCmd.AddCommand(podCmd)
}