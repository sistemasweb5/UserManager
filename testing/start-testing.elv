#! /usr/bin/env nix-shell
#! nix-shell -i elvish -p go

var home = "/home/zrock/Documents/web-development/no-name/testing"
var log_file = (mktemp)
var tmp_log = (mktemp)

go version
while $true {
    printf "Log file at: %s\n" $log_file
    echo "Press enter to run, Ctrl-C+Enter to exit"

    read-line
    clear
    printf "\n\n== Execution %s ==\n" (date) >> $log_file
    try {
        go run $home > $tmp_log 2>$tmp_log
        cat $tmp_log >> $log_file
        cat $tmp_log
    } catch e {
        put $e
        tail -n 20 $tmp_log
    }
}
