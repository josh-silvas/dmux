# Usage Guide
NBot uses a custom framework model where the core Go code focuses on system functionality,
whereas the `plugin` system offers a wide variety of sub-commands that utilize the core framework.

### Keystore
The `keystore` command that allows users to interact with the system keyring that NBot
uses to house credentials. You can use `keystore` to `--reset` existing credentials
or to `--view` one of your credentials.

```
Arguments:

  -h  --help   Print help information
      --reset  Reset keyring creds: [neteng-awx,nautobot,librenms,radius]
      --view   Reset keyring creds: [neteng-awx,nautobot,librenms,radius]
```

**Command Usage**

* `nbot keystore --view nautobot`
* `nbot keystore --reset nautobot`

### Info
The `info` command is a useful command to gather information about a device or
resource from systems. 

```
Arguments:

  positional-arg               One of DeviceName, IPAddress, DeviceID
  -h  --help                   Print help information
  -t  --tag                    Devices with a tag
  -s  --site                   Devices within a given site name.
  -m  --mac-address            Devices within a given mac-address.
      --serial                 Devices by serial number.
```

**Command Usage**

* `nbot info <device_id>`
* `nbot info <device_name>`
* `nbot info <device_ip_address>`
* `nbot info --serial <serial_number>`


### SSH
The `ssh` command provides an easy helper when connecting to network devices via ssh.
Since NBot handles credentials via the system keyring service, it can also handle 
setting up your ssh session and logging into the device. 

```
Arguments:

  positional-arg              One of DeviceName, IPAddress, DeviceID
  -h  --help                  Print help information
  -p  --port                  SSH port to use. Default: 22. Default: 22
```

**Command Usage**

* `nbot ssh <device_id>`
* `nbot ssh <device_name>`
* `nbot ssh <device_ip_address>`
* `nbot ssh <device_name> --port 2022`

### Version
The `version` command is core functionality to NBot and is used to display details
about the NBot installation.

**Command Usage**

* `nbot version`

```
NBot Version: v0.2.1
 ° Runtime: darwin_amd64
 ° Version Checked At: 2022-09-02 13:54:20 -0500 CDT
 ° Next Version Check At: 2022-09-03 13:54:20 -0500 CDT
```
