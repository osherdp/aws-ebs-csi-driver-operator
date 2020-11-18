// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// APIServers returns a APIServerInformer.
	APIServers() APIServerInformer
	// Authentications returns a AuthenticationInformer.
	Authentications() AuthenticationInformer
	// Builds returns a BuildInformer.
	Builds() BuildInformer
	// ClusterOperators returns a ClusterOperatorInformer.
	ClusterOperators() ClusterOperatorInformer
	// ClusterVersions returns a ClusterVersionInformer.
	ClusterVersions() ClusterVersionInformer
	// Consoles returns a ConsoleInformer.
	Consoles() ConsoleInformer
	// DNSes returns a DNSInformer.
	DNSes() DNSInformer
	// FeatureGates returns a FeatureGateInformer.
	FeatureGates() FeatureGateInformer
	// Images returns a ImageInformer.
	Images() ImageInformer
	// Infrastructures returns a InfrastructureInformer.
	Infrastructures() InfrastructureInformer
	// Ingresses returns a IngressInformer.
	Ingresses() IngressInformer
	// Networks returns a NetworkInformer.
	Networks() NetworkInformer
	// OAuths returns a OAuthInformer.
	OAuths() OAuthInformer
	// OperatorHubs returns a OperatorHubInformer.
	OperatorHubs() OperatorHubInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// Proxies returns a ProxyInformer.
	Proxies() ProxyInformer
	// Schedulers returns a SchedulerInformer.
	Schedulers() SchedulerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// APIServers returns a APIServerInformer.
func (v *version) APIServers() APIServerInformer {
	return &aPIServerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Authentications returns a AuthenticationInformer.
func (v *version) Authentications() AuthenticationInformer {
	return &authenticationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Builds returns a BuildInformer.
func (v *version) Builds() BuildInformer {
	return &buildInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterOperators returns a ClusterOperatorInformer.
func (v *version) ClusterOperators() ClusterOperatorInformer {
	return &clusterOperatorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterVersions returns a ClusterVersionInformer.
func (v *version) ClusterVersions() ClusterVersionInformer {
	return &clusterVersionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Consoles returns a ConsoleInformer.
func (v *version) Consoles() ConsoleInformer {
	return &consoleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DNSes returns a DNSInformer.
func (v *version) DNSes() DNSInformer {
	return &dNSInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// FeatureGates returns a FeatureGateInformer.
func (v *version) FeatureGates() FeatureGateInformer {
	return &featureGateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Images returns a ImageInformer.
func (v *version) Images() ImageInformer {
	return &imageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Infrastructures returns a InfrastructureInformer.
func (v *version) Infrastructures() InfrastructureInformer {
	return &infrastructureInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Ingresses returns a IngressInformer.
func (v *version) Ingresses() IngressInformer {
	return &ingressInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Networks returns a NetworkInformer.
func (v *version) Networks() NetworkInformer {
	return &networkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OAuths returns a OAuthInformer.
func (v *version) OAuths() OAuthInformer {
	return &oAuthInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OperatorHubs returns a OperatorHubInformer.
func (v *version) OperatorHubs() OperatorHubInformer {
	return &operatorHubInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Proxies returns a ProxyInformer.
func (v *version) Proxies() ProxyInformer {
	return &proxyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Schedulers returns a SchedulerInformer.
func (v *version) Schedulers() SchedulerInformer {
	return &schedulerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}