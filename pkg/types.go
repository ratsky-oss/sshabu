// Copyright (C) 2023  Shovra Nikita, Livitsky Andrey

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sshabu

import (
	"errors"
	"reflect"
)

func inheritOptions(item interface{}, addition interface{}) {
	itemValue := reflect.ValueOf(item).Elem()
	addValue := reflect.ValueOf(addition).Elem()

	for i := 0; i < itemValue.NumField(); i++ {
		if itemValue.Field(i).Interface() == nil {
			itemValue.Field(i).Set(addValue.Field(i))
		}
	}
}

func findNamesInStruct(value reflect.Value, names *[]string) {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		structType := value.Type()

		for i := 0; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			fieldType := structType.Field(i)

			if fieldType.Name == "Name" {
				*names = append(*names, fieldValue.String())
			}

			findNamesInStruct(fieldValue, names)
		}
	} else if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			findNamesInStruct(element, names)
		}
	}
}

type Shabu struct{
    Options     Options     `mapstructure:"globaloptions,omitempty" yaml:"Options,omitempty"`
    Hosts       []Host      `mapstructure:"hosts,omitempty" yaml:"Hosts,omitempty"`
    Groups      []Group     `mapstructure:"groups,omitempty" yaml:"Groups,omitempty"`
}

func (shabu Shabu) FindNamesInShabu() []string {
	var names []string

	findNamesInStruct(reflect.ValueOf(shabu), &names)

	return names
}

func (shabu Shabu) areAllUnique() (bool, *string) {
	seen := make(map[string]bool)
    items := shabu.FindNamesInShabu()
	for _, item := range items {
		if seen[item] {
			return false, &item // The item is not unique
		}
		seen[item] = true
	}

	return true, nil // All items are unique
}


func (shabu *Shabu) Boil() error {
    if uniq, name := shabu.areAllUnique(); !uniq{
        return errors.New("'Name' fields must be unique - '"+ *name +"' aready used")
    }
    for i := range shabu.Groups {
        shabu.Groups[i].solveGroup(shabu.Groups[i].Options)
    }
    return nil
}

func (shabu *Shabu) AddHost(host Host) error {
    var names []string
	findNamesInStruct(reflect.ValueOf(shabu), &names)

    for _, v := range names {
        if host.Name == v {
            return errors.New("'Name' fields must be unique - '"+ host.Name +"' aready used")
        }
    }
    shabu.Hosts = append(shabu.Hosts, host)
    return nil
}


type Host struct{
    Name     string  `mapstructure:"name" yaml:"Name"`
	Options  Options `mapstructure:",squash,omitempty" yaml:",inline,omitempty"`
}

func (host *Host) inheritOptions(groupOptions Options) error {
    inheritOptions(&host.Options, &groupOptions)
    return nil 
}

type Group struct{
    Name            string  `mapstructure:"name" yaml:"Name"`
	Options         Options `mapstructure:"options,omitempty" yaml:"Options,omitempty"`
    Hosts           []Host  `mapstructure:"hosts,omitempty" yaml:"Hosts,omitempty"`
    Subgroups       []Group `mapstructure:"subgroups,omitempty" yaml:"Subgroups,omitempty"`
}

func (group *Group) inheritOptions(parentOptions Options) error {
    inheritOptions(&group.Options, &parentOptions)
    return nil
}

func (group *Group) solveGroup(parentOptions Options) error {
    group.inheritOptions(parentOptions)
    for i := range group.Hosts {
        group.Hosts[i].inheritOptions(group.Options)
    }
    
    for i := range group.Subgroups {
        group.Subgroups[i].solveGroup(group.Options)
    }
    
    return nil
}

