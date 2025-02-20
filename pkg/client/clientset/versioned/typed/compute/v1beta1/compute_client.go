// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ComputeV1beta1Interface interface {
	RESTClient() rest.Interface
	ComputeAddressesGetter
	ComputeBackendBucketsGetter
	ComputeBackendServicesGetter
	ComputeDisksGetter
	ComputeExternalVPNGatewaysGetter
	ComputeFirewallsGetter
	ComputeFirewallPoliciesGetter
	ComputeForwardingRulesGetter
	ComputeHTTPHealthChecksGetter
	ComputeHTTPSHealthChecksGetter
	ComputeHealthChecksGetter
	ComputeImagesGetter
	ComputeInstancesGetter
	ComputeInstanceGroupsGetter
	ComputeInstanceGroupManagersGetter
	ComputeInstanceTemplatesGetter
	ComputeInterconnectAttachmentsGetter
	ComputeNetworksGetter
	ComputeNetworkEndpointGroupsGetter
	ComputeNetworkPeeringsGetter
	ComputeNodeGroupsGetter
	ComputeNodeTemplatesGetter
	ComputeProjectMetadatasGetter
	ComputeReservationsGetter
	ComputeResourcePoliciesGetter
	ComputeRoutesGetter
	ComputeRoutersGetter
	ComputeRouterInterfacesGetter
	ComputeRouterNATsGetter
	ComputeRouterPeersGetter
	ComputeSSLCertificatesGetter
	ComputeSSLPoliciesGetter
	ComputeSecurityPoliciesGetter
	ComputeSharedVPCHostProjectsGetter
	ComputeSharedVPCServiceProjectsGetter
	ComputeSnapshotsGetter
	ComputeSubnetworksGetter
	ComputeTargetGRPCProxiesGetter
	ComputeTargetHTTPProxiesGetter
	ComputeTargetHTTPSProxiesGetter
	ComputeTargetInstancesGetter
	ComputeTargetPoolsGetter
	ComputeTargetSSLProxiesGetter
	ComputeTargetTCPProxiesGetter
	ComputeTargetVPNGatewaysGetter
	ComputeURLMapsGetter
	ComputeVPNGatewaysGetter
	ComputeVPNTunnelsGetter
}

// ComputeV1beta1Client is used to interact with features provided by the compute.cnrm.cloud.google.com group.
type ComputeV1beta1Client struct {
	restClient rest.Interface
}

