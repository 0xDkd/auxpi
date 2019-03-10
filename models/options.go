package models

type Option struct {
	Model

	OptionsKey   string
	OptionsValue string
}

func MigrateOptions() error {
	if db.HasTable(&Option{}) {
		err := db.DropTable(&Option{}).Error
		err = db.CreateTable(&Option{}).Error
		return err
	} else {
		err := db.CreateTable(&Option{}).Error
		return err
	}
}
