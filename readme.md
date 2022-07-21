# wirestat

- Filesystem data is in MB
- Memory is in MB
- Warning about low golang skill
- Default port = 8930, might need opening on firewall
- Default rules file = /etc/wirestat/rules.txt
- Update instructions
- Command to download from GitHub release
- Build & attach on release
- Quickstart script with the below commands via a curl | bash script?

- Tested on:
  - Ubuntu 20.04
  - Fedora 36

## Install

```bash
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

## Uninstall

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

## Possible Issues

##### SELinux preventing run via Systemd

You may need to reset the SELinux security context of the `/usr/local/bin/wirestat` binary after moving on systems with SELinux:

```bash
restorecon /usr/local/bin/wirestat
```

## Rule Format

cpu.all_active_percent >= 50 : CPU should not go over 50%

## Dodgy Practice List

- Gets disk info by parsing `df` command output.

## Low Maintenance Project

This is a low maintenance project. The scope of features and support are purposefully kept narrow for my purposes to ensure longer term maintenance is viable. I'm not looking to grow this into a bigger project at all.

Issues and PRs raised for bugs are perfectly fine assuming they don't significantly increase the scope of the project. Please don't open PRs for new features that expand the scope.