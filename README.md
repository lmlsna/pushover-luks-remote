# pushover-luks-remote
Requests confirmation to provide the decryption password to a LUKS root partition during boot with remote-unlock.

## Configuration
  1. Rename config.sample to config.
  2. Add USER_KEY and APP_TOKEN
  3. Script "go" is called by remote-unlock hook during boot.
  4. Script "unlock" is called by "go" if user reaches endpoint.
