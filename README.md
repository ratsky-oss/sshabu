# Ratskt Sshabu

<!-- ![alt text](docs/logo_gradient_square.svg "containers"){:height="100px" width="100px"} -->
<!-- <img src="docs/logo_gradient_square.svg" alt="alt text" width="300" height="300"> -->



# Table of contents

- [Overview](#overview)
- [Requirements](#requirements)
- [Quick start](#quick-start)
- [Installation](#installation)
- [Configuration](#configuration)
    - [Settings reference](#settings-reference) 
    - [Crucial settings](#crucial-settings)
    - [Service management](#service-management)
- [Troubleshooting](#troubleshooting)
- [License](#license)
- [Contacts](#contact)

# Overview

<!-- TODO -->

# Requirements

<!-- TODO -->
<!-- TODO: Not supported directivies -->

- Openssh

# Quick start
1. Download the binary file `sshabu`. You can change the default path from `/usr/bin/sshabu` to your own, but make sure that the path is included in your `PATH` environment variable.
```bash
wget -O /usr/bin/sshabu https://git.ratsky.ru/ratsky/ssh-client/sshabu/-/jobs/2334/artifacts/download?file_type=archive
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

## Usage
<!-- TODO -->

## Configuration

All unmentioned opions will be inherited from parent(s) groups 'till top group in <Group> derictive

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

 <!-- // TODO: Host is generally the same as Name -->

```
HostName: <str>
IdentityFile: <str>
...
```
> Suggestion - "Host" better be unique if used \
> Considerations: \
> Host wildcards?? \
> Using "Host" inside:
> - "GlabalOptions" is useless
> - "\<Host>" will override \<Host>."Name" section in destination openssh config

### Configuration example

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
# License


# Contact

If you have any questions or feedback about the Ratsky Sshabu, please contact us at 28xxgs3im@mozmail.com. We would be pleased to receive your feedback!