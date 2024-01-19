# Usage Guide
DMux uses a custom framework model where the core Go code focuses on system functionality,
whereas the `plugin` system offers a wide variety of sub-commands that utilize the core framework.

### Keystore
The `keystore` command that allows users to interact with the system keyring that DMux
uses to house credentials. You can use `keystore` to `--reset` existing credentials
or to `--view` one of your credentials.

```
Arguments:

  -h  --help   Print help information
      --reset  Reset keyring creds: [neteng-awx,nautobot,librenms,radius]
      --view   Reset keyring creds: [neteng-awx,nautobot,librenms,radius]
```

**Command Usage**

* `dmux keystore --view nautobot`
* `dmux keystore --reset nautobot`

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

* `dmux info <device_id>`
* `dmux info <device_name>`
* `dmux info <device_ip_address>`
* `dmux info --serial <serial_number>`


### SSH
The `ssh` command provides an easy helper when connecting to network devices via ssh.
Since DMux handles credentials via the system keyring service, it can also handle 
setting up your ssh session and logging into the device. 

```
Arguments:

  positional-arg              One of DeviceName, IPAddress, DeviceID
  -h  --help                  Print help information
  -p  --port                  SSH port to use. Default: 22. Default: 22
```

**Command Usage**

* `dmux ssh <device_id>`
* `dmux ssh <device_name>`
* `dmux ssh <device_ip_address>`
* `dmux ssh <device_name> --port 2022`

### Version
The `version` command is core functionality to DMux and is used to display details
about the DMux installation.

**Command Usage**

* `dmux version`

```
DMux Version: v0.2.1
 ° Runtime: darwin_amd64
 ° Version Checked At: 2022-09-02 13:54:20 -0500 CDT
 ° Next Version Check At: 2022-09-03 13:54:20 -0500 CDT
```
