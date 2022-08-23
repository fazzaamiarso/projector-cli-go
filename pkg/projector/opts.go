package projector

import "github.com/hellflame/argparse"

type ProjectorOptions struct {
	Pwd       string
	Config    string
	Arguments []string
}

func GetOptions() (*ProjectorOptions, error) {
	parser := argparse.NewParser("Projecter", "Get all key-pair values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	args := parser.Strings("f", "foo", &argparse.Option{
		Positional:  true,
		Default: "",
        Required: false,
	})

	config := parser.String("c", "config", &argparse.Option{
		Default: "",
		Required: false,
	})
	pwd := parser.String("p", "pwd", &argparse.Option{
		Default: "",
		Required: false,
	})

	err := parser.Parse(nil)
	if err != nil {
        return nil, err
    }

	return &ProjectorOptions{
		Pwd:  *pwd,
		Config:  *config,
		Arguments:  *args,
	}, nil


}