func (c *ComputeV1beta1Client) ComputeAddresses(namespace string) ComputeAddressInterface {
	return newComputeAddresses(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeBackendBuckets(namespace string) ComputeBackendBucketInterface {
	return newComputeBackendBuckets(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeBackendServices(namespace string) ComputeBackendServiceInterface {
	return newComputeBackendServices(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeDisks(namespace string) ComputeDiskInterface {
	return newComputeDisks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeExternalVPNGateways(namespace string) ComputeExternalVPNGatewayInterface {
	return newComputeExternalVPNGateways(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeFirewalls(namespace string) ComputeFirewallInterface {
	return newComputeFirewalls(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeFirewallPolicies(namespace string) ComputeFirewallPolicyInterface {
	return newComputeFirewallPolicies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeForwardingRules(namespace string) ComputeForwardingRuleInterface {
	return newComputeForwardingRules(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeHTTPHealthChecks(namespace string) ComputeHTTPHealthCheckInterface {
	return newComputeHTTPHealthChecks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeHTTPSHealthChecks(namespace string) ComputeHTTPSHealthCheckInterface {
	return newComputeHTTPSHealthChecks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeHealthChecks(namespace string) ComputeHealthCheckInterface {
	return newComputeHealthChecks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeImages(namespace string) ComputeImageInterface {
	return newComputeImages(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeInstances(namespace string) ComputeInstanceInterface {
	return newComputeInstances(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeInstanceGroups(namespace string) ComputeInstanceGroupInterface {
	return newComputeInstanceGroups(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeInstanceGroupManagers(namespace string) ComputeInstanceGroupManagerInterface {
	return newComputeInstanceGroupManagers(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeInstanceTemplates(namespace string) ComputeInstanceTemplateInterface {
	return newComputeInstanceTemplates(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeInterconnectAttachments(namespace string) ComputeInterconnectAttachmentInterface {
	return newComputeInterconnectAttachments(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeNetworks(namespace string) ComputeNetworkInterface {
	return newComputeNetworks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeNetworkEndpointGroups(namespace string) ComputeNetworkEndpointGroupInterface {
	return newComputeNetworkEndpointGroups(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeNetworkPeerings(namespace string) ComputeNetworkPeeringInterface {
	return newComputeNetworkPeerings(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeNodeGroups(namespace string) ComputeNodeGroupInterface {
	return newComputeNodeGroups(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeNodeTemplates(namespace string) ComputeNodeTemplateInterface {
	return newComputeNodeTemplates(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeProjectMetadatas(namespace string) ComputeProjectMetadataInterface {
	return newComputeProjectMetadatas(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeReservations(namespace string) ComputeReservationInterface {
	return newComputeReservations(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeResourcePolicies(namespace string) ComputeResourcePolicyInterface {
	return newComputeResourcePolicies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeRoutes(namespace string) ComputeRouteInterface {
	return newComputeRoutes(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeRouters(namespace string) ComputeRouterInterface {
	return newComputeRouters(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeRouterInterfaces(namespace string) ComputeRouterInterfaceInterface {
	return newComputeRouterInterfaces(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeRouterNATs(namespace string) ComputeRouterNATInterface {
	return newComputeRouterNATs(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeRouterPeers(namespace string) ComputeRouterPeerInterface {
	return newComputeRouterPeers(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSSLCertificates(namespace string) ComputeSSLCertificateInterface {
	return newComputeSSLCertificates(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSSLPolicies(namespace string) ComputeSSLPolicyInterface {
	return newComputeSSLPolicies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSecurityPolicies(namespace string) ComputeSecurityPolicyInterface {
	return newComputeSecurityPolicies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSharedVPCHostProjects(namespace string) ComputeSharedVPCHostProjectInterface {
	return newComputeSharedVPCHostProjects(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSharedVPCServiceProjects(namespace string) ComputeSharedVPCServiceProjectInterface {
	return newComputeSharedVPCServiceProjects(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSnapshots(namespace string) ComputeSnapshotInterface {
	return newComputeSnapshots(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeSubnetworks(namespace string) ComputeSubnetworkInterface {
	return newComputeSubnetworks(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetGRPCProxies(namespace string) ComputeTargetGRPCProxyInterface {
	return newComputeTargetGRPCProxies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetHTTPProxies(namespace string) ComputeTargetHTTPProxyInterface {
	return newComputeTargetHTTPProxies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetHTTPSProxies(namespace string) ComputeTargetHTTPSProxyInterface {
	return newComputeTargetHTTPSProxies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetInstances(namespace string) ComputeTargetInstanceInterface {
	return newComputeTargetInstances(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetPools(namespace string) ComputeTargetPoolInterface {
	return newComputeTargetPools(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetSSLProxies(namespace string) ComputeTargetSSLProxyInterface {
	return newComputeTargetSSLProxies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetTCPProxies(namespace string) ComputeTargetTCPProxyInterface {
	return newComputeTargetTCPProxies(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeTargetVPNGateways(namespace string) ComputeTargetVPNGatewayInterface {
	return newComputeTargetVPNGateways(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeURLMaps(namespace string) ComputeURLMapInterface {
	return newComputeURLMaps(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeVPNGateways(namespace string) ComputeVPNGatewayInterface {
	return newComputeVPNGateways(c, namespace)
}

func (c *ComputeV1beta1Client) ComputeVPNTunnels(namespace string) ComputeVPNTunnelInterface {
	return newComputeVPNTunnels(c, namespace)
}

// NewForConfig creates a new ComputeV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*ComputeV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ComputeV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new ComputeV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ComputeV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ComputeV1beta1Client for the given RESTClient.
func New(c rest.Interface) *ComputeV1beta1Client {
	return &ComputeV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ComputeV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
