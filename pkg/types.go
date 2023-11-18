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
    Options     Options     `mapstructure:"globaloptions,omitempty"`
    Hosts       []Host      `mapstructure:"hosts,omitempty"`
    Groups      []Group     `mapstructure:"groups,omitempty"`
    // names       []string
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


type Host struct{
	Options  Options `mapstructure:",squash,omitempty"`
    Name     string  `mapstructure:"name"`
}

func (host *Host) inheritOptions(groupOptions Options) error {
    inheritOptions(&host.Options, &groupOptions)
    return nil 
}

type Group struct{
	Options         Options `mapstructure:"options,omitempty"`
    Hosts           []Host  `mapstructure:"hosts,omitempty"`
    Name            string  `mapstructure:"name"`
    Subgroups       []Group `mapstructure:"subgroups,omitempty"`
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
    AddKeysToAgent                       Option          `mapstructure:"addkeystoagent,omitempty"`
    AddressFamily                        Option        `mapstructure:"addressfamily,omitempty"`
    BatchMode                            Option          `mapstructure:"batchmode,omitempty"`
    BindAddress                          Option        `mapstructure:"bindaddress,omitempty"`
    CanonicalDomains                     Option        `mapstructure:"canonicaldomains,omitempty"`
    CanonicalizeFallbackLocal            Option          `mapstructure:"CanonicalizeFallbackLocal,omitempty"`
    CanonicalizeHostname                 Option        `mapstructure:"canonicalizehostname,omitempty"`
    CanonicalizeMaxDots                  Option           `mapstructure:"canonicalizemaxdots,omitempty"`
    CanonicalizePermittedCNAMEs          Option        `mapstructure:"CanonicalizePermittedCNAMEs,omitempty"`
    CASignatureAlgorithms                Option        `mapstructure:"casignaturealgorithms,omitempty"`
    CertificateFile                      Option        `mapstructure:"certificatefile,omitempty"`
    CheckHostIP                          Option        `mapstructure:"checkHostip,omitempty"`
    Ciphers                              Option        `mapstructure:"ciphers,omitempty"`
    ClearAllForwardings                  Option          `mapstructure:"clearallforwardings,omitempty"`
    Compression                          Option        `mapstructure:"compression,omitempty"`
    ConnectionAttempts                   Option           `mapstructure:"connectionattempts,omitempty"`
    ConnectTimeout                       Option `mapstructure:"connecttimeout,omitempty"`
    ControlMaster                        Option        `mapstructure:"controlmaster,omitempty"`
    ControlPath                          Option        `mapstructure:"controlpath,omitempty"`
    ControlPersist                       Option        `mapstructure:"controlpersist,omitempty"`
    DynamicForward                       Option        `mapstructure:"dynamicforward,omitempty"`
    EscapeChar                           Option        `mapstructure:"escapechar,omitempty"`
    ExitOnForwardFailure                 Option          `mapstructure:"exitonforwardfailure,omitempty"`
    FingerprintHash                      Option        `mapstructure:"fingerprinthash,omitempty"`
    ForkAfterAuthentication              Option          `mapstructure:"forkafterauthentication,omitempty"`
    ForwardAgent                         Option          `mapstructure:"forwardagent,omitempty"`
    ForwardX11                           Option          `mapstructure:"forwardx11,omitempty"`
    ForwardX11Timeout                    Option `mapstructure:"forwardx11timeout,omitempty"`
    ForwardX11Trusted                    Option          `mapstructure:"forwardx11trusted,omitempty"`
    GatewayPorts                         Option        `mapstructure:"gatewayports,omitempty"`
    GlobalKnownHostsFile                 Option        `mapstructure:"globalknownhostsfile,omitempty"`
    GSSAPIAuthentication                 Option          `mapstructure:"gssapiauthentication,omitempty"`
    GSSAPIDelegateCredentials            Option          `mapstructure:"gssapidelegatecredentials,omitempty"`
    HashKnownHosts                       Option          `mapstructure:"hashknownhosts,omitempty"`
    Host                                 Option        `mapstructure:"host,omitempty"`
    HostbasedAcceptedAlgorithms          Option        `mapstructure:"hostbasedacceptedalgorithms,omitempty"`
    HostbasedAuthentication              Option          `mapstructure:"hostbasedauthentication,omitempty"`
    HostKeyAlgorithms                    Option        `mapstructure:"hostkeyalgorithms,omitempty"`
    HostKeyAlias                         Option        `mapstructure:"hostkeyalias,omitempty"`
    Hostname                             Option        `mapstructure:"hostname,omitempty"`
    IdentitiesOnly                       Option          `mapstructure:"IdentitiesOnly,omitempty"`
    IdentityAgent                        Option        `mapstructure:"identityagent,omitempty"`
    IdentityFile                         Option        `mapstructure:"identityfile,omitempty"`
    IPQoS                                Option        `mapstructure:"ipqos,omitempty"`
    KbdInteractiveDevices                Option        `mapstructure:"kbdinteractivedevices,omitempty"`
    KexAlgorithms                        Option        `mapstructure:"kexalgorithms,omitempty"`
    KnownHostsCommand                    Option        `mapstructure:"knownhostscommand,omitempty"`
    LocalCommand                         Option        `mapstructure:"localcommand,omitempty"`
    LocalForward                         Option        `mapstructure:"localforward,omitempty"`
    LogLevel                             Option        `mapstructure:"loglevel,omitempty"`
    MACs                                 Option        `mapstructure:"macs,omitempty"`
    // Match                                Option        `mapstructure:"match,omitempty"`
    NoHostAuthenticationForLocalhost     Option          `mapstructure:"nohostauthenticationforlocalhost,omitempty"`
    NumberOfPasswordPrompts              Option           `mapstructure:"numberofpasswordprompts,omitempty"`
    PasswordAuthentication               Option          `mapstructure:"passwordauthentication,omitempty"`
    PermitLocalCommand                   Option          `mapstructure:"permitlocalcommand,omitempty"`
    PermitRemoteOpen                     Option          `mapstructure:"permitremoteopen,omitempty"`
    PKCS11Provider                       Option        `mapstructure:"pkcs11provider,omitempty"`
    Port                                 Option           `mapstructure:"port,omitempty"`
    PreferredAuthentications             Option        `mapstructure:"preferredauthentications,omitempty"`
    ProxyCommand                         Option        `mapstructure:"proxycommand,omitempty"`
    ProxyJump                            Option        `mapstructure:"proxyjump,omitempty"`
    ProxyUseFdpass                       Option          `mapstructure:"proxyusefdpass,omitempty"`
    PubkeyAcceptedAlgorithms             Option        `mapstructure:"pubkeyacceptedalgorithms,omitempty"`
    PubkeyAuthentication                 Option          `mapstructure:"pubkeyauthentication,omitempty"`
    RekeyLimit                           Option        `mapstructure:"rekeylimit,omitempty"`
    RemoteCommand                        Option        `mapstructure:"remotecommand,omitempty"`
    RemoteForward                        Option        `mapstructure:"remoteforward,omitempty"`
    RequestTTY                           Option        `mapstructure:"requesttty,omitempty"`
    SendEnv                              Option        `mapstructure:"sendenv,omitempty"`
    ServerAliveInterval                  Option `mapstructure:"serveraliveinterval,omitempty"`
    ServerAliveCountMax                  Option           `mapstructure:"serveralivecountmax,omitempty"`
    SessionType                          Option        `mapstructure:"sessionType,omitempty"`
    SetEnv                               Option        `mapstructure:"setenv,omitempty"`
    StdinNull                            Option          `mapstructure:"stdinnull,omitempty"`
    StreamLocalBindMask                  Option        `mapstructure:"streamlocalbindmask,omitempty"`
    StreamLocalBindUnlink                Option          `mapstructure:"streamlocalbindunlink,omitempty"`
    StrictHostKeyChecking                Option          `mapstructure:"stricthostkeychecking,omitempty"`
    TCPKeepAlive                         Option          `mapstructure:"tcpkeepalive,omitempty"`
    Tunnel                               Option          `mapstructure:"tunnel,omitempty"`
    TunnelDevice                         Option        `mapstructure:"tunneldevice,omitempty"`
    UpdateHostKeys                       Option          `mapstructure:"updatehostkeys,omitempty"`
    UseKeychain                          Option          `mapstructure:"usekeychain,omitempty"`
    User                                 Option        `mapstructure:"user,omitempty"`
    UserKnownHostsFile                   Option        `mapstructure:"userknownhostsfile,omitempty"`
    VerifyHostKeyDNS                     Option          `mapstructure:"verifyhostkeydns,omitempty"`
    VisualHostKey                        Option          `mapstructure:"visualhostkey,omitempty"`
    XAuthLocation                        Option        `mapstructure:"xauthlocation,omitempty"`
}