module github.com/ezraisw/tanshogyo/services/user

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/render v1.0.1
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang/mock v1.7.0-rc.1
	github.com/google/wire v0.5.0
	github.com/onsi/ginkgo/v2 v2.1.4
	github.com/onsi/gomega v1.19.0
	github.com/ezraisw/tanshogyo/pkg/common v0.0.0-00010101000000-000000000000
	github.com/ezraisw/tanshogyo/pkg/gormds v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.25.2
)

require (
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/tools v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
)

replace (
	github.com/ezraisw/tanshogyo/pkg/common => ../../pkg/common
	github.com/ezraisw/tanshogyo/pkg/gormds => ../../pkg/gormds
)
