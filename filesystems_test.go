package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDfOutputToFileSystem(t *testing.T) {

	dfOutput := []byte(`Filesystem     1024-blocks      Used Available Capacity Mounted on
devtmpfs              4096MB         0MB      4096MB       0% /dev
tmpfs              8098376MB      5328MB   8093048MB       1% /dev/shm
/dev/nvme0n1p3   236815184MB 207116448MB  17596336MB      93% /
/dev/sdb1        229639500MB 193740116MB  24161528MB      89% /mnt/afd4af7a-2809-4c07-9152-1d5c3933fcd3
/dev/nvme0n1p1      523248MB     48204MB    475044MB      10% /boot/efi`)

	fileSystems, err := parseDfOutputToFileSystem(dfOutput)

	assert.NoError(t, err)
	assert.Len(t, fileSystems, 5)
	assert.Equal(t, fileSystems[0].Name, "devtmpfs")
	assert.Equal(t, fileSystems[4].Name, "/dev/nvme0n1p1")
	assert.Equal(t, fileSystems[0].Used, uint64(0))
	assert.Equal(t, fileSystems[0].Capacity, uint64(4096))
	assert.Equal(t, fileSystems[0].Available, uint64(4096))
	assert.Equal(t, fileSystems[0].UsedPercent, uint(0))
	assert.Equal(t, fileSystems[3].Used, uint64(193740116))
	assert.Equal(t, fileSystems[3].Capacity, uint64(217901644))
	assert.Equal(t, fileSystems[3].Available, uint64(24161528))
	assert.Equal(t, fileSystems[3].UsedPercent, uint(89))
}
