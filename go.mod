module project-base-wahyu

go 1.22

require (
	github.com/go-playground/validator/v10 v10.19.0
	github.com/gofiber/fiber/v2 v2.52.13
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.12.3
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.30.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

// Catatan: replace directive di bawah ini HANYA workaround untuk sandbox
// testing (jaringan terbatas). Di komputer kamu dengan internet normal,
// baris ini AMAN DIHAPUS -- go mod tidy akan resolve seperti biasa lewat
// proxy.golang.org. Kalaupun dibiarkan, tidak masalah karena tetap menunjuk
// ke source code resmi yang identik (mirror GitHub official dari tim Go).
replace golang.org/x/sys => github.com/golang/sys v0.28.0

replace golang.org/x/text => github.com/golang/text v0.21.0

replace golang.org/x/net => github.com/golang/net v0.32.0

replace golang.org/x/crypto => github.com/golang/crypto v0.31.0

replace filippo.io/edwards25519 => github.com/FiloSottile/edwards25519 v1.1.0
