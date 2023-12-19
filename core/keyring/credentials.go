package keyring

// NetBox : Helper method on the `netbox` credential.
func (s *Settings) NetBox() (Credential, error) {
	return s.Get("netbox", WithPasswdPromptText("Enter your NetBox API key"))
}

// Nautobot : Helper method on the `nautobot` credential.
func (s *Settings) Nautobot() (Credential, error) {
	return s.Get("nautobot", WithPasswdPromptText("Enter your Nautobot API key"))
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
