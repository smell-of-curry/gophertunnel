module github.com/sandertv/gophertunnel

go 1.22

toolchain go1.22.1

require (
	github.com/go-gl/mathgl v1.1.0
	github.com/go-jose/go-jose/v3 v3.0.3
	github.com/golang/snappy v0.0.4
	github.com/google/uuid v1.6.0
	github.com/klauspost/compress v1.17.9
	github.com/muhammadmuzzammil1998/jsonc v1.0.0
	github.com/pelletier/go-toml v1.9.5
	github.com/sandertv/go-raknet v1.14.1
	golang.org/x/net v0.28.0
	golang.org/x/oauth2 v0.22.0
	golang.org/x/text v0.17.0
)

require golang.org/x/sys v0.24.0 // indirect

require (
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/image v0.19.0 // indirect
)

replace github.com/sandertv/go-raknet => github.com/smell-of-curry/go-raknet v0.0.0-20240910211656-1f6850ca73b1
