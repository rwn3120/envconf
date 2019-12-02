# ENVCONF

Simple library for loading configuration from 
- yaml
- environment

Target purpose is a demo with topic `Reflection in Golang`

## Usage
See `envconf_test.go`:
```
type testConf struct {
	PtrStringParameter *string `env:"PTR_STRING_PARAMETER"`
	StringParameter    string  `env:"STRING_PARAMETER"`
    ...
}
```

### Load configuration from environment
```
os.Setenv("PTR_STRING_PARAMETER", someString)
os.Setenv("STRING_PARAMETER", someString)

c := testConf{}
if err := FromEnv(&c); err != nil {
    return err
}
```


### Load configuration from YAML
```
c := testConf{}
if err := FromYAML(os.Open(yaml), &c); err != nil {
    return err
}
```

### Load configuration from YAML file and override it by variables from environment
```
os.Setenv("STRING_PARAMETER", "my string")

c := testConf{}
if err := Load(strings.NewReader(yaml), &c); err != nil {
    return err
}
```

### Dump configuration (YAML)
```
os.Setenv("STRING_PARAMETER", "my string")

c := testConf{}
...
println(ToYaml(c))
```