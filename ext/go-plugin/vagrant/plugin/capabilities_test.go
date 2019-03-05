package plugin

import (
	"testing"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/vagrant/ext/go-plugin/vagrant"
)

func TestCapabilities_GuestCapabilities(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &GuestCapabilitiesPlugin{Impl: &MockGuestCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(GuestCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}
	resp, err := impl.GuestCapabilities()
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	if len(resp) != 1 {
		t.Fatalf("length %d != 1", len(resp))
	}
	if resp[0].Name != "test_cap" {
		t.Errorf("name - %s != test_cap", resp[0].Name)
	}
	if resp[0].Platform != "testOS" {
		t.Errorf("platform - %s != testOS", resp[0].Platform)
	}
}

func TestCapabilities_GuestCapability(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &GuestCapabilitiesPlugin{Impl: &MockGuestCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(GuestCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.SystemCapability{
		Name:     "test_cap",
		Platform: "TestOS"}
	m := &vagrant.Machine{}
	args := []string{"test_value", "next_test_value"}

	resp, err := impl.GuestCapability(cap, args, m)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
	if result[1] != "test_value" {
		t.Errorf("%s != test_value", result[1])
	}
}

func TestCapabilities_GuestCapability_noargs(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &GuestCapabilitiesPlugin{Impl: &MockGuestCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(GuestCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.SystemCapability{
		Name:     "test_cap",
		Platform: "TestOS"}
	m := &vagrant.Machine{}
	var args interface{}
	args = nil

	resp, err := impl.GuestCapability(cap, args, m)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
}

func TestCapabilities_HostCapabilities(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &HostCapabilitiesPlugin{Impl: &MockHostCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(HostCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}
	resp, err := impl.HostCapabilities()
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	if len(resp) != 1 {
		t.Fatalf("length %d != 1", len(resp))
	}
	if resp[0].Name != "test_cap" {
		t.Errorf("name - %s != test_cap", resp[0].Name)
	}
	if resp[0].Platform != "testOS" {
		t.Errorf("platform - %s != testOS", resp[0].Platform)
	}
}

func TestCapabilities_HostCapability(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &HostCapabilitiesPlugin{Impl: &MockHostCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(HostCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.SystemCapability{
		Name:     "test_cap",
		Platform: "TestOS"}
	e := &vagrant.Environment{}
	args := []string{"test_value", "next_test_value"}

	resp, err := impl.HostCapability(cap, args, e)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
	if result[1] != "test_value" {
		t.Errorf("%s != test_value", result[1])
	}
}

func TestCapabilities_HostCapability_noargs(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &HostCapabilitiesPlugin{Impl: &MockHostCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(HostCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.SystemCapability{
		Name:     "test_cap",
		Platform: "TestOS"}
	e := &vagrant.Environment{}
	var args interface{}
	args = nil

	resp, err := impl.HostCapability(cap, args, e)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
}

func TestCapabilities_ProviderCapabilities(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &ProviderCapabilitiesPlugin{Impl: &MockProviderCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(ProviderCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}
	resp, err := impl.ProviderCapabilities()
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	if len(resp) != 1 {
		t.Fatalf("length %d != 1", len(resp))
	}
	if resp[0].Name != "test_cap" {
		t.Errorf("name - %s != test_cap", resp[0].Name)
	}
	if resp[0].Provider != "testProvider" {
		t.Errorf("provider - %s != testProvdier", resp[0].Provider)
	}
}

func TestCapabilities_ProviderCapability(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &ProviderCapabilitiesPlugin{Impl: &MockProviderCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(ProviderCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.ProviderCapability{
		Name:     "test_cap",
		Provider: "test_provider"}
	m := &vagrant.Machine{}
	args := []string{"test_value", "next_test_value"}

	resp, err := impl.ProviderCapability(cap, args, m)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
	if result[1] != "test_value" {
		t.Errorf("%s != test_value", result[1])
	}
}

func TestCapabilities_ProviderCapability_noargs(t *testing.T) {
	client, server := plugin.TestPluginGRPCConn(t, map[string]plugin.Plugin{
		"caps": &ProviderCapabilitiesPlugin{Impl: &MockProviderCapabilities{}}})
	defer server.Stop()
	defer client.Close()

	raw, err := client.Dispense("caps")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	impl, ok := raw.(ProviderCapabilities)
	if !ok {
		t.Fatalf("bad %#v", raw)
	}

	cap := &vagrant.ProviderCapability{
		Name:     "test_cap",
		Provider: "test_provider"}
	m := &vagrant.Machine{}
	var args interface{}
	args = nil

	resp, err := impl.ProviderCapability(cap, args, m)
	if err != nil {
		t.Fatalf("bad resp: %s", err)
	}
	result, ok := resp.([]interface{})
	if !ok {
		t.Fatalf("bad %#v", result)
	}
	if result[0] != "test_cap" {
		t.Errorf("%s != test_cap", result[0])
	}
}