// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package translation

import (
	"slices"
	"sort"
	"testing"

	envoy_config_core_v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	envoy_config_listener "github.com/cilium/proxy/go/envoy/config/listener/v3"
	httpConnectionManagerv3 "github.com/cilium/proxy/go/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/cilium/proxy/go/envoy/extensions/transport_sockets/tls/v3"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/cilium/cilium/operator/pkg/model"
)

func TestNewHTTPListener(t *testing.T) {
	t.Run("without TLS", func(t *testing.T) {
		res, err := NewHTTPListener("dummy-name", "dummy-secret-namespace", nil, nil)
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)

		require.Equal(t, "dummy-name", listener.Name)
		require.Len(t, listener.GetListenerFilters(), 1)
		require.Len(t, listener.GetFilterChains(), 1)
	})

	t.Run("with default XffNumTrustedHops", func(t *testing.T) {
		res, err := NewHTTPListener("dummy-name", "dummy-secret-namespace", nil, nil)
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)
		require.Len(t, listener.GetFilterChains(), 1)
		require.Len(t, listener.GetFilterChains()[0].Filters, 1)
		httpConnectionManager := &httpConnectionManagerv3.HttpConnectionManager{}
		err = proto.Unmarshal(listener.GetFilterChains()[0].Filters[0].ConfigType.(*envoy_config_listener.Filter_TypedConfig).TypedConfig.Value, httpConnectionManager)
		require.Nil(t, err)
		// Default value is 0
		require.Equal(t, uint32(0), httpConnectionManager.XffNumTrustedHops)
	})

	t.Run("without TLS with Proxy Protocol", func(t *testing.T) {
		res, err := NewHTTPListener("dummy-name", "dummy-secret-namespace", nil, nil, WithProxyProtocol())
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)

		require.Equal(t, "dummy-name", listener.Name)

		listenerNames := []string{}
		for _, l := range listener.GetListenerFilters() {
			listenerNames = append(listenerNames, l.Name)
		}
		slices.Sort(listenerNames)
		require.Equal(t, []string{proxyProtocolType, tlsInspectorType}, listenerNames)
		require.Len(t, listener.GetFilterChains(), 1)
	})

	t.Run("TLS filterchain order", func(t *testing.T) {
		res1, err1 := NewHTTPListener("dummy-name", "dummy-secret-namespace", map[model.TLSSecret][]string{
			{Name: "dummy-secret-1", Namespace: "dummy-namespace"}: {"dummy.server.com"},
			{Name: "dummy-secret-2", Namespace: "dummy-namespace"}: {"dummy.anotherserver.com"},
		}, nil)
		res2, err2 := NewHTTPListener("dummy-name", "dummy-secret-namespace", map[model.TLSSecret][]string{
			{Name: "dummy-secret-2", Namespace: "dummy-namespace"}: {"dummy.anotherserver.com"},
			{Name: "dummy-secret-1", Namespace: "dummy-namespace"}: {"dummy.server.com"},
		}, nil)

		require.NoError(t, err1)
		require.NoError(t, err2)

		diffOutput := cmp.Diff(res1, res2, protocmp.Transform())
		if len(diffOutput) != 0 {
			t.Errorf("Listeners filterchain order did not match:\n%s\n", diffOutput)
		}
	})

	t.Run("TLS", func(t *testing.T) {
		res, err := NewHTTPListener("dummy-name", "dummy-secret-namespace", map[model.TLSSecret][]string{
			{Name: "dummy-secret-1", Namespace: "dummy-namespace"}: {"dummy.server.com"},
			{Name: "dummy-secret-2", Namespace: "dummy-namespace"}: {"dummy.anotherserver.com"},
		}, nil)
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)

		require.Equal(t, "dummy-name", listener.Name)
		require.Len(t, listener.GetListenerFilters(), 1)
		require.Len(t, listener.GetFilterChains(), 3)
		require.Equal(t, "raw_buffer", listener.GetFilterChains()[0].GetFilterChainMatch().TransportProtocol)
		require.Equal(t, "tls", listener.GetFilterChains()[1].GetFilterChainMatch().TransportProtocol)
		require.Equal(t, "tls", listener.GetFilterChains()[2].GetFilterChainMatch().TransportProtocol)
		require.Len(t, listener.GetFilterChains()[1].GetFilters(), 1)
		require.Equal(t, []string{"dummy.server.com"}, listener.GetFilterChains()[1].GetFilterChainMatch().ServerNames)
		require.Equal(t, []string{"dummy.anotherserver.com"}, listener.GetFilterChains()[2].GetFilterChainMatch().ServerNames)

		downStreamTLS := &envoy_extensions_transport_sockets_tls_v3.DownstreamTlsContext{}
		err = proto.Unmarshal(listener.FilterChains[1].TransportSocket.ConfigType.(*envoy_config_core_v3.TransportSocket_TypedConfig).TypedConfig.Value, downStreamTLS)
		require.NoError(t, err)

		var secretNames []string
		require.Len(t, downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs, 1)
		sort.Slice(downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs, func(i, j int) bool {
			return downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[i].Name < downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[j].Name
		})
		secretNames = append(secretNames, downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[0].GetName())

		err = proto.Unmarshal(listener.FilterChains[2].TransportSocket.ConfigType.(*envoy_config_core_v3.TransportSocket_TypedConfig).TypedConfig.Value, downStreamTLS)
		require.NoError(t, err)

		require.Len(t, downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs, 1)
		sort.Slice(downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs, func(i, j int) bool {
			return downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[i].Name < downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[j].Name
		})
		secretNames = append(secretNames, downStreamTLS.CommonTlsContext.TlsCertificateSdsSecretConfigs[0].GetName())

		sort.Strings(secretNames)
		require.Equal(t, "dummy-secret-namespace/dummy-namespace-dummy-secret-1", secretNames[0])
		require.Equal(t, "dummy-secret-namespace/dummy-namespace-dummy-secret-2", secretNames[1])
	})
}

