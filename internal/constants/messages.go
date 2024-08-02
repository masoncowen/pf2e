package constants

type NewSessionMsg struct {
    Campaign Campaign
}

type ReloadSessionMsg struct {
    SessionPath string
}

type BackMsg struct {}
