// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package toolregistry

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"go.uber.org/zap"
)

var (
	kubectlInstallScriptTmpl   = template.Must(template.New("kubectl").Parse(kubectlInstallScript))
	kustomizeInstallScriptTmpl = template.Must(template.New("kustomize").Parse(kustomizeInstallScript))
	helmInstallScriptTmpl      = template.Must(template.New("helm").Parse(helmInstallScript))
)

func (r *registry) installKubectl(ctx context.Context, version string) error {
	workingDir, err := ioutil.TempDir("", "kubectl-install")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDir)

	var (
		buf  bytes.Buffer
		data = map[string]string{
			"WorkingDir": workingDir,
			"Version":    version,
			"BinDir":     r.binDir,
		}
	)
	if err := kubectlInstallScriptTmpl.Execute(&buf, data); err != nil {
		r.logger.Error("failed to render kubectl install script",
			zap.String("version", version),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install kubectl %s (%v)", version, err)
	}

	var (
		script = buf.String()
		cmd    = exec.CommandContext(ctx, "/bin/sh", "-c", script)
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		r.logger.Error("failed to install kubectl",
			zap.String("version", version),
			zap.String("script", script),
			zap.String("out", string(out)),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install kubectl %s (%v)", version, err)
	}

	r.logger.Info("just installed kubectl", zap.String("version", version))
	return nil
}

func (r *registry) installKustomize(ctx context.Context, version string) error {
	workingDir, err := ioutil.TempDir("", "kustomize-install")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDir)

	var (
		buf  bytes.Buffer
		data = map[string]string{
			"WorkingDir": workingDir,
			"Version":    version,
			"BinDir":     r.binDir,
		}
	)
	if err := kustomizeInstallScriptTmpl.Execute(&buf, data); err != nil {
		r.logger.Error("failed to render kustomize install script",
			zap.String("version", version),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install kustomize %s (%v)", version, err)
	}

	var (
		script = buf.String()
		cmd    = exec.CommandContext(ctx, "/bin/sh", "-c", script)
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		r.logger.Error("failed to install kustomize",
			zap.String("version", version),
			zap.String("script", script),
			zap.String("out", string(out)),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install kustomize %s (%v)", version, err)
	}

	r.logger.Info("just installed kustomize", zap.String("version", version))
	return nil
}

func (r *registry) installHelm(ctx context.Context, version string) error {
	workingDir, err := ioutil.TempDir("", "helm-install")
	if err != nil {
		return err
	}
	defer os.RemoveAll(workingDir)

	var (
		buf  bytes.Buffer
		data = map[string]string{
			"WorkingDir": workingDir,
			"Version":    version,
			"BinDir":     r.binDir,
		}
	)
	if err := helmInstallScriptTmpl.Execute(&buf, data); err != nil {
		r.logger.Error("failed to render helm install script",
			zap.String("version", version),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install helm %s (%v)", version, err)
	}

	var (
		script = buf.String()
		cmd    = exec.CommandContext(ctx, "/bin/sh", "-c", script)
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		r.logger.Error("failed to install helm",
			zap.String("version", version),
			zap.String("script", script),
			zap.String("out", string(out)),
			zap.Error(err),
		)
		return fmt.Errorf("failed to install helm %s (%v)", version, err)
	}

	r.logger.Info("just installed helm", zap.String("version", version))
	return nil
}