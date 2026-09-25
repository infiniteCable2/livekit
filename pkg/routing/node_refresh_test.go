// Copyright 2026 LiveKit, Inc.
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

package routing

import (
	"testing"

	"github.com/livekit/protocol/livekit"
	"github.com/stretchr/testify/require"
)

func TestLocalNodeSetNodeIP(t *testing.T) {
	node, err := NewLocalNodeFromNodeProto(&livekit.Node{Ip: "198.51.100.10"})
	require.NoError(t, err)
	previous := node.Clone()

	node.SetNodeIP("198.51.100.20")
	require.Equal(t, "198.51.100.20", node.NodeIP())
	require.Equal(t, "198.51.100.20", node.Clone().Ip)
	require.Equal(t, "198.51.100.10", previous.Ip)
}
