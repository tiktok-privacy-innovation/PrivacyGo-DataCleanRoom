package utils

import (
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func RunningInsideKubernetes() bool {
	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount"); os.IsNotExist(err) {
		hlog.Info("Not running inside Kubernetes")
		return false
	}
	hlog.Info("Running inside Kubernetes")
	return true
}

func GetNamespace() string {
	ns, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		hlog.Errorf("Failed to get namespace %s", err.Error())
	}
	return strings.TrimSpace(string(ns))
}
