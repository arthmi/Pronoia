package config

import (
	"errors"
	"fmt"
)

// Validate checks the validity of the device configuration
func Validate(cfg *DeviceConfig) error {
	var errs []error

	if cfg.Meta.DeviceID == "" {
		errs = append(errs, fmt.Errorf("meta.device_id is required"))
	}

	for _, cmd := range cfg.Commands {
		errs = append(errs, validateCommand(cmd)...)
	}

	for _, get := range cfg.Gets {
		errs = append(errs, validateGet(get)...)
	}

	for _, stream := range cfg.Streams {
		errs = append(errs, validateStream(stream)...)
	}

	return errors.Join(errs...)
}

// Check if the command is valid
func validateCommand(cmd Command) []error {
	var errs []error

	if cmd.Name == "" {
		errs = append(errs, fmt.Errorf("a command has an empty name"))
	}

	if cmd.Exec == "" {
		errs = append(errs, fmt.Errorf("command %q: exec is required", cmd.Name))
	}

	// Check if the command is dangerous and requires a confirm parameter
	if cmd.Dangerous && !hasConfirmParam(cmd.Params) {
		errs = append(errs, fmt.Errorf("command %q: dangerous=true requires a bool parameter named confirm", cmd.Name))
	}

	for _, p := range cmd.Params {
		errs = append(errs, validateParam(p, cmd.Name)...)
	}

	return errs
}

// Check if the get command is valid
func validateGet(get GetCommand) []error {
	var errs []error

	if get.Name == "" {
		errs = append(errs, fmt.Errorf("a get has an empty name"))
	}

	if get.Exec == "" {
		errs = append(errs, fmt.Errorf("get %q: exec is required", get.Name))
	}

	switch get.Returns {
	case ReturnTypeText, ReturnTypeJSON, ReturnTypeInt:
		// Valid return type
	default:
		errs = append(errs, fmt.Errorf("get %q: returns %q unknown", get.Name, get.Returns))
	}

	for _, p := range get.Params {
		errs = append(errs, validateParam(p, get.Name)...)
	}

	return errs
}

// Check if the stream is valid
func validateStream(s Stream) []error {
	var errs []error

	if s.Name == "" {
		errs = append(errs, fmt.Errorf("a stream has an empty name"))
	}

	switch s.Transport {
	case TransportMJPEG:
		if s.Source == "" {
			errs = append(errs, fmt.Errorf("stream %q: source is required for mjpeg", s.Name))
		}
	case TransportWebSocket:
		if len(s.Fields) == 0 {
			errs = append(errs, fmt.Errorf("stream %q: at least one field is required for websocket", s.Name))
		}
	default:
		errs = append(errs, fmt.Errorf("stream %q: transport %q unknown", s.Name, s.Transport))
	}

	return errs
}

// Check if the the parameter is valid
func validateParam(p Param, ownerName string) []error {
	var errs []error

	if p.Name == "" {
		errs = append(errs, fmt.Errorf("%q: a param has an empty name", ownerName))
	}

	switch p.Type {
	case ParamTypeInt, ParamTypeString, ParamTypePath, ParamTypeEnum, ParamTypeBool:
		// Valid parameter type
	default:
		errs = append(errs, fmt.Errorf("%q: param %q has an unknown type %q", ownerName, p.Name, p.Type))
	}

	if p.Type == ParamTypeEnum && len(p.EnumValues) == 0 {
		errs = append(errs, fmt.Errorf("%q: param %q of type enum requires enum_values", ownerName, p.Name))
	}

	return errs
}

// Check if the "confirm" parameter is True
func hasConfirmParam(params []Param) bool {
	for _, p := range params {
		if p.Name == "confirm" && p.Type == ParamTypeBool {
			return true
		}
	}
	return false
}