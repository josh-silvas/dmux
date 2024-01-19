package keyring

// NetBox : Helper method on the `netbox` credential.
func (s *Settings) NetBox() (Credential, error) {
	return s.Get("netbox", WithPasswdPromptText("Enter your NetBox API key"))
}

// NautobotV1 : Helper method on the `nautobot` credential.
func (s *Settings) NautobotV1() (Credential, error) {
	return s.Get("nautobot_v1", WithPasswdPromptText("Enter your Nautobot API key"))
}

// NautobotV2 : Helper method on the `nautobot` credential.
func (s *Settings) NautobotV2() (Credential, error) {
	return s.Get("nautobot_v2", WithPasswdPromptText("Enter your Nautobot API key"))
}

// Netbox : Helper method on the `Netbox` credential.
func (s *Settings) Netbox() (Credential, error) {
	return s.Get("netbox", WithPasswdPromptText("Enter your Netbox API key"))
}

// DeviceAuth : Helper method on the `DeviceAuth` credential.
func (s *Settings) DeviceAuth() (Credential, error) {
	return s.Get(
		"device-auth",
		PromptUser(),
		WithUserPromptText("Enter your RADIUS, AAA or local username"),
		WithPasswdPromptText("Enter your RADIUS, AAA or local password"),
	)
}
