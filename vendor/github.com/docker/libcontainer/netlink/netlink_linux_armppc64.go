// +build arm ppc64 ppc64le

package netlink

func ifrDataByte(b byte) uint8 {
	return uint8(b)
}
