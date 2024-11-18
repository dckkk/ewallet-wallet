package constants

const (
	SuccessMessage      = "success"
	ErrFailedBadRequest = "data tidak sesuai"
	ErrServerError      = "terjadi kesalahan pada server"
)

var MappingClient = map[string]string{
	"fastcampus_ecommerce": "ini_secret_key",
}
