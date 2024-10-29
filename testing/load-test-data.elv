#! /usr/bin/env nix-shell
#! nix-shell -i elvish -p postgresql_16_jit

set E:PGPASSWORD = "ecliptic"

psql -h localhost -p 5100 -U phaeton -d helios -f resources/test-data.sql



