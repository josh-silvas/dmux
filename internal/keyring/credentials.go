package keyring

// Nautobot : Helper method on the `nautobot` credential.
func (s *Settings) Nautobot() (Credential, error) {
	return s.Get("nautobot", WithPasswdPromptText("Enter your Nautobot API key"))
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
