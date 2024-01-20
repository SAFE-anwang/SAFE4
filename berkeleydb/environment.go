package berkeleydb

/*
#cgo windows CFLAGS: -I./libdb/windows
#cgo windows LDFLAGS: -L./libdb/windows -ldb
#cgo linux CFLAGS: -I./libdb/linux
#cgo linux LDFLAGS: -L./libdb/linux -ldb

#include <db.h>

int go_env_open(DB_ENV *env, char *dirname, u_int32_t flags, u_int32_t mode) {
	return env->open(env, dirname, flags, mode);
}

int go_env_close(DB_ENV *env, u_int32_t flags) {
	return env->close(env, flags);
}
*/
import "C"

// Environment holds the database environment
type Environment struct {
	environ *C.DB_ENV
}

// NewEnvironment creates a new database environment
func NewEnvironment() (*Environment, error) {
	var env *C.DB_ENV
	err := C.db_env_create(&env, 0)
	if err > 0 {
		return nil, createError(err)
	}

	return &Environment{env}, nil
}

// Open a database in the environment
func (env *Environment) Open(path string, flags C.u_int32_t, fileMode int) error {
	mode := C.u_int32_t(fileMode)

	err := C.go_env_open(env.environ, C.CString(path), flags, mode)
	return createError(err)
}

// Close a database in the environment
func (env *Environment) Close() error {
	err := C.go_env_close(env.environ, 0)
	return createError(err)
}
