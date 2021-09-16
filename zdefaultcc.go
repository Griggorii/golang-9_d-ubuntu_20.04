package cfg

func DefaultGCCGO(goos, goarch string) string { return "/usr/bin/x86_64-linux-gnu-gccgo-9" }
func DefaultCC(goos, goarch string) string { return "x86_64-linux-gnu-gcc-9" }
func DefaultCXX(goos, goarch string) string { return "x86_64-linux-gnu-g++-9" }
const DefaultPkgConfig = "pkg-config"
var OSArchSupportsCgo = map[string]bool{}
