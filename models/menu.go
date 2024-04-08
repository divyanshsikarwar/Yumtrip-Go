package models

type Menu struct {
	ID          string `json:"_id"`
	Key 	   string `json:"key"`
	Items []Item `json:"items"`
	BackgroundImage string `json:"backgroundImage"`
}

func (m *Menu) GetID() string {
	return m.ID
}

func UpsertMenu (menu Menu) error {
	return nil
}