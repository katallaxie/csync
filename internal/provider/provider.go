package provider

// Backup is the backup interface.
type Backup interface {
	// Folder returns the backup folder.
	Folder(f string) (string, error)
}
