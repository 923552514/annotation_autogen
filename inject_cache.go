package annotation_autogen

// CacheInjectMethodInfo specifies the method that to be injected all the params could be used
type CacheInjectMethodInfo struct {
	StructName string // method's struct object's name
	PkgName    string // method's struct object's pkg name
	MethodName string // method's name
}

// KeyGenerator is an interface that contains any methods:
//
//	GenKey: specifies a method could be impl, it has one output param as type string that means the cache key name;
//	 	it has 4 input params:
//			method[CacheInjectMethodInfo]: injected method's info
//			keyValue[interface]: as annotation as key field's value
//			kwargs[map[string]interface]: specifies as a map withs key as param's name and value as param's value
//			args[[]interface]: specifies as a slice contains value as param's value
type KeyGenerator interface {
	GenKey(method *CacheInjectMethodInfo, keyValue interface{}, kwargs map[string]interface{}, args ...interface{}) string
}

// CacheSourceHandler specifies an interface that could be impl
//
//	Set Method: specifies the cache set function, it contains 4 params:
//		key: cache's key name, as specifies as annotation's key field
//		value: cache's value, as specifies as annotation's value field
//		expire: cache's expire time, as specifies as annotation's expire field
type CacheSourceHandler interface {
	Set(key string, value interface{}, expire int64, kwargs map[string]interface{}) error
	Get(key string, kwargs map[string]interface{}) (interface{}, error)
	Del(key string, kwargs map[string]interface{}) error
}

/*
// CacheHandlerConfig
type CacheHandlerConfig struct {
}
*/

/*
type CacheInjectConfig struct {
	CacheSourceHandler CacheSourceHandler
	CacheHandlerConfig map[string]*CacheHandlerConfig
	KeyGenerators      map[string]KeyGenerator
}

func (*CacheInjectConfig) Inject() {

}
*/

// CacheRegister
// register cache inject
// it contains two input params:
// sourceHandler impl of CacheSourceHandler, it must not be nil
// keyGenerator type of map[string]KeyGenerator, it could be nil
func CacheRegister(sourceHandler CacheSourceHandler, keyGenerator map[string]KeyGenerator) string {
	return "implementation not generated, run cmd"
}
