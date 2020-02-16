module github.com/davidwalter0/tools/x/k8s/nodes

go 1.13

require (
	github.com/davidwalter0/go-cfg v1.2.3
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	github.com/davidwalter0/tools/x/k8s/kubeconfig v0.0.3-k8s-alpha0

	k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/client-go v0.0.0-20191016111102-bec269661e48
	k8s.io/utils v0.0.0-20200124190032-861946025e34 // indirect
)