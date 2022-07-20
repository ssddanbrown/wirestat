

Filesystem data is in MB
Memory is in MB

## Rule Format

cpu.all_active_percent >= 50 : CPU should not go over 50%

## TODO

- SystemD unit writer
  - (Takes the binary path and uses that in config)
  - Advise downloading/moving to /usr/local/bin

## Dodgy Practice List

- Gets disk info by parsing `df` command output.