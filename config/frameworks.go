package config

type frameworksConfig struct {
	config *MainConfigApp
}

// ADD NEW FRAMEWORK CONFIGURED IN HERE
// 'id': {
//	  config: `methodHandler()`,
// },
var frameworks = map[string]frameworksConfig{
	"next": {
		config: newNextJs(),
	},
	"gatsby": {
		config: newGatsbyJs(),
	},
	"vite-vue3": {
		config: newViteApp(),
	},
	"react": {
		config: newReactApp(),
	},
}
