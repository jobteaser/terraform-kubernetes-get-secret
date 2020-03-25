package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
)

type Query struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Context   string `json:"context"`
}

type Result struct {
	Value string `json:"value"`
}

func main() {
	// Get input type json see Query struct for json format
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err.Error())
	}

	// Loading json input using the Query struct
	var q Query
	err = json.Unmarshal(input, &q)
	if err != nil {
		panic(err.Error())
	}

	//  Get the local kube config with context override.
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{CurrentContext: q.Context}).ClientConfig()
	if err != nil {
		panic(err.Error())
	}

	// Creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Get the secret
	secret, err := clientset.CoreV1().Secrets(q.Namespace).Get(q.Name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Check if key exist
	val, ok := secret.Data[q.Key]
	if !ok {
		panic("Key not found")
	}

	// Output result as json then write to Stdout
	r := &Result{
		Value: string(val),
	}

	o, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		panic(err.Error())
	}

	os.Stdout.Write(o)

	return
}
