{{- /*
Copyright (C) 2023  Shovra Nikita, Livitsky Andrey

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.*/ -}}
# -----------------------
# RATSKY SSHABU
{{ with .Options }}
{{ include "option" . }}
{{- end -}}
{{- range .Hosts -}}
{{ include "host" . }}
{{- end -}}
{{- range .Groups -}}
{{ include "group" . }}
{{- end -}}
{{- /*----------------------------------------------*/ -}}
{{ define "group" }}
{{- range .Hosts -}}
{{ include "host" . }}
{{- end -}}
{{- range .Subgroups -}}
{{ include "group" . }}
{{- end -}}
{{- end -}}
{{- /*----------------------------------------------*/ -}}
{{ define "host" }}
{{- if .Options.Host }}
Host {{ .Options.Host }}
{{- else }}
Host {{ .Name }}
{{- end -}}
{{- with .Options }}
{{ include "option" . | indent 4 }}
{{- end -}}
{{- end -}}
{{- /*----------------------------------------------*/ -}}
{{ define "option" -}}
{{- if .AddKeysToAgent }}AddKeysToAgent {{ .AddKeysToAgent }}
{{ end }}
{{- if .AddressFamily }}AddressFamily {{ .AddressFamily }}
{{ end }}
{{- if .BatchMode }}BatchMode {{ .BatchMode }}
{{ end }}
{{- if .BindAddress }}BindAddress {{ .BindAddress }}
{{ end }}
{{- if .CanonicalDomains }}CanonicalDomains {{ .CanonicalDomains }}
{{ end }}
{{- if .CanonicalizeFallbackLocal }}CanonicalizeFallbackLocal {{ .CanonicalizeFallbackLocal }}
{{ end }}
{{- if .CanonicalizeHostname }}CanonicalizeHostname {{ .CanonicalizeHostname }}
{{ end }}
{{- if .CanonicalizeMaxDots }}CanonicalizeMaxDots {{ .CanonicalizeMaxDots }}
{{ end }}
{{- if .CanonicalizePermittedCNAMEs }}CanonicalizePermittedCNAMEs {{ .CanonicalizePermittedCNAMEs }}
{{ end }}
{{- if .CASignatureAlgorithms }}CASignatureAlgorithms {{ .CASignatureAlgorithms }}
{{ end }}
{{- if .CertificateFile }}CertificateFile {{ .CertificateFile }}
{{ end }}
{{- if .CheckHostIP }}CheckHostIP {{ .CheckHostIP }}
{{ end }}
{{- if .Ciphers }}Ciphers {{ .Ciphers }}
{{ end }}
{{- if .ClearAllForwardings }}ClearAllForwardings {{ .ClearAllForwardings }}
{{ end }}
{{- if .Compression }}Compression {{ .Compression }}
{{ end }}
{{- if .ConnectionAttempts }}ConnectionAttempts {{ .ConnectionAttempts }}
{{ end }}
{{- if .ConnectTimeout }}ConnectTimeout {{ .ConnectTimeout }}
{{ end }}
{{- if .ControlMaster }}ControlMaster {{ .ControlMaster }}
{{ end }}
{{- if .ControlPath }}ControlPath {{ .ControlPath }}
{{ end }}
{{- if .ControlPersist }}ControlPersist {{ .ControlPersist }}
{{ end }}
{{- if .DynamicForward }}DynamicForward {{ .DynamicForward }}
{{ end }}
{{- if .EscapeChar }}EscapeChar {{ .EscapeChar }}
{{ end }}
{{- if .ExitOnForwardFailure }}ExitOnForwardFailure {{ .ExitOnForwardFailure }}
{{ end }}
{{- if .FingerprintHash }}FingerprintHash {{ .FingerprintHash }}
{{ end }}
{{- if .ForkAfterAuthentication }}ForkAfterAuthentication {{ .ForkAfterAuthentication }}
{{ end }}
{{- if .ForwardAgent }}ForwardAgent {{ .ForwardAgent }}
{{ end }}
{{- if .ForwardX11 }}ForwardX11 {{ .ForwardX11 }}
{{ end }}
{{- if .ForwardX11Timeout }}ForwardX11Timeout {{ .ForwardX11Timeout }}
{{ end }}
{{- if .ForwardX11Trusted }}ForwardX11Trusted {{ .ForwardX11Trusted }}
{{ end }}
{{- if .GatewayPorts }}GatewayPorts {{ .GatewayPorts }}
{{ end }}
{{- if .GlobalKnownHostsFile }}GlobalKnownHostsFile {{ .GlobalKnownHostsFile }}
{{ end }}
{{- if .GSSAPIAuthentication }}GSSAPIAuthentication {{ .GSSAPIAuthentication }}
{{ end }}
{{- if .GSSAPIDelegateCredentials }}GSSAPIDelegateCredentials {{ .GSSAPIDelegateCredentials }}
{{ end }}
{{- if .HashKnownHosts }}HashKnownHosts {{ .HashKnownHosts }}
{{ end }}
{{- if .HostbasedAcceptedAlgorithms }}HostbasedAcceptedAlgorithms {{ .HostbasedAcceptedAlgorithms }}
{{ end }}
{{- if .HostbasedAuthentication }}HostbasedAuthentication {{ .HostbasedAuthentication }}
{{ end }}
{{- if .HostKeyAlgorithms }}HostKeyAlgorithms {{ .HostKeyAlgorithms }}
{{ end }}
{{- if .HostKeyAlias }}HostKeyAlias {{ .HostKeyAlias }}
{{ end }}
{{- if .Hostname }}Hostname {{ .Hostname }}
{{ end }}
{{- if .IdentitiesOnly }}IdentitiesOnly {{ .IdentitiesOnly }}
{{ end }}
{{- if .IdentityAgent }}IdentityAgent {{ .IdentityAgent }}
{{ end }}
{{- if .IdentityFile }}IdentityFile {{ .IdentityFile }}
{{ end }}
{{- if .IPQoS }}IPQoS {{ .IPQoS }}
{{ end }}
{{- if .KbdInteractiveDevices }}KbdInteractiveDevices {{ .KbdInteractiveDevices }}
{{ end }}
{{- if .KexAlgorithms }}KexAlgorithms {{ .KexAlgorithms }}
{{ end }}
{{- if .KnownHostsCommand }}KnownHostsCommand {{ .KnownHostsCommand }}
{{ end }}
{{- if .LocalCommand }}LocalCommand {{ .LocalCommand }}
{{ end }}
{{- if .LocalForward }}LocalForward {{ .LocalForward }}
{{ end }}
{{- if .LogLevel -}}LogLevel {{ .LogLevel }}
{{ end }}
{{- if .MACs }}MACs {{ .MACs }}
{{ end }}
{{- if .NoHostAuthenticationForLocalhost }}NoHostAuthenticationForLocalhost {{ .NoHostAuthenticationForLocalhost }}
{{ end }}
{{- if .NumberOfPasswordPrompts }}NumberOfPasswordPrompts {{ .NumberOfPasswordPrompts }}
{{ end }}
{{- if .PasswordAuthentication }}PasswordAuthentication {{ .PasswordAuthentication }}
{{ end }}
{{- if .PermitLocalCommand }}PermitLocalCommand {{ .PermitLocalCommand }}
{{ end }}
{{- if .PermitRemoteOpen }}PermitRemoteOpen {{ .PermitRemoteOpen }}
{{ end }}
{{- if .PKCS11Provider }}PKCS11Provider {{ .PKCS11Provider }}
{{ end }}
{{- if .Port }}Port {{ .Port }}
{{ end }}
{{- if .PreferredAuthentications }}PreferredAuthentications {{ .PreferredAuthentications }}
{{ end }}
{{- if .ProxyCommand }}ProxyCommand {{ .ProxyCommand }}
{{ end }}
{{- if .ProxyJump }}ProxyJump {{ .ProxyJump }}
{{ end }}
{{- if .ProxyUseFdpass }}ProxyUseFdpass {{ .ProxyUseFdpass }}
{{ end }}
{{- if .PubkeyAcceptedAlgorithms }}PubkeyAcceptedAlgorithms {{ .PubkeyAcceptedAlgorithms }}
{{ end }}
{{- if .PubkeyAuthentication }}PubkeyAuthentication {{ .PubkeyAuthentication }}
{{ end }}
{{- if .RekeyLimit }}RekeyLimit {{ .RekeyLimit }}
{{ end }}
{{- if .RemoteCommand }}RemoteCommand {{ .RemoteCommand }}
{{ end }}
{{- if .RemoteForward }}RemoteForward {{ .RemoteForward }}
{{ end }}
{{- if .RequestTTY }}RequestTTY {{ .RequestTTY }}
{{ end }}
{{- if .SendEnv }}SendEnv {{ .SendEnv }}
{{ end }}
{{- if .ServerAliveInterval }}ServerAliveInterval {{ .ServerAliveInterval }}
{{ end }}
{{- if .ServerAliveCountMax }}ServerAliveCountMax {{ .ServerAliveCountMax }}
{{ end }}
{{- if .SessionType }}SessionType {{ .SessionType }}
{{ end }}
{{- if .SetEnv }}SetEnv {{ .SetEnv }}
{{ end }}
{{- if .StdinNull }}StdinNull {{ .StdinNull }}
{{ end }}
{{- if .StreamLocalBindMask }}StreamLocalBindMask {{.StreamLocalBindMask}}
{{ end }}
{{- if .StreamLocalBindUnlink }}StreamLocalBindUnlink {{.StreamLocalBindUnlink}}
{{ end }}
{{- if .StrictHostKeyChecking }}StrictHostKeyChecking {{.StrictHostKeyChecking}}
{{ end }}
{{- if .TCPKeepAlive }}TCPKeepAlive {{.TCPKeepAlive}}
{{ end }}
{{- if .Tunnel }}Tunnel {{.Tunnel}}
{{ end }}
{{- if .TunnelDevice }}TunnelDevice {{.TunnelDevice}}
{{ end }}
{{- if .UpdateHostKeys }}UpdateHostKeys {{.UpdateHostKeys}}
{{ end }}
{{- if .UseKeychain }}UseKeychain {{.UseKeychain}}
{{ end }}
{{- if .User }}User {{.User}}
{{ end }}
{{- if .UserKnownHostsFile }}UserKnownHostsFile {{.UserKnownHostsFile}}
{{ end }}
{{- if .VerifyHostKeyDNS }}VerifyHostKeyDNS {{.VerifyHostKeyDNS}}
{{ end }}
{{- if .VisualHostKey }}VisualHostKey {{.VisualHostKey}}
{{ end }}
{{- if .XAuthLocation }}XAuthLocation {{.XAuthLocation}}
{{ end }}
{{- end -}}