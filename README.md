# Ratsky Sshabu


# Table of contents

- [Overview](#overview)
- [Quick start](#quick-start)
- [Installation](#installation)
- [Commands](#commands)
- [Configuration](#configuration)
- [License](#license)
- [Contacts](#contact)

# Overview

`Ratsky Sshabu` is a robust SSH client management tool designed to streamline the process of connecting to multiple servers effortlessly. This tool leverages OpenSSH and offers a user-friendly interface to enhance the overall SSH experience. With Sshabu, managing SSH configurations becomes more intuitive, allowing users to organize and connect to their servers efficiently.

<img src="docs/sshabu-quick.gif" alt="alt text">

> Openssh should be installed on your system.

# Quick start

1. Download the binary file `sshabu`. You can view them on the release page:
https://github.com/Ra-sky/sshabu/releases/tag/v0.0.1-alpha
```bash
wget https://github.com/Ra-sky/sshabu/releases/download/v0.0.1-alpha/sshabu_Darwin_arm64.tar.gz
```
2. Unzip and move binary file `sshabu` to `/usr/local/bin/`.
```bash
mkdir sshabu_Darwin_arm64 && tar -xvzf sshabu_Darwin_arm64.tar.gz -C sshabu_Darwin_arm64 && \
  cd sshabu_Darwin_arm64 && \
  mv sshabu /usr/local/bin/sshabu
```
2. Initialize your sshabu configuration.
```bash
sshabu init
```
3. Enable auto-completion. Several options are available; use the following command to view them:
```bash
sshabu completion --help
```
4. Begin editing the config with this convenient command:
```bash
sshabu edit
```
5. Connect to servers by specifying the name.
```bash
sshabu connect <Name>
```
6. Yep-Yep-Yep! It's time for shabu!

# Installation
## Brew
1. Adding Ratsky third-party repository 
```bash
brew tap ratsky-oss/taps
```
2. Installing sshabu
```
brew install sshabu
```
3. Validate sshabu binary
```
which sshabu
```
4. Initialize your `sshabu` configuration.
```bash
sshabu init
```
## Easy
1. Download the binary file `sshabu` to `/usr/bin/sshabu`. You can change the default path from `/usr/bin/sshabu` to your own, but make sure that the path is included in your `PATH` environment variable.
<!-- ```bash
wget -O /usr/bin/sshabu https://git.ratsky.ru/ratsky/ssh-client/sshabu/-/jobs/2334/artifacts/download?file_type=archive
``` -->
2. Initialize your `sshabu` configuration.
```bash
sshabu init
```
## Build from source
1. Clone the git repository.
<!-- ```bash
git clone https://git.ratsky.ru/ratsky/ssh-client/sshabu.git
``` -->
2. Change the directory to the cloned project.
```bash
cd ./sshabu
```
3. Build the project.
```bash
go build .
```
4. Move the binary file `sshabu`. You can change the default path from `/usr/bin/sshabu` to your own, but make sure that the path is included in your `PATH` environment variable.
```bash
mv sshabu /usr/bin/sshabu
```
5. Initialize your `sshabu` configuration.
```bash
sshabu init
```

## Commands

#### `sshabu init`
Create `~/$HOME/.sshabu/` directory and generate example `sshabu.yaml` config
#### `sshabu apply`
Generate `openssh.config` based on `sshabu.yaml`
#### `sshabu edit`
Open `sshabu.yaml` with editor and runs `sshabu apply after that`
#### `sshabu connect`
Runs openssh command with `openssh.config`

> Find out more info by using `--help` flag

## Configuration

All unmentioned options will be inherited from parent(s) groups 'till top group in <Group> derictive

> ~/.sshabu/sshabu.yaml

Config structure

```
GlobalOptions:
    <Option>

Groups:
    [- <Group>] 

Hosts:
    [- <Host>]
```


### \<Host>

Target of ssh configuration

> All \<Host>.Name(s) must be unique

```
Name: <string>
<Option>
```
Name field is an identifire of following config. \
Name field value will be used as "Host" derective in final Openssh config, if no "Host" option in <Host> is defined.

### \<Group>

Groups allow you to define some options for all included entities such s \<Host> in "Hosts" and \<Group> in "Subgroups".

 Groups was designed to provide option inheritance

 Option precedence list (higher option will override lower):
 - \<Host>.option
 - \<Group>.\<Host>.option
 ...
 - \<Group>.\<Group>....option
 ...
 - GlobalOption

```
Name: <string>
Hosts:
    - [<Host>]
Options:
    <Option>
Subgroups:
    - [<Group>]
```

### \<Option>

Generally just openssh config options in "key: value" format\
man page - ssh_config(5)
https://linux.die.net/man/5/ssh_config \

> Avoid using 'Host' option unless you know what you are doing.  

```
HostName: <str>
IdentityFile: <str>
StrictHostKeyChecking: <yes/accept-new/no/off/ask>
...
```
> Suggestion - "Host" better be unique if used \
> Considerations: \
> Host wildcards?? \
> Using "Host" inside:
> - "GlabalOptions" is useless
> - "\<Host>" will override \<Host>."Name" section in destination openssh config

<details>
<summary>All options reference</summary>

|Openssh option  | Sshabu | Tested |
|---|:---:|:---:|
| AddKeysToAgent  | ✅ | ❌ |
| AddressFamily | ✅ | ❌ |
| BatchMode | ✅ | ❌ |
| BindAddress | ✅ | ❌ |
| CanonicalDomains  | ✅ | ❌ |
| CanonicalizeFallbackLocal | ✅ | ❌ |
| CanonicalizeHostname  | ✅ | ❌ |
| CanonicalizeMaxDots | ✅ | ❌ |
| CanonicalizePermittedCNAMEs | ✅ | ❌ |
| CASignatureAlgorithms | ✅ | ❌ |
| CertificateFile | ✅ | ❌ |
| CheckHostIP | ✅ | ❌ |
| Ciphers | ✅ | ❌ |
| ClearAllForwardings | ✅ | ❌ |
| Compression | ✅ | ❌ |
| ConnectionAttempts  | ✅ | ❌ |
| ConnectTimeout  | ✅ | ❌ |
| ControlMaster | ✅ | ❌ |
| ControlPath | ✅ | ❌ |
| ControlPersist  | ✅ | ❌ |
| DynamicForward  | ✅ | ❌ |
| EnableEscapeCommandline | ❌ | ❌ |
| EscapeChar  | ✅ | ❌ |
| ExitOnForwardFailure  | ✅ | ❌ |
| FingerprintHash | ✅ | ❌ |
| ForkAfterAuthentication | ✅ | ❌ |
| ForwardAgent  | ✅ | ❌ |
| ForwardX11  | ✅ | ❌ |
| ForwardX11Timeout | ✅ | ❌ |
| ForwardX11Trusted | ✅ | ❌ |
| GatewayPorts  | ✅ | ❌ |
| GlobalKnownHostsFile  | ✅ | ❌ |
| GSSAPIAuthentication  | ✅ | ❌ |
| GSSAPIDelegateCredentials | ✅ | ❌ |
| HashKnownHosts  | ✅ | ❌ |
| Host  | ✅ | ✅ |
| HostbasedAcceptedAlgorithms | ✅ | ❌ |
| HostbasedAuthentication | ✅ | ❌ |
| HostKeyAlgorithms | ✅ | ❌ |
| HostKeyAlias  | ✅ | ❌ |
| Hostname  | ✅ | ✅ |
| IdentitiesOnly  | ✅ | ❌ |
| IdentityAgent | ✅ | ❌ |
| IdentityFile  | ✅ | ✅ |
| IPQoS | ✅ | ❌ |
| KbdInteractiveAuthentication  | ❌ | ❌ |
| KbdInteractiveDevices | ✅ | ❌ |
| KexAlgorithms | ✅ | ❌ |
| KnownHostsCommand | ✅ | ❌ |
| LocalCommand  | ✅ | ❌ |
| LocalForward  | ✅ | ❌ |
| LogLevel  | ✅ | ❌ |
| MACs  | ✅ | ❌ |
| Match | ❌ | ❌ |
| NoHostAuthenticationForLocalhost  | ✅ | ❌ |
| NumberOfPasswordPrompts | ✅ | ❌ |
| PasswordAuthentication  | ✅ | ❌ |
| PermitLocalCommand  | ✅ | ❌ |
| PermitRemoteOpen  | ✅ | ❌ |
| PKCS11Provider  | ✅ | ❌ |
| Port  | ✅ | ✅ |
| PreferredAuthentications  | ✅ | ❌ |
| ProxyCommand  | ✅ | ❌ |
| ProxyJump | ✅ | ❌ |
| ProxyUseFdpass  | ✅ | ❌ |
| PubkeyAcceptedAlgorithms  | ✅ | ❌ |
| PubkeyAuthentication  | ✅ | ❌ |
| RekeyLimit  | ✅ | ❌ |
| RemoteCommand | ✅ | ❌ |
| RemoteForward | ✅ | ❌ |
| RequestTTY  | ✅ | ❌ |
| RequiredRSASize | ❌ | ❌ |
| SendEnv | ✅ | ❌ |
| ServerAliveInterval | ✅ | ❌ |
| ServerAliveCountMax | ✅ | ❌ |
| SessionType | ✅ | ❌ |
| SetEnv  | ✅ | ❌ |
| StdinNull | ✅ | ❌ |
| StreamLocalBindMask | ✅ | ❌ |
| StreamLocalBindUnlink | ✅ | ❌ |
| StrictHostKeyChecking | ✅ | ✅ |
| TCPKeepAlive  | ✅ | ❌ |
| Tunnel  | ✅ | ❌ |
| TunnelDevice  | ✅ | ❌ |
| UpdateHostKeys  | ✅ | ❌ |
| UseKeychain | ✅ | ❌ |
| User  | ✅ | ✅ |
| UserKnownHostsFile  | ✅ | ❌ |
| VerifyHostKeyDNS  | ✅ | ❌ |
| VisualHostKey | ✅ | ❌ |
| XAuthLocation | ✅ | ❌ |

</details>

<details>
<summary>Coniguration example</summary>

```
# ----------------------------------------------------------------------
# Default options for all hosts

GlobalOptions:
  LogLevel: INFO
  User: user
  IdentityFile: /.ssh/id_rsa

# ----------------------------------------------------------------------
# Top level standalone host list

Hosts:
  - Name: smth_ungrouped                          # Host example
    HostName: host.example.com                    # Key: value ssh_config(5) 
    User: user2
    IdentityFile: /path/to/key2
    Port: 2223

# ----------------------------------------------------------------------

# Top level group list
Groups:
  - Name: work                                     # Some group
    Hosts:                                         # List of Target host
      - Name: project1-lab                         # This host will inherit User, IdentityFile from group "work"
        HostName: lab.project1.ratsky.local
    Options:                                       # ssh_config(5)
      User: alivitskiy
      IdentityFile: /.ssh/id_rsa.work
    Subgroups:                                     # List of subgroups that will inherit "work" Options

     - Name: prod
       Options:
          IdentityFile: /.ssh/id_rsa.work2
       Hosts:
          - Name: project1
            HostName: 192.168.1.2
            Port: 2222

     - Name: test
       Hosts:
          - Name: project1-test
            HostName: 192.168.11.3

  - Name: home                                     # Another group
    Hosts:
      - Name: home-gitlab
        HostName: gitlab.ratsky.local
      - Name: home-nextcloud
        HostName: nc.ratsky.local
 
# ----------------------------------------------------------------------
```

Result
```
# sshabu apply

IdentityFile /.ssh/id_rsa
LogLevel INFO

Host smth_ungrouped
    Hostname host.example.com
    IdentityFile /path/to/key2
    Port 2223
    
Host project1-lab
    Hostname lab.project1.ratsky.local
    IdentityFile /.ssh/id_rsa.work
    
Host project1
    Hostname 192.168.1.2
    IdentityFile /.ssh/id_rsa.work2
    Port 2222
    
Host project1-test
    Hostname 192.168.11.3
    IdentityFile /.ssh/id_rsa.work
    
Host home-gitlab
    Hostname gitlab.ratsky.local
    
Host home-nextcloud
    Hostname nc.ratsky.local

```

</details>

# License

Ratsky Sshabu is released under the Apache 2.0 license. See LICENSE.

# Contact

If you have any questions or feedback about the Ratsky Sshabu, please contact us at 28xxgs3im@mozmail.com. We would be pleased to receive your feedback!