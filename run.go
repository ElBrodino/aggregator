package main

func run(cmds commands, appState *state, cmd command) error {
	err := cmds.run(appState, cmd)
	if err != nil {
		return err
	}
	return nil
}
