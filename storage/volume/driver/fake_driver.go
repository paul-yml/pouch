package driver

import (
	"path"

	"github.com/alibaba/pouch/storage/volume/types"
)

// FakeDriver is a fake volume driver.
type FakeDriver struct {
	name string
}

// NewFakeDriver returns a fake voluem driver.
func NewFakeDriver(name string) Driver {
	return &FakeDriver{
		name: name,
	}
}

// Name returns the fake driver's name.
func (f *FakeDriver) Name(ctx Context) string {
	return f.name
}

// StoreMode returns the fake driver's store model.
func (f *FakeDriver) StoreMode(ctx Context) VolumeStoreMode {
	return UseLocalMetaStore | LocalStore
}

// Create a fake volume
func (f *FakeDriver) Create(ctx Context, id types.VolumeID) (*types.Volume, error) {
	// generate the mountPath
	mountPath := path.Join("/fake", id.Name)

	return types.NewVolumeFromID(mountPath, "", id), nil
}

// Remove a fake volume
func (f *FakeDriver) Remove(ctx Context, volume *types.Volume) error {
	return nil
}

// Path returns fake volume's path.
func (f *FakeDriver) Path(ctx Context, volume *types.Volume) (string, error) {
	return path.Join("/fake", volume.Name), nil
}
