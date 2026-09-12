package config

// 
type ParamType string
const (
	ParamTypeInt    ParamType = "int"
	ParamTypeString ParamType = "string"
	ParamTypePath   ParamType = "path"
	ParamTypeEnum   ParamType = "enum"
	ParamTypeBool   ParamType = "bool"
)

// 
type ReturnType string
const (
	ReturnTypeText ReturnType = "text"
	ReturnTypeJSON ReturnType = "json"
	ReturnTypeInt  ReturnType = "int"
)

// 
type StreamTransport string
const (
	TransportMJPEG     StreamTransport = "mjpeg"
	TransportWebSocket StreamTransport = "websocket"
)

// 
type Param struct {
	Name        string    `yaml:"name"`
	Type        ParamType `yaml:"type"`
	Description string    `yaml:"description"`

	Min        *int     `yaml:"min,omitempty"`
	Max        *int     `yaml:"max,omitempty"`
	EnumValues []string `yaml:"enum_values,omitempty"`
	PathBase   string   `yaml:"path_base,omitempty"`
	MaxLength  *int     `yaml:"max_length,omitempty"`
}

// 
type Command struct {
	Name          string   `yaml:"name"`
	Exec          string   `yaml:"exec"`
	Args          []string `yaml:"args"`
	Params        []Param  `yaml:"params,omitempty"`
	Description   string   `yaml:"description"`
	Dangerous     bool     `yaml:"dangerous"`
	Reversible    bool     `yaml:"reversible"`
	SideEffects   string   `yaml:"side_effects,omitempty"`
}

// 
type GetCommand struct {
	Name               string     `yaml:"name"`
	Exec               string     `yaml:"exec"`
	Args               []string   `yaml:"args"`
	Params             []Param    `yaml:"params,omitempty"`
	Returns            ReturnType `yaml:"returns"`
	Description        string     `yaml:"description"`
}

// 
type StreamField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Unit string `yaml:"unit,omitempty"`
}

// 
type Stream struct {
	Name      string          `yaml:"name"`
	Transport StreamTransport `yaml:"transport"`
	Source    string          `yaml:"source"`

	FPS        int    `yaml:"fps,omitempty"`
	Resolution string `yaml:"resolution,omitempty"`

	RateHz     float64       `yaml:"rate_hz,omitempty"`
	Fields     []StreamField `yaml:"fields,omitempty"`
	BufferSize int           `yaml:"buffer_size,omitempty"`
}

// 
type Meta struct {
	DeviceID  string `yaml:"device_id"`
	AuthToken string `yaml:"auth_token"`
	Room      string `yaml:"room,omitempty"`
}

// 
type DeviceConfig struct {
	Meta     Meta         `yaml:"meta"`
	Commands []Command    `yaml:"commands"`
	Gets     []GetCommand `yaml:"gets"`
	Streams  []Stream     `yaml:"streams"`
}