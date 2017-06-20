package torpedo_registry

type BotAPI struct {
	API interface{}
	Bot struct {
		Build struct {
			Build     string
			BuildDate string
			Version   string
		}
		Config struct {
			GoogleWebAppKey    *string
			SoundCloudClientID *string
			LastFmKey          *string
			LastFmSecret       *string
			PinterestToken     *string
		}
		GetCachedItem      func(string) string
		SetCachedItems     func(string, map[int]string) string
		GetCommandHandlers func() map[string]func(*BotAPI, interface{}, string)
		GetHelp            func() map[string]string
		Stats              struct {
			StartTimestamp         int64
			ProcessedMessages      int64
			ProcessedMessagesTotal int64
			ConnectedAccounts      int64
			TotalAccounts          int64
		}
		PostMessage func(interface{}, string, *BotAPI, ...interface{})
	}
	CommandPrefix string
}

var (
	handlers    = make(map[string]func(*BotAPI, interface{}, string))
	help        = make(map[string]string)
	preparsers  = make(map[string]func())
	postparsers = make(map[string]func())
)

func RegisterHandler(name string, f func(*BotAPI, interface{}, string)) {
	handlers[name] = f
}

func RegisterHelp(name, help_str string) {
	help[name] = help_str
}

func GetHandlers() map[string]func(*BotAPI, interface{}, string) {
	return handlers
}

func GetHelp() map[string]string {
	return help
}

func RegisterPreParser(name string, f func()) {
	preparsers[name] = f
}

func RegisterPostParser(name string, f func()) {
	postparsers[name] = f
}

func GetPreParsers() map[string]func() {
	return preparsers
}

func GetPostParsers() map[string]func() {
	return postparsers
}
