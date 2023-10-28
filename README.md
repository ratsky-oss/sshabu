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

- Openssh

# Quick start

<!-- TODO -->

# Installation

## Usage
<!-- TODO -->

## Configuration

All unmentioned opions will be inherited from parent(s) groups 'till top group in <Group> derictive

> Config path 

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

```
Name: <string>(must be unique)    #alias for group options 
<Option>
```
### \<Group>
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
GlobalOptions:
  LogLevel: VERBOSE
  ForwardAgent: true
  PasswordAuthentication: yes

Groups:
  - Name: group1
    Options:
      User: user1
      IdentityFile: /path/to/key1
    Subgroups:
     - Name: host1
       Hosts:
          - Name: tetetete
    Hosts:
      - Name: host1
        HostName: host1.example.com

  - Name: group2
    Subgroups:
      - Name: subgroup1
        Hosts:
          - Name: host3
            HostName: host3.example.com
            User: user3
            IdentityFile: /path/to/key3
            Port: 2222

Hosts:
  - Name: host2
    HostName: host2.example.com
    User: user2
    IdentityFile: /path/to/key2
    Port: 2222
```

# License

Copyright Disclaimer:

Ratsky Walnut is the application developed by Shvora Nikita and Livitsky Andrey, licensed under the GNU General Public License version 3.0 (GPL 3.0).

All source code, design, and other intellectual property rights of Ratsky Walnut, including but not limited to text, graphics, logos, images, and software, are the property of the authors and contributors of the respective open source projects, and are protected by international copyright laws.

The information provided by Ratsky Walnut is for general informational purposes only and cannot be reproduced, distributed, or transmitted in any form or by any means without the prior written permission of the authors and contributors of the respective open source projects. Unauthorized use of any content on this app is strictly prohibited and may result in legal action.

Ratsky Walnut is licensed under the GNU General Public License version 3.0 (GPL 3.0), which provides users with the freedom to run, copy, distribute, study, change, and improve the software. However, the authors and contributors of the respective open source projects make no representations or warranties of any kind, express or imply, the completeness, accuracy, reliability, suitability, or availability of the app or the information, products, services or related graphics contained in the app for any purpose. Any reliance you place on such information is therefore strictly at your own risk.

In no event will the authors and contributors of the respective open source projects be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the app or the use or other dealings in the app.

For more information on the GPL 3.0 license, consult the [LICENSE file](LICENSE) included with the app.

All third party license agreements could be found in the third_party_licenses directory.

All rights reserved. 

# Contact

If you have any questions or feedback about the Ratsky Walnut, please contact us at 28xxgs3im@mozmail.com. We would be pleased to receive your feedback!