// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package state

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type LocalState struct{}

func (l *LocalState) WaitForVolumeAttachment(context.Context) error {
	return nil
}

func (l *LocalState) MarkInProgress(context.Context) error {
	log.Info("Scanning is in progress")
	return nil
}

func (l *LocalState) MarkDone(_ context.Context, errs []error) error {
	if len(errs) > 0 {
		log.Errorf("scan has been completed with errors: %v", errs)
		return nil
	}
	log.Info("Scan has been completed")
	return nil
}

func (l *LocalState) IsAborted(context.Context) (bool, error) {
	return false, nil
}

func NewLocalState() (*LocalState, error) {
	return &LocalState{}, nil
}
