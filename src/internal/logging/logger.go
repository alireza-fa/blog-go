package logging

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})

	Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
}

//func NewLogger() Logger {
//	switch os.Getenv(constants.Logger) {
//	}
//}
