package sshabu

// import (
//     "time"
// )

type Shabu struct{
    Options     Options     `mapstructure:"globaloptions,omitempty"`
    Hosts       []Host      `mapstructure:"hosts,omitempty"`
    Groups      []Group     `mapstructure:"groups,omitempty"`
}

type Host struct{
	Options  Options `mapstructure:",squash,omitempty"`
    Name     string  `mapstructure:"name"`
}

type Group struct{
	Options         Options `mapstructure:"options,omitempty"`
    Hosts           []Host  `mapstructure:"hostname,omitempty"`
    Name            string  `mapstructure:"name"`
    Subgroups       []Group `mapstructure:"subgroups"`
}

type Option interface{}
type Options struct {
    AddKeysToAgent                       Option          `mapstructure:"addkeystoagent,omitempty"`
    // *AddKeysToAgent                       *bool          `mapstructure:"AddKeysToAgent,omitempty"`
    AddressFamily                        Option        `mapstructure:"addressfamily,omitempty"`
    BatchMode                            Option          `mapstructure:"batchmode,omitempty"`
    // *BatchMode                            bool          `mapstructure:"BatchMode,omitempty"`
    BindAddress                          Option        `mapstructure:"bindaddress,omitempty"`
    CanonicalDomains                     Option        `mapstructure:"canonicaldomains,omitempty"`
    CanonicalizeFallbackLocal            Option          `mapstructure:"CanonicalizeFallbackLocal,omitempty"`
    // *CanonicalizeFallbackLocal            bool          `mapstructure:"CanonicalizeFallbackLocal,omitempty"`
    CanonicalizeHostname                 Option        `mapstructure:"canonicalizehostname,omitempty"`
    CanonicalizeMaxDots                  Option           `mapstructure:"canonicalizemaxdots,omitempty"`
    CanonicalizePermittedCNAMEs          Option        `mapstructure:"CanonicalizePermittedCNAMEs,omitempty"`
    CASignatureAlgorithms                Option        `mapstructure:"casignaturealgorithms,omitempty"`
    CertificateFile                      Option        `mapstructure:"certificatefile,omitempty"`
    CheckHostIP                          Option        `mapstructure:"checkHostip,omitempty"`
    Ciphers                              Option        `mapstructure:"ciphers,omitempty"`
    ClearAllForwardings                  Option          `mapstructure:"clearallforwardings,omitempty"`
    // *ClearAllForwardings                  bool          `mapstructure:"ClearAllForwardings,omitempty"`
    Compression                          Option        `mapstructure:"compression,omitempty"`
    ConnectionAttempts                   Option           `mapstructure:"connectionattempts,omitempty"`
    ConnectTimeout                       Option `mapstructure:"connecttimeout,omitempty"`
    ControlMaster                        Option        `mapstructure:"controlmaster,omitempty"`
    ControlPath                          Option        `mapstructure:"controlpath,omitempty"`
    ControlPersist                       Option        `mapstructure:"controlpersist,omitempty"`
    DynamicForward                       Option        `mapstructure:"dynamicforward,omitempty"`
    EscapeChar                           Option        `mapstructure:"escapechar,omitempty"`
    ExitOnForwardFailure                 Option          `mapstructure:"exitonforwardfailure,omitempty"`
    // *ExitOnForwardFailure                 bool          `mapstructure:"ExitOnForwardFailure,omitempty"`
    FingerprintHash                      Option        `mapstructure:"fingerprinthash,omitempty"`
    ForkAfterAuthentication              Option          `mapstructure:"forkafterauthentication,omitempty"`
    // *ForkAfterAuthentication              bool          `mapstructure:"ForkAfterAuthentication,omitempty"`
    ForwardAgent                         Option          `mapstructure:"forwardagent,omitempty"`
    // *ForwardAgent                         bool          `mapstructure:"ForwardAgent,omitempty"`
    ForwardX11                           Option          `mapstructure:"forwardx11,omitempty"`
    // *ForwardX11                           bool          `mapstructure:"ForwardX11,omitempty"`
    ForwardX11Timeout                    Option `mapstructure:"forwardx11timeout,omitempty"`
    ForwardX11Trusted                    Option          `mapstructure:"forwardx11trusted,omitempty"`
    // *ForwardX11Trusted                    bool          `mapstructure:"ForwardX11Trusted,omitempty"`
    GatewayPorts                         Option        `mapstructure:"gatewayports,omitempty"`
    GlobalKnownHostsFile                 Option        `mapstructure:"globalknownhostsfile,omitempty"`
    GSSAPIAuthentication                 Option          `mapstructure:"gssapiauthentication,omitempty"`
    // *GSSAPIAuthentication                 bool          `mapstructure:"GSSAPIAuthentication,omitempty"`
    GSSAPIDelegateCredentials            Option          `mapstructure:"gssapidelegatecredentials,omitempty"`
    // *GSSAPIDelegateCredentials            bool          `mapstructure:"GSSAPIDelegateCredentials,omitempty"`
    HashKnownHosts                       Option          `mapstructure:"hashknownhosts,omitempty"`
    // *HashKnownHosts                       bool          `mapstructure:"HashKnownHosts,omitempty"`
    Host                                 Option        `mapstructure:"host,omitempty"`
    HostbasedAcceptedAlgorithms          Option        `mapstructure:"hostbasedacceptedalgorithms,omitempty"`
    HostbasedAuthentication              Option          `mapstructure:"hostbasedauthentication,omitempty"`
    // *HostbasedAuthentication              bool          `mapstructure:"HostbasedAuthentication,omitempty"`
    HostKeyAlgorithms                    Option        `mapstructure:"hostkeyalgorithms,omitempty"`
    HostKeyAlias                         Option        `mapstructure:"hostkeyalias,omitempty"`
    Hostname                             Option        `mapstructure:"hostname,omitempty"`
    IdentitiesOnly                       Option          `mapstructure:"IdentitiesOnly,omitempty"`
    // *IdentitiesOnly                       bool          `mapstructure:"IdentitiesOnly,omitempty"`
    IdentityAgent                        Option        `mapstructure:"identityagent,omitempty"`
    IdentityFile                         Option        `mapstructure:"identityfile,omitempty"`
    IPQoS                                Option        `mapstructure:"ipqos,omitempty"`
    // *KbdInteractiveAuthentication         bool          `mapstructure:"KbdInteractiveAuthentication,omitempty"`
    KbdInteractiveDevices                Option        `mapstructure:"kbdinteractivedevices,omitempty"`
    KexAlgorithms                        Option        `mapstructure:"kexalgorithms,omitempty"`
    KnownHostsCommand                    Option        `mapstructure:"knownhostscommand,omitempty"`
    LocalCommand                         Option        `mapstructure:"localcommand,omitempty"`
    LocalForward                         Option        `mapstructure:"localforward,omitempty"`
    LogLevel                             Option        `mapstructure:"loglevel,omitempty"`
    MACs                                 Option        `mapstructure:"macs,omitempty"`
    Match                                Option        `mapstructure:"match,omitempty"`
    NoHostAuthenticationForLocalhost     Option          `mapstructure:"nohostauthenticationforlocalhost,omitempty"`
    // *NoHostAuthenticationForLocalhost     bool          `mapstructure:"NoHostAuthenticationForLocalhost,omitempty"`
    NumberOfPasswordPrompts              Option           `mapstructure:"numberofpasswordprompts,omitempty"`
    PasswordAuthentication               Option          `mapstructure:"passwordauthentication,omitempty"`
    // *PasswordAuthentication               bool          `mapstructure:"PasswordAuthentication,omitempty"`
    PermitLocalCommand                   Option          `mapstructure:"permitlocalcommand,omitempty"`
    // *PermitLocalCommand                   bool          `mapstructure:"PermitLocalCommand,omitempty"`
    PermitRemoteOpen                     Option          `mapstructure:"permitremoteopen,omitempty"`
    // *PermitRemoteOpen                     bool          `mapstructure:"PermitRemoteOpen,omitempty"`
    PKCS11Provider                       Option        `mapstructure:"pkcs11provider,omitempty"`
    Port                                 Option           `mapstructure:"port,omitempty"`
    PreferredAuthentications             Option        `mapstructure:"preferredauthentications,omitempty"`
    ProxyCommand                         Option        `mapstructure:"proxycommand,omitempty"`
    ProxyJump                            Option        `mapstructure:"proxyjump,omitempty"`
    ProxyUseFdpass                       Option          `mapstructure:"proxyusefdpass,omitempty"`
    // *ProxyUseFdpass                       bool          `mapstructure:"ProxyUseFdpass,omitempty"`
    PubkeyAcceptedAlgorithms             Option        `mapstructure:"pubkeyacceptedalgorithms,omitempty"`
    PubkeyAuthentication                 Option          `mapstructure:"pubkeyauthentication,omitempty"`
    // *PubkeyAuthentication                 bool          `mapstructure:"PubkeyAuthentication,omitempty"`
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
    // *StdinNull                            bool          `mapstructure:"StdinNull,omitempty"`
    StreamLocalBindMask                  Option        `mapstructure:"streamlocalbindmask,omitempty"`
    StreamLocalBindUnlink                Option          `mapstructure:"streamlocalbindunlink,omitempty"`
    // *StreamLocalBindUnlink                bool          `mapstructure:"StreamLocalBindUnlink,omitempty"`
    StrictHostKeyChecking                Option          `mapstructure:"stricthostkeychecking,omitempty"`
    TCPKeepAlive                         Option          `mapstructure:"tcpkeepalive,omitempty"`
    // *TCPKeepAlive                         bool          `mapstructure:"TCPKeepAlive,omitempty"`
    Tunnel                               Option          `mapstructure:"tunnel,omitempty"`
    // *Tunnel                               bool          `mapstructure:"Tunnel,omitempty"`
    TunnelDevice                         Option        `mapstructure:"tunneldevice,omitempty"`
    UpdateHostKeys                       Option          `mapstructure:"updatehostkeys,omitempty"`
    // *UpdateHostKeys                       bool          `mapstructure:"UpdateHostKeys,omitempty"`
    UseKeychain                          Option          `mapstructure:"usekeychain,omitempty"`
    // *UseKeychain                          bool          `mapstructure:"UseKeychain,omitempty"`
    User                                 Option        `mapstructure:"user,omitempty"`
    UserKnownHostsFile                   Option        `mapstructure:"userknownhostsfile,omitempty"`
    VerifyHostKeyDNS                     Option          `mapstructure:"verifyhostkeydns,omitempty"`
    // *VerifyHostKeyDNS                     bool          `mapstructure:"VerifyHostKeyDNS,omitempty"`
    VisualHostKey                        Option          `mapstructure:"visualhostkey,omitempty"`
    // *VisualHostKey                        bool          `mapstructure:"VisualHostKey,omitempty"`
    XAuthLocation                        Option        `mapstructure:"xauthlocation,omitempty"`
}