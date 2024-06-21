package rootfs

var Profiles = []string{"", "oci"}

type profile struct {
	BaseImage       string
	BootstrapScript string
	Size            int
}

var profiles map[string]profile = map[string]profile{
	"oci": {BaseImage: "cjrash/nex:debian", BootstrapScript: oci_profile, Size: 250000000},
}