func TestNewSNIListener(t *testing.T) {
	t.Run("normal SNI listener", func(t *testing.T) {
		res, err := NewSNIListener("dummy-name",
			map[string][]string{
				"foo-namespace/dummy-service:443": {
					"foo.bar",
				},
				"dummy-namespace/dummy-service:443": {
					"example.org",
					"example.com",
				},
			},
			nil,
		)
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)

		require.Equal(t, "dummy-name", listener.Name)
		require.Len(t, listener.GetListenerFilters(), 1)
		require.Len(t, listener.GetFilterChains(), 2)
		require.Equal(t, []string{"example.com", "example.org"}, listener.GetFilterChains()[0].FilterChainMatch.ServerNames)
		require.Equal(t, []string{"foo.bar"}, listener.GetFilterChains()[1].FilterChainMatch.ServerNames)
	})

	t.Run("SNI listener with stable filterchain sort-order", func(t *testing.T) {
		res1, err1 := NewSNIListener("dummy-name",
			map[string][]string{
				"dummy-namespace/dummy-service:443": {
					"example.org",
					"example.com",
				},
				"foo-namespace/dummy-service:443": {
					"foo.bar",
				},
			},
			nil,
		)
		res2, err2 := NewSNIListener("dummy-name",
			map[string][]string{
				"foo-namespace/dummy-service:443": {
					"foo.bar",
				},
				"dummy-namespace/dummy-service:443": {
					"example.org",
					"example.com",
				},
			},
			nil,
		)
		require.Nil(t, err1)
		require.Nil(t, err2)

		diffOutput := cmp.Diff(res1, res2, protocmp.Transform())
		if len(diffOutput) != 0 {
			t.Errorf("Listeners did not match:\n%s\n", diffOutput)
		}
	})

	t.Run("normal SNI listener with Proxy Protocol", func(t *testing.T) {
		res, err := NewSNIListener("dummy-name", map[string][]string{"dummy-namespace/dummy-service:443": {"example.org", "example.com"}}, nil, WithProxyProtocol())
		require.Nil(t, err)

		listener := &envoy_config_listener.Listener{}
		err = proto.Unmarshal(res.Value, listener)
		require.Nil(t, err)

		require.Equal(t, "dummy-name", listener.Name)
		listenerNames := []string{}
		for _, l := range listener.GetListenerFilters() {
			listenerNames = append(listenerNames, l.Name)
		}
		slices.Sort(listenerNames)
		require.Equal(t, []string{proxyProtocolType, tlsInspectorType}, listenerNames)
		require.Len(t, listener.GetFilterChains(), 1)
		require.Len(t, listener.GetFilterChains()[0].FilterChainMatch.ServerNames, 2)
	})
}

