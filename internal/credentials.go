package internal

import "os"

const COOKIE_FILE = ".aoc"

type CredentialsManager struct {
	Cookie string
}

func NewCredentialsManager() (*CredentialsManager, error) {
	credentials := CredentialsManager{}
	err := credentials.LoadCookie()

	if err != nil {
		panic(err)
	}

	return &credentials, nil
}

func (cm *CredentialsManager) LoadCookie() error {
	cookie, err := cm.readCookieFile()

	if err != nil {
		panic(err)
	}

	cm.Cookie = cookie
	return nil
}

func (cm *CredentialsManager) Update() error {
	content := []byte(cm.Cookie)
	err := os.WriteFile(cm.getCredentialsFile(), content, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (cm *CredentialsManager) getCredentialsFile() string {
	return os.Getenv("HOME") + "/" + COOKIE_FILE
}

func (cm *CredentialsManager) readCookieFile() (cookie string, err error) {
	credentialsFile := cm.getCredentialsFile()
	dat, err := os.ReadFile(credentialsFile)

	if err != nil {
		return "", err
	}

	return string(dat), nil
}
