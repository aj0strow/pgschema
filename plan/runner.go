package plan

type Conn interface {
	Exec(string) error
}

type Logger interface {
	Info(string)
}

type Runner struct {
	Conn   Conn
	Logger Logger
	DryRun bool
}

func (r *Runner) Run(changes []Change) error {
	for _, change := range changes {
		r.Logger.Info(change.String())
		if !r.DryRun {
			err := r.Conn.Exec(change.String())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
