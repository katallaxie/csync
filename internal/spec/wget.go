package spec

// Wget is the configuration for the wget downloader.
func Wget() App {
	return App{
		Name: "wget",
		Files: Files{
			".wgetrc",
			".wget-hsts",
		},
	}
}
