//go:build syncer
// +build syncer

package opensourcegocli

// TODO: Auto generate this from the config.yaml file
import (
	_ "github.com/cresta/public-sync-modules/autoapprove"
	_ "github.com/cresta/public-sync-modules/buildgolib"
	_ "github.com/cresta/public-sync-modules/golangcilint"
	_ "github.com/cresta/public-sync-modules/setlicense"
	_ "github.com/cresta/public-sync-modules/synceractions"
)