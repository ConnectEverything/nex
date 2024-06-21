package rootfs

const (
	copy_fs = `#!/bin/sh
set -xe

ls / | grep -v dev | grep -v proc | grep -v sys | grep -v tmp | while read d; do tar c "/$d" | tar x -C /tmp/rootfs; done
for dir in dev proc sys tmp; do mkdir /tmp/rootfs/${dir}; done

chmod 1777 /tmp/rootfs/tmp
mkdir -p /tmp/rootfs/home/nex/
chown 1000:1000 /tmp/rootfs/home/nex/`

	oci_profile = `#!/bin/sh
set -xe

echo "installing profile dependencies"
apt-get install -y ca-certificates dbus-user-session libseccomp2 curl

get_arch() {
	a=$(uname -m)
	case ${a} in
	"x86_64" | "amd64")
		echo "amd64"
		;;
	"aarch64" | "arm64" | "arm")
		echo "arm64"
		;;
	*)
		echo ${NIL}
		;;
	esac
}

arch=$(get_arch)
runc_uri="https://github.com/opencontainers/runc/releases/latest/download/runc.${arch}"
curl -O --output-dir /usr/local/bin ${runc_uri}
`
)
