package spec

// List of apps.
func List() []App {
	return []App{
		Docker(),
		AWS(),
		Alacritty(),
		Azure(),
		Bartender(),
		Bash(),
		Bat(),
		Zsh(),
		Wget(),
		Git(),
		GnuPG(),
		Homebrew(),
		Hyper(),
		Kubectl(),
	}
}
