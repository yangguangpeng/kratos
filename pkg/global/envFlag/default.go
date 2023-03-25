package envFlag

type Handler struct {
	env string
}

const (
	EnvDev     = "dev"
	EnvUniTest = "unit-test"
	EnvFat     = "fat"
	EnvPre     = "pre"
	EnvPro     = "pro"
)

var Instance *Handler

func HandlerInstance(env string) *Handler {

	instance := &Handler{env}
	return instance
}

func (t *Handler) SetMode2UnitTest() {
	t.env = EnvUniTest
}

func (t *Handler) IsEnvPro() bool {
	return t.env == EnvPro
}

func (t *Handler) IsUnitTestMode() bool {
	return t.env == EnvUniTest
}

func (t *Handler) IsEnvMode() bool {
	return t.env == EnvDev
}
