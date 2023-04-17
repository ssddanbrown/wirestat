# wirestat

This is a simple application that, when running on a machine, will output core system metrics in a JSON response at a HTTP endpoint. Rules can be defined against these metrics for alerting. The existence of any alerts will cause the JSON response to be served with a HTTP error status code. 

The application does not outwardly send any notifications (email, sms etc..) itself but it can be paired with a low cost, [or even self-hosted](https://github.com/louislam/uptime-kuma), website monitoring tool to form a simple system monitoring & alerting setup.

## Example JSON Response

A typical JSON response from the application, with active rules and alerts, looks like this:

```json
{
  "alerts": [
    "Root drive is almost full"
  ],
  "rules": {
    "CPU is over 50%": {
      "property": "cpu.all_active_percent",
      "operator": "\u003e=",
      "value": 50
    },
    "Memory is over 70%": {
      "property": "memory.used_percent",
      "operator": "\u003e",
      "value": 70
    },
    "Root drive is almost full": {
      "property": "filesystem./dev/nvme0n1p3.used_percent",
      "operator": "\u003e",
      "value": 80
    }
  },
  "metrics": {
    "cpu.all_active_percent": 5,
    "cpu.cpu0_active_percent": 2,
    "cpu.cpu1_active_percent": 5,
    "cpu.cpu2_active_percent": 3,
    "cpu.cpu3_active_percent": 8,
    "filesystem./dev/nvme0n1p3.available": 10955,
    "filesystem./dev/nvme0n1p3.capacity": 230107,
    "filesystem./dev/nvme0n1p3.used": 219152,
    "filesystem./dev/nvme0n1p3.used_percent": 96,
    "memory.available": 9550,
    "memory.buffers": 442,
    "memory.cached": 4047,
    "memory.free": 5333,
    "memory.swap_cached": 0,
    "memory.swap_free": 10589,
    "memory.swap_total": 10589,
    "memory.swap_used": 0,
    "memory.swap_used_percent": 0,
    "memory.total": 16585,
    "memory.used": 6760,
    "memory.used_percent": 40,
    "uptime.days": 12,
    "uptime.hours": 2,
    "uptime.minutes": 9,
    "uptime.seconds": 4
  },
  "metrics_updated_at": "2022-07-22T12:44:56.39827463+01:00"
}
```

## Advisory & Compatibility

This application has been thrown together, primarily for my own use, with little serious experience of golang, and likely contains bugs. For some stats, such as disk info, command output is read (`df command`) and parsed. It has been tested on the following systems:

- Ubuntu 20.04 (x86_64)
- Fedora 36 (x86_64)

Only Linux x86_64 systems are supported at this time. 
I would be willing to support other linux architectures, upon request & PR & testing from others, but not other operating systems (Windows, MacOS) nor any other init systems or installation setups.

## Install

Listed below are the commands that can get wirestat set-up as a Systemd service on a modern amd64 linux system.
Commands that would typically require root permissions are prefixed with `sudo`.

```bash
# Download latest release binary
curl https://github.com/ssddanbrown/wirestat/releases/latest/download/wirestat_linux_amd64 -Lo wirestat

# Make binary executable & move binary to /usr/local/bin
chmod +x wirestat
sudo mv wirestat /usr/local/bin/wirestat

# Create your rules file
sudo mkdir /etc/wirestat
sudo touch /etc/wirestat/rules.txt

# Install as a systemd service
wirestat systemd | sudo tee /etc/systemd/system/wirestat.service
sudo systemctl enable wirestat
sudo systemctl start wirestat
```

After these commands have been ran you should be able to access the wirestat output at `http://<machine_ip>:8930`.
You will need to ensure that the wirestat port (8930 by default) is open for any required access.
You can check the wirestat service status using `sudo systemctl status wirestat`.

## Update

Assuming the above steps have been used for install, updating simply requires downloading and replacing the wirestat binary: 

```bash
# Replace binary & restart systemd service
sudo curl https://github.com/ssddanbrown/wirestat/releases/latest/download/wirestat_linux_amd64 -Lo /usr/local/bin/wirestat
sudo chmod +x /usr/local/bin/wirestat
sudo systemctl restart wirestat
```

## Uninstall

Assuming the above steps have been used for install, you can pretty much do the reverse for uninstall:

```bash
# Stop and remove systemd service
sudo systemctl stop wirestat
sudo systemctl disable wirestat
sudo rm /etc/systemd/system/wirestat.service

# Delete rules
sudo rm -r /etc/wirestat

# Delete binary
sudo rm /usr/local/bin/wirestat
```

## Defining Rules

Rules are defined in a plaintext file. By default wirestat will look for a `/etc/wirestat/rules.txt` file. See "Options" below for details on using a custom path.

The rules file should contain one rule per line, with each rule following this format:

```
<metric> <operator> <value> : <rule name/label>
```

Where:

- **metric**: Is the full property name of the metric you want to check against.
- **operator**: One of: `>`, `<`, `=`, `!=`, `<=`, `>=`.
- **value**: The numeric value you want compare the metric against against.
- **rule name/label**: A unique human friendly label to describe the rule.

Blanks lines, and lines starting with a `#`, will be ignored.

**The rule should be true in the scenario you want to alert upon.**

#### Rule File Example

```txt
cpu.all_active_percent >= 50 : CPU utilization over 50%
memory.used_percent > 80 : Memory usage over 80%
filesystem./dev/nvme0n1p3.used_percent > 90 : Main disk is almost full

# Barry advised we need to reboot this machine every 30 days
uptime.days >= 30 : System has been up for over 30 days
```

#### Applying Rule Changes

Rules are read by wirestat upon start-up. Simply restart your wirestat process to re-read the rule file (Commonly `systemctl restart wirestat`).

## Alerts

When a rule passes, its name will be added to the `alerts` output in the JSON response. When any alerts are present in this output, the JSON response will be served with a non-2xx (Typically 500) HTTP response code.

Upon rules, internal issues may also show up in alerts. For example, invalidly defined rules may be advised of within the alerts output.

## Metrics

Metrics are polled in the background at ~5 second intervals. A `metrics_updated_at` property in the response reflects the last update time for this polling.

All metrics are reflected as an unsigned int64 value. Any metric that represents a percentage value is specifically labelled as such.

#### CPU

All metrics are shown as percentages. Activity it taken from a 1 second sample from `/proc/stat`. Metrics will be shown for each thread upon a value for all threads.

#### Filesystems

These metrics are take from parsing the command `df -P -B MB`. 
All filesystems from this output will be reflected.
Non-percentage values shown are in MegaBytes.

#### Memory

Non-percentage values shown are in MegaBytes.

### Uptime

Uptime is shown broken down into a series of units. The lesser units are not the total amount of time passed for that type, but instead provide precision to the greater units. For example, `uptime.seconds` will not go above 59, it only shows the number of seconds since the last counted minute.

## Options & Arguments

wirestat can be provided some command line options. If it's setup as a Systemd service as above, you'd need to edit the command line options within the `ExecStart` line of your `/etc/systemd/system/wirestat.service` file.

#### `-port`

The port to use to run the wirestat server on. Defaults to `8930`. Should be a port that's not used by any other service.

```bash
# Example of using wirestat on port 9090
wirestat -port 9090
```

#### `-rules`

The path to the rules file to read. Defaults to `/etc/wirestat/rules.txt`.

```bash
# Example of using a custom rules file location
wirestat -rules /home/barry/wirestat-rules.txt
```

#### `-accesskey`

Define a key string to be required to access the JSON output, as a very basic form of authorization.

```bash
# Example of setting a custom access key
wirestat -accesskey "hunter2"
```

When this option is set you'll need to provide this key value, when accessing the JSON data, as either:

- A `key` query parameter value in the URL.
- A `X-Access-Key` header value.

#### `-ruledelimiter`

The delimeter to use when parsing rules. Use this if you have metrics that have : within them.

```bash
# Example of using ~ as the rule delimeter instead of :
wirestat -ruledelimeter "~"
```

#### `systemd`

This argument will tell wirestat to print systemd configuration to stdout, instead of running the application as normal. Other options can be passed with this argument, and those options will be used within the command output.

```bash
wirestat -port 9090 -rules /root/rules.txt systemd
```

## Possible Issues

##### SELinux preventing run via Systemd

You may need to reset the SELinux security context of the `/usr/local/bin/wirestat` binary after moving on systems with SELinux:

```bash
restorecon /usr/local/bin/wirestat
```

## Low Maintenance Project

This is a low maintenance project. The scope of features and support are purposefully kept narrow for my purposes to ensure longer term maintenance is viable. I'm not looking to grow this into a bigger project at all.

Issues and PRs raised for bugs are perfectly fine assuming they don't significantly increase the scope of the project. Please don't open PRs for new features that expand the scope.

## Attribution

The following great projects are directly used within this project:

- [goprocinfo](https://github.com/c9s/goprocinfo) - MIT Licensed
- [testify](github.com/stretchr/testify) - MIT Licensed