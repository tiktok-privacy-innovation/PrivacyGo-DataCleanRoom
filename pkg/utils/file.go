// Copyright 2024 TikTok Pte. Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"archive/tar"
	"compress/gzip"
	stdErrors "errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func GetDcrConfDir() string {
	return "/usr/local/dcr_conf"
}

func GetConfigFile() string {
	return fmt.Sprintf("%s/%s", GetDcrConfDir(), "config.yaml")
}

func GetDockerFile() string {
	return fmt.Sprintf("%s/%s", GetDcrConfDir(), "Dockerfile")
}

func GetWorkDirectory() (string, error) {
	return "/app/data", nil
}

// GetCurrentAbPathByExecutable Get current executable ab path
func GetCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", errors.Wrap(err, "failed to get executable path")
	}
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	if err != nil {
		return "", errors.Wrap(err, "failed to get executable path")
	}
	return res, nil
}

func CopyFile(src string, dest string) error {
	src = filepath.Clean(src)
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "failed to open source file")
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			return
		}
	}(srcFile)

	dest = filepath.Clean(dest)
	dstFile, err := os.Create(dest)
	if err != nil {
		return errors.Wrap(err, "failed to open dest file")
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			return
		}
	}(dstFile)
	// Copy file content
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return errors.Wrap(err, "failed to copy file")
	}
	return nil
}

func safeJoin(base, path string) (string, error) {
	absPath := filepath.Join(base, path)
	if !strings.HasPrefix(absPath, filepath.Clean(base)+string(os.PathSeparator)) {
		return "", fmt.Errorf("%s is an invalid path", path)
	}
	return absPath, nil
}

func UnTarGz(compressedFile string, destDirectory string) error {
	destDirectory, err := filepath.Abs(destDirectory)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to get abs path of %s", destDirectory))
	}
	if !strings.HasSuffix(compressedFile, "tar.gz") {
		return errors.Wrap(stdErrors.New("only support file ending with tar.gz"), "")
	}
	compressedFile = filepath.Clean(compressedFile)
	file, err := os.Open(compressedFile)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to open compressed file: %s", compressedFile))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	// create gzip reader
	gr, err := gzip.NewReader(file)
	if err != nil {
		return errors.Wrap(err, "failed to open gzip reader")
	}
	defer func(r *gzip.Reader) {
		_ = r.Close()
	}(gr)
	// create tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "failed to traverse directory")
		}

		target, err := safeJoin(destDirectory, hdr.Name)
		if err != nil {
			return err
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			target = filepath.Clean(target)
			if err := os.Mkdir(target, hdr.FileInfo().Mode()); err != nil {
				return errors.Wrap(err, "failed to mkdir")
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, hdr.FileInfo().Mode())
			if err != nil {
				return errors.Wrap(err, "failed to create file")
			}
			if _, err := io.Copy(f, tr); err != nil {
				return errors.Wrap(err, "failed to write file")
			}
		case tar.TypeLink:
			linkname, err := safeJoin(destDirectory, hdr.Linkname)
			if err != nil {
				return err
			}
			if err := os.Link(linkname, target); err != nil {
				return errors.Wrap(err, "failed to create link")
			}
		case tar.TypeSymlink:
			if err := os.Symlink(hdr.Linkname, target); err != nil {
				return errors.Wrap(err, "failed to create symlink")
			}
		default:
			return errors.Wrap(fmt.Errorf("%s unkown type: %c", hdr.Name, hdr.Typeflag), "")
		}
	}
	return nil
}
