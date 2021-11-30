package config

import (
	"github.com/spf13/pflag"
	"github.com/vesoft-inc/nebula-br/pkg/storage"
)

const (
	FlagStorage  = "storage"
	FlagMetaAddr = "meta"
	FlagSpaces   = "spaces"

	FlagLogPath = "log"

	flagBackupName = "name"
)

func AddCommonFlags(flags *pflag.FlagSet) {
	flags.String(FlagLogPath, "br.log", "Specify br detail log path")
	storage.AddFlags(flags)
}

type NodeInfo struct {
	Addrs   string
	RootDir string
	DataDir []string
}
