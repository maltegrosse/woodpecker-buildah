package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

const (
	buildahPath = "/usr/bin/buildah"
)

type options struct {
	Username      string
	Password      string
	Registry      string
	Repository    string
	ImageName     string
	Tag           string
	Context       string
	ManifestName  string
	Architectures []string
	Transport string
	Flags []string
	LoginArgs []string
	ManifestArgs []string
	BuildArgs []string
	PushArgs []string
	Steps []string

	CurrentPath string
}

func main() {
	log.Println("INFO: starting buildah plugin")
	opts, err := readEnv()
	if err != nil {
		log.Fatalln("failed to execute plugin", err)
	}
	err = execute(opts)
	if err != nil {
		log.Fatalln("failed to execute plugin", err)
	}
	log.Println("INFO: finish buildah plugin")

}
func readEnv() (*options,error){
	viper.SetEnvPrefix("plugin")
	viper.AutomaticEnv()
	viper.SetTypeByDefaultValue(true)
	viper.BindEnv("username")
	viper.BindEnv("password")
	viper.BindEnv("registry")
	viper.BindEnv("repository")
	viper.BindEnv("imagename")
	viper.SetDefault("tag", "latest")
	viper.BindEnv("tag")
	viper.SetDefault("context", "Dockerfile")
	viper.BindEnv("context")
	viper.BindEnv("manifestname")
	viper.SetDefault("architectures", []string{"amd64"})
	viper.BindEnv("architectures")
	viper.SetDefault("transport", "docker")
	viper.BindEnv("transport")
	viper.BindEnv("flags")
	viper.BindEnv("loginargs")
	viper.BindEnv("manifestargs")
	viper.BindEnv("buildargs")
	viper.BindEnv("pushargs")
	viper.SetDefault("steps", []string{"login","build","push"})
	viper.BindEnv("steps")

	var opts options
	err := viper.Unmarshal(&opts)
	if err != nil {
		return nil, err
	}
	opts.CurrentPath = os.Getenv("CI_WORKSPACE")
	return &opts, nil
}
func execute(opts *options) error {
	for _, step := range opts.Steps {
		switch step {
		case "login":
			err := login(opts)
			if err != nil {
				return err	

			}
		case "build":
			err := createManifest(opts)
			if err != nil {
				return err	

			}
			err = buildArchs(opts)
			if err != nil {
				return err	

			}
		case "push":
			err := push(opts)
			if err != nil {
				return err	

			}
		}
	}
	return nil
}
func login(opts *options) error{
	if len(opts.Username) == 0 || len(opts.Password) == 0 {
		return errors.New("username and password are required")
	}
	if len(opts.Registry) == 0 {
		return errors.New("registry is required")
	}

	cmd := exec.Command(buildahPath,"login", "--username", opts.Username, "--password-stdin",opts.Registry)
	cmd.Stdin = bytes.NewBufferString(opts.Password)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("login failed: %s", err)
	}
	log.Println("INFO: login success at registry", opts.Registry)
	return nil
}
func createManifest(opts *options) error{
	if len(opts.ManifestName) == 0 {
		opts.ManifestName = os.Getenv("CI_COMMIT_SHA")
	}
	
	args := []string{"manifest", "create", opts.ManifestName}
	args = append(args, opts.Flags...)
	args = append(args, opts.ManifestArgs...)
	out, err := exec.Command(buildahPath,args...).Output()
	if err != nil {
		return fmt.Errorf("creating manifest failed: %s\n%s", err, out)
	}
	log.Println("INFO: created manifest", opts.ManifestName)
	return nil
}
func buildArchs(opts *options)error{
	
	path := opts.CurrentPath +"/"+ opts.Context 
	tag := opts.Registry+"/"+opts.Repository+"/"+opts.ImageName+":"+opts.Tag
	for _, arch := range opts.Architectures {
		fmt.Println("INFO: building for architecture", arch)
		args:= []string{"build", "--manifest", opts.ManifestName, "--arch",arch,"--tag",tag}
		args = append(args, opts.Flags...)
		args = append(args, opts.BuildArgs...)
		args = append(args, path)
		out, err := exec.Command(buildahPath,args...).Output()
		if err != nil {
			return fmt.Errorf("building arch %s failed: %s\n%s", arch,err, out)
		}
		log.Println("INFO: build successfull for architecture", arch)
	}
	log.Println("INFO: build successfull finished for tag", tag)
	return nil
}
func push(opts *options) error{
	path := opts.Transport+ "://"+ opts.Registry+"/"+opts.Repository+"/"+opts.ImageName+":"+opts.Tag
	args:= []string{"manifest", "push","--all", opts.ManifestName, path}
	args = append(args, opts.Flags...)
	args = append(args, opts.PushArgs...)
	out, err := exec.Command(buildahPath,args...).Output()
	if err != nil {
		return fmt.Errorf("pushing image failed: %s\n%s", err, out)
	}
	log.Println("INFO: pushed successfully to", path)
	return nil
}


