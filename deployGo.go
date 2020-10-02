package main

import (
	"fmt"
	//"github.com/interconnectedcloud/qdr-operator/test/e2e/framework"
	"github.com/rh-messaging/shipshape/pkg/framework"
	"github.com/rh-messaging/shipshape/pkg/apps/qdrouterd/deployment"
	corev1 "k8s.io/api/core/v1"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"k8s.io/client-go/kubernetes"
)

type RouterRoleType string
type PlacementType string

type Address struct {
	Prefix         string `json:"prefix,omitempty"`
	Pattern        string `json:"pattern,omitempty"`
	Distribution   string `json:"distribution,omitempty"`
	Waypoint       bool   `json:"waypoint,omitempty"`
	IngressPhase   *int32 `json:"ingressPhase,omitempty"`
	EgressPhase    *int32 `json:"egressPhase,omitempty"`
	Priority       *int32 `json:"priority,omitempty"`
	EnableFallback bool   `json:"enableFallback,omitempty"`
}

type SslProfile struct {
	Name                string `json:"name,omitempty"`
	Credentials         string `json:"credentials,omitempty"`
	CaCert              string `json:"caCert,omitempty"`
	GenerateCredentials bool   `json:"generateCredentials,omitempty"`
	GenerateCaCert      bool   `json:"generateCaCert,omitempty"`
	MutualAuth          bool   `json:"mutualAuth,omitempty"`
	Ciphers             string `json:"ciphers,omitempty"`
	Protocols           string `json:"protocols,omitempty"`
}

type LinkRoute struct {
	Prefix               string `json:"prefix,omitempty"`
	Pattern              string `json:"pattern,omitempty"`
	Direction            string `json:"direction,omitempty"`
	ContainerId          string `json:"containerId,omitempty"`
	Connection           string `json:"connection,omitempty"`
	AddExternalPrefix    string `json:"addExternalPrefix,omitempty"`
	RemoveExternalPrefix string `json:"removeExternalPrefix,omitempty"`
}

type Listener struct {
	Name             string `json:"name,omitempty"`
	Host             string `json:"host,omitempty"`
	Port             int32  `json:"port"`
	RouteContainer   bool   `json:"routeContainer,omitempty"`
	Http             bool   `json:"http,omitempty"`
	Cost             int32  `json:"cost,omitempty"`
	SslProfile       string `json:"sslProfile,omitempty"`
	SaslMechanisms   string `json:"saslMechanisms,omitempty"`
	AuthenticatePeer bool   `json:"authenticatePeer,omitempty"`
	Expose           bool   `json:"expose,omitempty"`
}

type DeploymentPlanType struct {
	Image        string                      `json:"image,omitempty"`
	Size         int32                       `json:"size,omitempty"`
	Role         RouterRoleType              `json:"role,omitempty"`
	Placement    PlacementType               `json:"placement,omitempty"`
	Resources    corev1.ResourceRequirements `json:"resources,omitempty"`
	Issuer       string                      `json:"issuer,omitempty"`
	LivenessPort int32                       `json:"livenessPort,omitempty"`
	ServiceType  string                      `json:"serviceType,omitempty"`
}

// InterconnectSpec defines the desired state of Interconnect
type InterconnectSpec struct {
	DeploymentPlan        DeploymentPlanType `json:"deploymentPlan,omitempty"`
	Users                 string             `json:"users,omitempty"`
	Listeners             []Listener         `json:"listeners,omitempty"`
	InterRouterListeners  []Listener         `json:"interRouterListeners,omitempty"`
	EdgeListeners         []Listener         `json:"edgeListeners,omitempty"`
	SslProfiles           []SslProfile       `json:"sslProfiles,omitempty"`
	Addresses             []Address          `json:"addresses,omitempty"`
	AutoLinks             []AutoLink         `json:"autoLinks,omitempty"`
	LinkRoutes            []LinkRoute        `json:"linkRoutes,omitempty"`
	Connectors            []Connector        `json:"connectors,omitempty"`
	InterRouterConnectors []Connector        `json:"interRouterConnectors,omitempty"`
	EdgeConnectors        []Connector        `json:"edgeConnectors,omitempty"`
}

type Connector struct {
	Name           string `json:"name,omitempty"`
	Host           string `json:"host"`
	Port           int32  `json:"port"`
	RouteContainer bool   `json:"routeContainer,omitempty"`
	Cost           int32  `json:"cost,omitempty"`
	VerifyHostname bool   `json:"verifyHostname,omitempty"`
	SslProfile     string `json:"sslProfile,omitempty"`
}

type AutoLink struct {
	Address        string `json:"address"`
	Direction      string `json:"direction"`
	ContainerId    string `json:"containerId,omitempty"`
	Connection     string `json:"connection,omitempty"`
	ExternalPrefix string `json:"externalPrefix,omitempty"`
	Phase          *int32 `json:"phase,omitempty"`
	Fallback       bool   `json:"fallback,omitempty"`
}

func defaultInteriorSpec() *v1alpha1.InterconnectSpec {
	return &v1alpha1.InterconnectSpec{
		DeploymentPlan: v1alpha1.DeploymentPlanType{
			Size:      2,
			Image:     "quay.io/interconnectedcloud/qdrouterd:latest",
			Role:      "interior",
			Placement: "Any",
		},
	}
}

func deployInterconnect(ctx *framework.ContextData, icName string, icSpec *v1alpha1.InterconnectSpec) error {
	// Deploying Interconnect using provided context
	if _, err := deployment.CreateInterconnectFromSpec(*ctx, icSpec.DeploymentPlan.Size, icName, *icSpec); err != nil {
		return err
	}
	// Wait for Interconnect deployment
	err := framework.WaitForDeployment(ctx.Clients.KubeClient, ctx.Namespace, icName, int(icSpec.DeploymentPlan.Size), framework.RetryInterval, framework.Timeout)
	return err
}

func deployGo(kubeClient *kubernetes.Clientset) error {
	// Creates an instance of the framework
	Framework := framework.NewFrameworkBuilder("examples-basic").Build()
	ctx := Framework.GetFirstContext()


	IcInteriorEast := defaultInteriorSpec()
	err := deployInterconnect(ctx, "RenaDeploy", IcInteriorEast)
	fmt.Println("Deploy Debug", IcInteriorEast)
	if err != nil {
		fmt.Println("Unable to deploy Interconnect")
	}
	return nil
}