func TestGetHostNetworkListenerAddresses(t *testing.T) {
	testCases := []struct {
		desc                       string
		ports                      []uint32
		ipv4Enabled                bool
		ipv6Enabled                bool
		expectedPrimaryAdress      *envoy_config_core_v3.Address
		expectedAdditionalAdresses []*envoy_config_listener.AdditionalAddress
	}{
		{
			desc:                       "No ports - no address",
			ipv4Enabled:                true,
			ipv6Enabled:                true,
			expectedPrimaryAdress:      nil,
			expectedAdditionalAdresses: nil,
		},
		{
			desc:                       "No IP family - no address",
			ports:                      []uint32{55555},
			expectedPrimaryAdress:      nil,
			expectedAdditionalAdresses: nil,
		},
		{
			desc:        "IPv4 only",
			ports:       []uint32{55555},
			ipv4Enabled: true,
			ipv6Enabled: false,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "0.0.0.0",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 55555,
						},
					},
				},
			},
			expectedAdditionalAdresses: nil,
		},
		{
			desc:        "IPv6 only",
			ports:       []uint32{55555},
			ipv4Enabled: false,
			ipv6Enabled: true,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "::",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 55555,
						},
					},
				},
			},
			expectedAdditionalAdresses: nil,
		},
		{
			desc:        "IPv4 & IPv6",
			ports:       []uint32{55555},
			ipv4Enabled: true,
			ipv6Enabled: true,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "0.0.0.0",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 55555,
						},
					},
				},
			},
			expectedAdditionalAdresses: []*envoy_config_listener.AdditionalAddress{
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "::",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 55555,
								},
							},
						},
					},
				},
			},
		},
		{
			desc:        "IPv4 only with multiple ports",
			ports:       []uint32{44444, 55555},
			ipv4Enabled: true,
			ipv6Enabled: false,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "0.0.0.0",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 44444,
						},
					},
				},
			},
			expectedAdditionalAdresses: []*envoy_config_listener.AdditionalAddress{
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "0.0.0.0",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 55555,
								},
							},
						},
					},
				},
			},
		},
		{
			desc:        "IPv6 only with multiple ports",
			ports:       []uint32{44444, 55555},
			ipv4Enabled: false,
			ipv6Enabled: true,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "::",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 44444,
						},
					},
				},
			},
			expectedAdditionalAdresses: []*envoy_config_listener.AdditionalAddress{
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "::",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 55555,
								},
							},
						},
					},
				},
			},
		},
		{
			desc:        "IPv4 & IPv6 with multiple ports",
			ports:       []uint32{44444, 55555},
			ipv4Enabled: true,
			ipv6Enabled: true,
			expectedPrimaryAdress: &envoy_config_core_v3.Address{
				Address: &envoy_config_core_v3.Address_SocketAddress{
					SocketAddress: &envoy_config_core_v3.SocketAddress{
						Protocol: envoy_config_core_v3.SocketAddress_TCP,
						Address:  "0.0.0.0",
						PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
							PortValue: 44444,
						},
					},
				},
			},
			expectedAdditionalAdresses: []*envoy_config_listener.AdditionalAddress{
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "::",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 44444,
								},
							},
						},
					},
				},
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "0.0.0.0",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 55555,
								},
							},
						},
					},
				},
				{
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  "::",
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: 55555,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			primaryAddress, additionalAddresses := getHostNetworkListenerAddresses(tC.ports, tC.ipv4Enabled, tC.ipv6Enabled)

			assert.Equal(t, tC.expectedPrimaryAdress, primaryAddress)
			assert.Equal(t, tC.expectedAdditionalAdresses, additionalAddresses)
		})
	}
}

func TestWithHostNetworkPortSorted(t *testing.T) {
	modifiedEnvoyListener1 := WithHostNetworkPort([]model.Listener{&model.HTTPListener{Port: 80}, &model.HTTPListener{Port: 443}}, true, true)(&envoy_config_listener.Listener{})
	modifiedEnvoyListener2 := WithHostNetworkPort([]model.Listener{&model.HTTPListener{Port: 443}, &model.HTTPListener{Port: 80}}, true, true)(&envoy_config_listener.Listener{})

	diffOutput := cmp.Diff(modifiedEnvoyListener1, modifiedEnvoyListener2, protocmp.Transform())
	if len(diffOutput) != 0 {
		t.Errorf("Modified Envoy Listeners did not match for different order of http listener ports:\n%s\n", diffOutput)
	}
}
