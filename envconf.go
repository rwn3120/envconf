package envconf

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v2"
)

const envTag = "env"

func fromEnv(val reflect.Value) error {
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		// fmt.Printf("Field Name: %s,\t Field Value: %v,\t, Field type: %v\t, Field kind: %v, Field tag:%s\n",
		// 	typeField.Name, valueField.Interface(), typeField.Type, valueField.Kind(), typeField.Tag.Get(envTag))

		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				valueField.Set(reflect.New(typeField.Type.Elem()))
			}
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.Struct {
			if err := fromEnv(valueField); err != nil {
				return err
			}
		} else {
			if tagValue, found := typeField.Tag.Lookup(envTag); found {
				if envValue, found := os.LookupEnv(tagValue); found {
					switch valueField.Kind() {
					case reflect.String:
						valueField.SetString(envValue)
					case reflect.Int:
						value, err := strconv.Atoi(envValue)
						if err != nil {
							return err
						}
						valueField.SetInt(int64(value))
					case reflect.Int32:
						value, err := strconv.ParseInt(envValue, 10, 32)
						if err != nil {
							return err
						}
						valueField.SetInt(value)
					case reflect.Int64:
						value, err := strconv.ParseInt(envValue, 10, 64)
						if err != nil {
							return err
						}
						valueField.SetInt(value)
					case reflect.Uint:
						value, err := strconv.ParseUint(envValue, 10, 64)
						if err != nil {
							return err
						}
						valueField.SetUint(value)
					case reflect.Uint32:
						value, err := strconv.ParseUint(envValue, 10, 32)
						if err != nil {
							return err
						}
						valueField.SetUint(value)
					case reflect.Uint64:
						value, err := strconv.ParseUint(envValue, 10, 64)
						if err != nil {
							return err
						}
						valueField.SetUint(value)
					case reflect.Float32:
						value, err := strconv.ParseFloat(envValue, 32)
						if err != nil {
							return err
						}
						valueField.SetFloat(value)
					case reflect.Float64:
						value, err := strconv.ParseFloat(envValue, 64)
						if err != nil {
							return err
						}
						valueField.SetFloat(value)
					case reflect.Bool:
						value, err := strconv.ParseBool(envValue)
						if err != nil {
							return err
						}
						valueField.SetBool(value)
					default:
						return fmt.Errorf("Kind %v not supported", valueField.Kind())
					}
				} else {
					// fmt.Println("ENV", tagValue, "not set")
				}
			} else {
				// fmt.Println("Annotation", envTag, "not set")
			}
		}
	}
	return nil
}

// FromEnv loads configuration from envinronment
func FromEnv(configuration interface{}) error {
	if configuration != nil {
		return fromEnv(reflect.ValueOf(configuration))
	}
	return nil
}

// FromYAML loads configuration from YAML file
func FromYAML(reader io.Reader, configuration interface{}) error {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, configuration)
}

// Load loads configuration from file and overries it by variables from environment
func Load(reader io.Reader, configuration interface{}) error {
	if err := FromYAML(reader, configuration); err != nil {
		return err
	}
	return FromEnv(configuration)
}

// ToYAML serializes configuration to YAML
func ToYAML(configuration interface{}) string {
	content, err := yaml.Marshal(configuration)
	if err != nil {
		return err.Error()
	}
	return string(content)
}
