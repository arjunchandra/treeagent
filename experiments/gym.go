package experiments

import (
	"github.com/unixpickle/anyrl"
	gym "github.com/unixpickle/gym-socket-api/binding-go"
)

func createGymEnv(e *EnvFlags) (gym.Env, anyrl.Env, error) {
	client, err := gym.Make(e.GymHost, e.Name)
	if err != nil {
		return nil, nil, err
	}
	if e.RecordDir != "" {
		err = client.Monitor(e.RecordDir, false, false, false)
		if err != nil {
			client.Close()
			return nil, nil, err
		}
	}
	env, err := anyrl.GymEnv(client, e.GymRender)
	if err != nil {
		client.Close()
		return nil, nil, err
	}
	return client, env, nil
}
