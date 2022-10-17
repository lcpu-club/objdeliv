// Package drivers includes several storage drivers
// for objdeliv
package drivers

import (
	_ "github.com/lcpu-club/objdeliv/storage/drivers/local"
	_ "github.com/lcpu-club/objdeliv/storage/drivers/memory"
)
