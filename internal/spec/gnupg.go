package spec

// GnuPG is the configuration for the GnuPG encryption software.
func GnuPG() App {
	return App{
		Name: "gnupg",
		Files: Files{
			".gnupg/gpg.conf",
			".gnupg/gpg-agent.conf",
			".gnupg/trustdb.gpg",
		},
	}
}
