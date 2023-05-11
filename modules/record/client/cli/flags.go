// nolint
package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagURI    = "uri"
	FlagMeta   = "meta"
	Encrypt    = "encrypt"
	Decrypt    = "decrypt"
	From       = "from"
	Encryption = "encryption"
)

// common flag sets to add to various functions
var (
	FsCreateRecord = flag.NewFlagSet("", flag.ContinueOnError)
	FsQureyRecord  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateRecord.String(FlagURI, "", "Source URI of the record, such as an IPFS link")
	FsCreateRecord.String(FlagMeta, "", "Metadata of the record")
	FsCreateRecord.Bool(Encrypt, false, "Encrypt of the record")
	FsCreateRecord.String(Encryption, "", "encryption")
	FsQureyRecord.Bool(Decrypt, false, "Decrypt of the record")
	FsQureyRecord.String(From, "", "From")
	FsQureyRecord.String(Encryption, "", "encryption")
}
