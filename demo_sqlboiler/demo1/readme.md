# Note: You must run this outside of your Go module directory. This must be done
# in GOPATH mode to get the correct result. If you'd like to pin the version
# manually via Go modules you can attempt other installation instructions.
go get -u -t github.com/volatiletech/sqlboiler

# Also install the driver of your choice, there exists psql, mysql, mssql
# These are separate binaries.
go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

run 
go generate