type Option interface{}
type Options struct {
    AddKeysToAgent                       Option         `mapstructure:"addkeystoagent,omitempty" yaml:"AddKeysToAgent,omitempty"`
    AddressFamily                        Option         `mapstructure:"addressfamily,omitempty" yaml:"AddressFamily,omitempty"`
    BatchMode                            Option         `mapstructure:"batchmode,omitempty" yaml:"BatchMode,omitempty"`
    BindAddress                          Option         `mapstructure:"bindaddress,omitempty" yaml:"BindAddress,omitempty"`
    CanonicalDomains                     Option         `mapstructure:"canonicaldomains,omitempty" yaml:"CanonicalDomains,omitempty"`
    CanonicalizeFallbackLocal            Option         `mapstructure:"CanonicalizeFallbackLocal,omitempty" yaml:"CanonicalizeFallbackLocal,omitempty"`
    CanonicalizeHostname                 Option         `mapstructure:"canonicalizehostname,omitempty" yaml:"CanonicalizeHostname,omitempty"`
    CanonicalizeMaxDots                  Option         `mapstructure:"canonicalizemaxdots,omitempty" yaml:"CanonicalizeMaxDots,omitempty"`
    CanonicalizePermittedCNAMEs          Option         `mapstructure:"CanonicalizePermittedCNAMEs,omitempty" yaml:"CanonicalizePermittedCNAMEs,omitempty"`
    CASignatureAlgorithms                Option         `mapstructure:"casignaturealgorithms,omitempty" yaml:"CASignatureAlgorithms,omitempty"`
    CertificateFile                      Option         `mapstructure:"certificatefile,omitempty" yaml:"CertificateFile,omitempty"`
    CheckHostIP                          Option         `mapstructure:"checkHostip,omitempty" yaml:"CheckHostIP,omitempty"`
    Ciphers                              Option         `mapstructure:"ciphers,omitempty" yaml:"Ciphers,omitempty"`
    ClearAllForwardings                  Option         `mapstructure:"clearallforwardings,omitempty" yaml:"ClearAllForwardings,omitempty"`
    Compression                          Option         `mapstructure:"compression,omitempty" yaml:"Compression,omitempty"`
    ConnectionAttempts                   Option         `mapstructure:"connectionattempts,omitempty" yaml:"ConnectionAttempts,omitempty"`
    ConnectTimeout                       Option         `mapstructure:"connecttimeout,omitempty" yaml:"ConnectTimeout,omitempty"`
    ControlMaster                        Option         `mapstructure:"controlmaster,omitempty" yaml:"ControlMaster,omitempty"`
    ControlPath                          Option         `mapstructure:"controlpath,omitempty" yaml:"ControlPath,omitempty"`
    ControlPersist                       Option         `mapstructure:"controlpersist,omitempty" yaml:"ControlPersist,omitempty"`
    DynamicForward                       Option         `mapstructure:"dynamicforward,omitempty" yaml:"DynamicForward,omitempty"`
    EscapeChar                           Option         `mapstructure:"escapechar,omitempty" yaml:"EscapeChar,omitempty"`
    ExitOnForwardFailure                 Option         `mapstructure:"exitonforwardfailure,omitempty" yaml:"ExitOnForwardFailure,omitempty"`
    FingerprintHash                      Option         `mapstructure:"fingerprinthash,omitempty" yaml:"FingerprintHash,omitempty"`
    ForkAfterAuthentication              Option          `mapstructure:"forkafterauthentication,omitempty" yaml:"ForkAfterAuthentication,omitempty"`
    ForwardAgent                         Option          `mapstructure:"forwardagent,omitempty" yaml:"ForwardAgent,omitempty"`
    ForwardX11                           Option          `mapstructure:"forwardx11,omitempty" yaml:"ForwardX11,omitempty"`
    ForwardX11Timeout                    Option          `mapstructure:"forwardx11timeout,omitempty" yaml:"ForwardX11Timeout,omitempty"`
    ForwardX11Trusted                    Option          `mapstructure:"forwardx11trusted,omitempty" yaml:"ForwardX11Trusted,omitempty"`
    GatewayPorts                         Option          `mapstructure:"gatewayports,omitempty" yaml:"GatewayPorts,omitempty"`
    GlobalKnownHostsFile                 Option          `mapstructure:"globalknownhostsfile,omitempty" yaml:"GlobalKnownHostsFile,omitempty"`
    GSSAPIAuthentication                 Option          `mapstructure:"gssapiauthentication,omitempty" yaml:"GSSAPIAuthentication,omitempty"`
    GSSAPIDelegateCredentials            Option          `mapstructure:"gssapidelegatecredentials,omitempty" yaml:"GSSAPIDelegateCredentials,omitempty"`
    HashKnownHosts                       Option          `mapstructure:"hashknownhosts,omitempty" yaml:"HashKnownHosts,omitempty"`
    Host                                 Option          `mapstructure:"host,omitempty" yaml:"Host,omitempty"`
    HostbasedAcceptedAlgorithms          Option          `mapstructure:"hostbasedacceptedalgorithms,omitempty" yaml:"HostbasedAcceptedAlgorithms,omitempty"`
    HostbasedAuthentication              Option          `mapstructure:"hostbasedauthentication,omitempty" yaml:"HostbasedAuthentication,omitempty"`
    HostKeyAlgorithms                    Option          `mapstructure:"hostkeyalgorithms,omitempty" yaml:"HostKeyAlgorithms,omitempty"`
    HostKeyAlias                         Option          `mapstructure:"hostkeyalias,omitempty" yaml:"HostKeyAlias,omitempty"`
    Hostname                             Option          `mapstructure:"hostname,omitempty" yaml:"Hostname,omitempty"`
    IdentitiesOnly                       Option          `mapstructure:"IdentitiesOnly,omitempty" yaml:"IdentitiesOnly,omitempty"`
    IdentityAgent                        Option          `mapstructure:"identityagent,omitempty" yaml:"IdentityAgent,omitempty"`
    IdentityFile                         Option          `mapstructure:"identityfile,omitempty" yaml:"IdentityFile,omitempty"`
    IPQoS                                Option          `mapstructure:"ipqos,omitempty" yaml:"IPQoS,omitempty"`
    KbdInteractiveDevices                Option          `mapstructure:"kbdinteractivedevices,omitempty" yaml:"KbdInteractiveDevices,omitempty"`
    KexAlgorithms                        Option          `mapstructure:"kexalgorithms,omitempty" yaml:"KexAlgorithms,omitempty"`
    KnownHostsCommand                    Option          `mapstructure:"knownhostscommand,omitempty" yaml:"KnownHostsCommand,omitempty"`
    LocalCommand                         Option          `mapstructure:"localcommand,omitempty" yaml:"LocalCommand,omitempty"`
    LocalForward                         Option          `mapstructure:"localforward,omitempty" yaml:"LocalForward,omitempty"`
    LogLevel                             Option          `mapstructure:"loglevel,omitempty" yaml:"LogLevel,omitempty"`
    MACs                                 Option          `mapstructure:"macs,omitempty" yaml:"MACs,omitempty"`
    // Match                                Option        `mapstructure:"match,omitempty" yam//l:",omitempty"`
    NoHostAuthenticationForLocalhost     Option          `mapstructure:"nohostauthenticationforlocalhost,omitempty" yaml:"NoHostAuthenticationForLocalhost,omitempty"`
    NumberOfPasswordPrompts              Option          `mapstructure:"numberofpasswordprompts,omitempty" yaml:"NumberOfPasswordPrompts,omitempty"`
    PasswordAuthentication               Option          `mapstructure:"passwordauthentication,omitempty" yaml:"PasswordAuthentication,omitempty"`
    PermitLocalCommand                   Option          `mapstructure:"permitlocalcommand,omitempty" yaml:"PermitLocalCommand,omitempty"`
    PermitRemoteOpen                     Option          `mapstructure:"permitremoteopen,omitempty" yaml:"PermitRemoteOpen,omitempty"`
    PKCS11Provider                       Option          `mapstructure:"pkcs11provider,omitempty" yaml:"PKCS11Provider,omitempty"`
    Port                                 Option          `mapstructure:"port,omitempty" yaml:"Port,omitempty"`
    PreferredAuthentications             Option          `mapstructure:"preferredauthentications,omitempty" yaml:"PreferredAuthentications,omitempty"`
    ProxyCommand                         Option          `mapstructure:"proxycommand,omitempty" yaml:"ProxyCommand,omitempty"`
    ProxyJump                            Option          `mapstructure:"proxyjump,omitempty" yaml:"ProxyJump,omitempty"`
    ProxyUseFdpass                       Option          `mapstructure:"proxyusefdpass,omitempty" yaml:"ProxyUseFdpass,omitempty"`
    PubkeyAcceptedAlgorithms             Option          `mapstructure:"pubkeyacceptedalgorithms,omitempty" yaml:"PubkeyAcceptedAlgorithms,omitempty"`
    PubkeyAuthentication                 Option          `mapstructure:"pubkeyauthentication,omitempty" yaml:"PubkeyAuthentication,omitempty"`
    RekeyLimit                           Option          `mapstructure:"rekeylimit,omitempty" yaml:"RekeyLimit,omitempty"`
    RemoteCommand                        Option          `mapstructure:"remotecommand,omitempty" yaml:"RemoteCommand,omitempty"`
    RemoteForward                        Option          `mapstructure:"remoteforward,omitempty" yaml:"RemoteForward,omitempty"`
    RequestTTY                           Option          `mapstructure:"requesttty,omitempty" yaml:"RequestTTY,omitempty"`
    SendEnv                              Option          `mapstructure:"sendenv,omitempty" yaml:"SendEnv,omitempty"`
    ServerAliveInterval                  Option          `mapstructure:"serveraliveinterval,omitempty" yaml:"ServerAliveInterval,omitempty"`
    ServerAliveCountMax                  Option          `mapstructure:"serveralivecountmax,omitempty" yaml:"ServerAliveCountMax,omitempty"`
    SessionType                          Option          `mapstructure:"sessionType,omitempty" yaml:"SessionType,omitempty"`
    SetEnv                               Option          `mapstructure:"setenv,omitempty" yaml:"SetEnv,omitempty"`
    StdinNull                            Option          `mapstructure:"stdinnull,omitempty" yaml:"StdinNull,omitempty"`
    StreamLocalBindMask                  Option          `mapstructure:"streamlocalbindmask,omitempty" yaml:"StreamLocalBindMask,omitempty"`
    StreamLocalBindUnlink                Option          `mapstructure:"streamlocalbindunlink,omitempty" yaml:"StreamLocalBindUnlink,omitempty"`
    StrictHostKeyChecking                Option          `mapstructure:"stricthostkeychecking,omitempty" yaml:"StrictHostKeyChecking,omitempty"`
    TCPKeepAlive                         Option          `mapstructure:"tcpkeepalive,omitempty" yaml:"TCPKeepAlive,omitempty"`
    Tunnel                               Option          `mapstructure:"tunnel,omitempty" yaml:"Tunnel,omitempty"`
    TunnelDevice                         Option          `mapstructure:"tunneldevice,omitempty" yaml:"TunnelDevice,omitempty"`
    UpdateHostKeys                       Option          `mapstructure:"updatehostkeys,omitempty" yaml:"UpdateHostKeys,omitempty"`
    UseKeychain                          Option          `mapstructure:"usekeychain,omitempty" yaml:"UseKeychain,omitempty"`
    User                                 Option          `mapstructure:"user,omitempty" yaml:"User,omitempty"`
    UserKnownHostsFile                   Option          `mapstructure:"userknownhostsfile,omitempty" yaml:"UserKnownHostsFile,omitempty"`
    VerifyHostKeyDNS                     Option          `mapstructure:"verifyhostkeydns,omitempty" yaml:"VerifyHostKeyDNS,omitempty"`
    VisualHostKey                        Option          `mapstructure:"visualhostkey,omitempty" yaml:"VisualHostKey,omitempty"`
    XAuthLocation                        Option          `mapstructure:"xauthlocation,omitempty" yaml:"XAuthLocation,omitempty"`
}

func GetAvaliableOptions() []string {
        var ava_options []string
        e := reflect.ValueOf(&(Options{})).Elem()
        
        for i := 0; i < e.NumField(); i++ {
            varName := e.Type().Field(i).Name
            ava_options = append(ava_options, varName)
        }
        return ava_options
}